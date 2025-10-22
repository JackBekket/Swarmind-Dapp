// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "./LLMNFT.sol";
import "./FeesCalculator.sol";
import "./hfswm.sol";
import "./WalletHub.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title LLM inference Pool
 * @author zer0_eX (JackBekket) 
 * @author m0rs (daseinsucks) 
 * @notice Pool for inference providers connected to shared WalletHub
 */

contract Pool is Ownable {
    using SafeERC20 for IERC20;

    enum Pool_type {
        private_pool,
        public_pool
    }
    enum Pay_type {
        token,
        request
    }

    LLMNFT public nft;
    IERC20 public credit;
    HFSWM public hfswm;
    WalletHub public walletHub;

    Pool_type public pt;
    uint256 private service_fee_percent = 1; // 1%
    address private treasury;

    mapping(string => address) worker_wallets; // lai pubkey -> worker wallet
    mapping(address => string[]) wallets_workers; // wallet -> lai pubkeys
    mapping(string => bool) public blacklist;
    mapping(string => bool) public isApproved;
    mapping(address => bool) public HFwhitelist;

    event Payout(address provider, uint256 llm_id, uint amount);
    event Response(
        uint256 id,
        uint256 llm_id,
        address worker,
        uint256 llmInputTokens,
        uint256 llmOutputTokens,
        uint256 cost,
        uint256 processingTime,
        uint256 timestamp
    );
    event NewWorker(address worker, bool isApproved);
    event ModelAdded(uint256 token_id);
    event ModelDeleted(uint256 token_id);

    struct LLM_meta {
        uint256 token_id;
        uint256 author_royalty;
        address author_wallet;
        uint256 hw_price_per_input_token;
        uint256 hw_price_per_output_token;
        Pay_type pay_type_;
    }

    struct CostDetails {
        uint256 hwpin;
        uint256 hwpout;
        uint256 cost_hw;
        uint256 awp;
        uint256 a_cost;
        uint256 tokens_sum;
    }

    mapping(uint => LLM_meta) llm_list;

    constructor(
        address llm_nft,
        address credit_,
        address hfswm_,
        address walletHub_,
         Pool_type _pt
    ) Ownable(msg.sender) {
        nft = LLMNFT(llm_nft);
        credit = IERC20(credit_);
        hfswm = HFSWM(hfswm_);
        walletHub = WalletHub(walletHub_);
        treasury = msg.sender;
        pt = _pt;
    }

    // ==================== Workers ====================

    function RegisterWorker(string memory lai_public_key) public {
        require(!blacklist[lai_public_key], "This key is blacklisted");
        require(worker_wallets[lai_public_key] == address(0), "Already registered");
        require(wallets_workers[msg.sender].length < 20, "Max workers reached");

        wallets_workers[msg.sender].push(lai_public_key);
        worker_wallets[lai_public_key] = msg.sender;

        isApproved[lai_public_key] = (pt == Pool_type.public_pool);
        emit NewWorker(msg.sender, isApproved[lai_public_key]);
    }

    function ApproveWorker(string memory lai_public_key) public onlyOwner {
        isApproved[lai_public_key] = true;
    }

    function Unregister(string memory lai_pub_key) public {
        require(worker_wallets[lai_pub_key] == msg.sender, "Not your key");
        string[] storage keys = wallets_workers[msg.sender];
        for (uint i = 0; i < keys.length; i++) {
            if (keccak256(bytes(keys[i])) == keccak256(bytes(lai_pub_key))) {
                keys[i] = keys[keys.length - 1];
                keys.pop();
                break;
            }
        }
        delete worker_wallets[lai_pub_key];
        delete isApproved[lai_pub_key];
    }

    function GetWorkerAddress(string memory lai_id) public view returns (address) {
        return worker_wallets[lai_id];
    }

    function Ban(string memory lai_pub_key) public onlyOwner {
        blacklist[lai_pub_key] = true;
    }

    function Unban(string memory lai_pub_key) public onlyOwner {
        blacklist[lai_pub_key] = false;
    }

     function setServiceFeePercent(uint256 _percent) external onlyOwner {
        require(_percent <= 100, "Invalid percent"); 
        service_fee_percent = _percent;
    }

    // ==================== HuggingFace Whitelist ====================

    function AddToHFWhitelist(address addr) public onlyOwner {
        HFwhitelist[addr] = true;
    }

    function RemoveFromHFWhitelist(address addr) public onlyOwner {
        HFwhitelist[addr] = false;
    }

    // ==================== Models ====================

    function AddModel(uint token_id, uint hw_price_in, uint hw_price_out) public onlyOwner {
        LLMNFT.LLM memory llm_struct = nft.GetLLM(token_id);
        Pay_type pt_ = llm_struct.model_type == LLMNFT.Llm_type.text
            ? Pay_type.token
            : Pay_type.request;

        llm_list[token_id] = LLM_meta({
            token_id: token_id,
            author_royalty: llm_struct.royalty_price,
            author_wallet: llm_struct.author_wallet,
            hw_price_per_input_token: hw_price_in,
            hw_price_per_output_token: hw_price_out,
            pay_type_: pt_
        });

        emit ModelAdded(token_id);
    }

        function UpdateModel(
        uint256 llm_id,
        uint new_royalty,
        address new_address,
        uint new_hw_price_in,
        uint new_hw_price_out
    ) public onlyOwner {
        Pool.LLM_meta memory lm = llm_list[llm_id];
        lm.author_royalty = new_royalty;
        lm.author_wallet = new_address;
        lm.hw_price_per_input_token = new_hw_price_in;
        lm.hw_price_per_output_token= new_hw_price_out;
        llm_list[llm_id] = lm;
    }

    function DeleteModel(uint llm_id) public onlyOwner {
        delete llm_list[llm_id];
        emit ModelDeleted(llm_id);
    }

    function GetModel(uint llm_id) public view returns (LLM_meta memory) {
        return llm_list[llm_id];
    }

    // ==================== Requests ====================

    function CalculateServiceFee(uint hw_cost) public view returns (uint) {
        return FeesCalculator.CalculateAbstractFee(hw_cost, 100, service_fee_percent);
    }

    function Pre_request(uint llm_id, address user) public view {
    LLMNFT.LLM memory llms = nft.GetLLM(llm_id);
    LLMNFT.Llm_type lm_type = llms.model_type;
    uint max_context = llms.max_context_window;
    LLM_meta memory meta = GetModel(llm_id);
    uint hw_price = meta.hw_price_per_input_token;
    uint price;

    if (lm_type == LLMNFT.Llm_type.text) {
        price = max_context * hw_price;
    } else {
        price = hw_price * 1;
    }
    
    require(walletHub.userBalance(user) >= price, "Insufficient deposited balance");
}

    function ProcessResponse(
        uint request_id,
        string memory worker_id,
        uint llm_id,
        uint256 llmInputTokens,
        uint256 llmOutputTokens,
        uint processingTime,
        address user,
        uint pics,
        uint imageArea, 
        uint imageSteps    
    ) public onlyOwner {
        require(!blacklist[worker_id], "Blacklisted worker");
        require(isApproved[worker_id], "Worker not approved");

        CostDetails memory cd = _calculateCosts(llm_id, llmInputTokens, llmOutputTokens, pics, imageArea, imageSteps);
        uint256 totalCost = cd.cost_hw + cd.a_cost;

        walletHub.debit(user, totalCost);
        _distributePayments(worker_id, llm_id, totalCost, cd);

        emit Response(
            request_id,
            llm_id,
            GetWorkerAddress(worker_id),
            llmInputTokens,
            llmOutputTokens,
            totalCost,
            processingTime,
            block.timestamp
        );
    }

    function _calculateCosts(
    uint llm_id,
    uint256 llmInputTokens,
    uint256 llmOutputTokens,
    uint pics,
    uint imageArea, 
    uint imageSteps    
) internal view returns (CostDetails memory) {
    LLM_meta memory lm = llm_list[llm_id];
    CostDetails memory cd;

    if (lm.pay_type_ == Pay_type.token) {
        // Текстовые LLM
        cd.tokens_sum = llmInputTokens + llmOutputTokens;
        cd.hwpin = lm.hw_price_per_input_token;
        cd.hwpout = lm.hw_price_per_output_token;
        cd.cost_hw = cd.hwpin * llmInputTokens + cd.hwpout * llmOutputTokens;
        cd.a_cost = lm.author_royalty * cd.tokens_sum;
    } else {
        cd.hwpin = lm.hw_price_per_input_token;
        cd.awp = lm.author_royalty;
        uint baseArea = 1024 * 1024;
        uint baseSteps = 25;

        uint areaMultiplier = (imageArea * 1e18) / baseArea; 
        uint stepsMultiplier = (imageSteps * 1e18) / baseSteps;

        uint totalMultiplier = (areaMultiplier * stepsMultiplier) / 1e18;

        cd.cost_hw = (cd.hwpin * pics * totalMultiplier) / 1e18;
        cd.a_cost = cd.awp;
    }

    return cd;
}


    function _distributePayments(
        string memory worker_id,
        uint llm_id,
        uint256 totalCost,
        CostDetails memory cd
    ) internal {
        uint256 serviceFee = CalculateServiceFee(totalCost);
        uint256 totalWithoutFee = totalCost - serviceFee;
        require(totalCost > 0, "zero total");

        uint256 workerPortion = (cd.cost_hw * totalWithoutFee) / totalCost;
        uint256 authorPortion = (cd.a_cost * totalWithoutFee) / totalCost;
        uint256 remainder = totalWithoutFee - workerPortion - authorPortion;

        address worker = GetWorkerAddress(worker_id);
        address author = llm_list[llm_id].author_wallet;

        if (worker != address(0)) {
            if (HFwhitelist[worker]) {
                hfswm.mint(worker, workerPortion);
            } else {
                walletHub.pay(worker, workerPortion);
            }
        }

        if (author != address(0)) {
            if (HFwhitelist[author]) {
                hfswm.mint(author, authorPortion);
            } else {
                walletHub.pay(author, authorPortion);
            }
        }

        walletHub.treasuryTransfer(serviceFee + remainder);
        emit Payout(worker, llm_id, workerPortion);
    }

}