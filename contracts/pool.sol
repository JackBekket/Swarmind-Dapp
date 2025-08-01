//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

//import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
//import "@openzeppelin/contracts/access/AccessControl.sol";
import "./LLMNFT.sol";
import "./FeesCalculator.sol";
import "./hfswm.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title LLM inference Pool
 * @author Sergey Ponomarev (aka JackBekket)
 * @notice This is draft for inference (LLM) provider pool
 */

contract Pool is Ownable {
  constructor(address llm_nft, address credit_, address hfswm_) Ownable(msg.sender) {
    nft = LLMNFT(llm_nft);
    credit = IERC20(credit_);
    hfswm = HFSWM(hfswm_);
}

    //global vars
    enum Pool_type {
        private_pool,
        public_pool
    }
    enum Pay_type {
        token,
        request
    }

    //address owner;
    Pool_type pt;
    //IERC721Metadata LLMNFT_token;
    LLMNFT nft;

    IERC20 credit;
    
    //huggingface token
    HFSWM public hfswm;


    uint service_fee_percent = 1; // 1%
    // TODO: add constant service address where we will collect fee
    //SUGGESTION @daseinsucks: use multisig?

    mapping(string => address) worker_wallets; // local_ai public key -> worker wallet address
    mapping(address => string[]) wallets_workers; // wallet address -> lai public key
    mapping(string => bool) public blacklist; //address => banned (true/false)
    mapping(string => bool) public isApproved; //for private pools
    mapping(address => bool) public HFwhitelist; //providers from huggingface
    //TODO: events

    event Deposit(address user, uint amount);
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

    // TODO: add other metadata from edgevpn into worker struct and make update methods

    // LLM models
    struct LLM_meta {
        uint256 token_id;
        uint256 author_royalty;
        address author_wallet;
        uint256 hw_price_per_input_token;
        uint256 hw_price_per_output_token;
        Pay_type pay_type_;
    }

    //We need this to avoid "Stack too deep" in ProcessResponse
    struct CostDetails {
        uint256 hwpin;
        uint256 hwpout;
        uint256 cost_hw;
        uint256 awp;
        uint256 a_cost;
        uint256 tokens_sum;
    }

    //LLM_meta[] llm_list;    // TODO: idk if we should use mapping or array.
    //SUGGESTION @daseinsucks: I say we use map, since for crud methods it'll be cheaper to fetch from map than iterate through array
    mapping(uint => LLM_meta) llm_list;

    function GetMetaLLM(uint llm_id) public view returns (LLM_meta memory) {
        return llm_list[llm_id];
    }

    function GetTotalPrice(uint llm_id) public view returns (uint) {
        LLM_meta memory lm = GetMetaLLM(llm_id);
        uint ap = lm.author_royalty;
        uint hwpin = lm.hw_price_per_input_token;
        uint hwpout = lm.hw_price_per_output_token;
        uint tp = ap + hwpin + hwpout;
        return tp;
    }

    function RegisterWorker(string memory lai_public_key) public {
        require(!blacklist[lai_public_key], "This key is in blacklist");
        require(
            worker_wallets[lai_public_key] == address(0),
            "Key already registered"
        );
        require(
            wallets_workers[msg.sender].length < 20,
            "Maximum number of workers (20) reached"
        );

        wallets_workers[msg.sender].push(lai_public_key);
        worker_wallets[lai_public_key] = msg.sender;

        if (pt == Pool_type.public_pool) {
            isApproved[lai_public_key] = true;
        } else {
            isApproved[lai_public_key] = false;
        }

        emit NewWorker(msg.sender, isApproved[lai_public_key]);
    }

    function GetWorkerAddress(
        string memory lai_id
    ) public view returns (address) {
        address worker = worker_wallets[lai_id];
        return worker;
    }

    function ApproveWorker(string memory lai_public_key) public onlyOwner {
        isApproved[lai_public_key] = true;
    }

    function Unregister(string memory lai_pub_key) public {
        require(
            worker_wallets[lai_pub_key] == msg.sender,
            "Not your worker key"
        );
        string[] storage keys = wallets_workers[msg.sender];
        for (uint i = 0; i < keys.length; i++) {
            if (keccak256(bytes(keys[i])) == keccak256(bytes(lai_pub_key))) {
                keys[i] = keys[keys.length - 1];
                keys.pop();
                break;
            }
        }

        delete worker_wallets[lai_pub_key];
    }

    function UpdateWorker(
    string memory old_lai_pub_key,
    string memory new_lai_pub_key
) public {
    require(
        worker_wallets[old_lai_pub_key] == msg.sender,
        "Not your worker key"
    );
    require(
        worker_wallets[new_lai_pub_key] == address(0),
        "New key already registered"
    );
    require(!blacklist[new_lai_pub_key], "New key is blacklisted");

    string[] storage keys = wallets_workers[msg.sender];
    bool updated = false;

    for (uint i = 0; i < keys.length; i++) {
        if (keccak256(bytes(keys[i])) == keccak256(bytes(old_lai_pub_key))) {
            keys[i] = new_lai_pub_key;
            updated = true;
            break;
        }
    }

    require(updated, "Old key not found");

    delete worker_wallets[old_lai_pub_key];
    worker_wallets[new_lai_pub_key] = msg.sender;

    isApproved[new_lai_pub_key] = isApproved[old_lai_pub_key];
    delete isApproved[old_lai_pub_key];

    emit NewWorker(msg.sender, isApproved[new_lai_pub_key]);
}

    // TODO: IMPORTANT -- change visibility to only Owner or only factory(?). idk what would be tipical deployment process prolly just onlyOwner works fine.
    function Ban(string memory lai_pub_key) public onlyOwner {
        blacklist[lai_pub_key] = true;
    }

    function Unban(string memory lai_pub_key) public onlyOwner {
        blacklist[lai_pub_key] = false;
    }

    function AddToHFWhitelist(address addr) public onlyOwner {
    HFwhitelist[addr] = true;
}

    function RemoveFromHFWhitelist(address addr) public onlyOwner {
            HFwhitelist[addr] = false;
}


    function AddModel(
        uint token_id,
        uint hw_price_in,
        uint hw_price_out
    ) public onlyOwner {
        LLMNFT.LLM memory llm_struct = nft.GetLLM(token_id);
        LLM_meta memory lm;
        Pay_type pt_;
        LLMNFT.Llm_type lm_type = llm_struct.model_type;
        if (lm_type == LLMNFT.Llm_type.text) {
            pt_ = Pay_type.token; // tarification per token
        } else {
            pt_ = Pay_type.request;
        }
        lm.pay_type_ = pt_;
        lm.author_royalty = llm_struct.royalty_price;
        lm.author_wallet = llm_struct.author_wallet;
        lm.token_id = token_id;
        lm.hw_price_per_input_token = hw_price_in;
        lm.hw_price_per_output_token = hw_price_out;
        //llm_list.push(lm);
        llm_list[token_id] = lm;
    }

    function GetModel(
        uint256 llm_id
    ) public view returns (Pool.LLM_meta memory) {
        return llm_list[llm_id];
    }

    function UpdateModel(
        uint256 llm_id,
        uint new_royalty,
        address new_address
    ) public onlyOwner {
        Pool.LLM_meta memory lm = llm_list[llm_id];
        lm.author_royalty = new_royalty;
        lm.author_wallet = new_address;
        llm_list[llm_id] = lm;
    }

    function DeleteModel(uint256 llm_id) public onlyOwner {
        delete llm_list[llm_id];
    }

    function DepositCredit(uint amount) public {
        require(credit.transferFrom(msg.sender, address(this), amount));
        emit Deposit(msg.sender, amount);
    }

    // Check that user have enough money before calling request
    function Pre_request(uint llm_id, address user) public view {
        LLMNFT.LLM memory llms = nft.GetLLM(llm_id);
        LLMNFT.Llm_type lm_type = llms.model_type;
        Pay_type pt_;
        uint max_context = llms.max_context_window;
        LLM_meta memory meta = GetMetaLLM(llm_id);
        uint hw_price = meta.hw_price_per_input_token;
        //uint royalty = meta.author_royalty;
        uint price;
        if (lm_type == LLMNFT.Llm_type.text) {
            pt_ = Pay_type.token; // tarification per token
            price = max_context * hw_price;
        } else {
            pt_ = Pay_type.request;
            price = hw_price * 1;
        }
    
        uint balance = credit.balanceOf(user);
         require(balance >= price, "ERC20: insufficient balance");
    }

    function CalculateServiceFee(uint hw_cost) public view returns (uint) {
        uint service_fee = FeesCalculator.CalculateAbstractFee(
            hw_cost,
            100,
            service_fee_percent
        );
        return service_fee;
    }

    function ProcessResponse(
    uint request_id,
    string memory worker_id,
    uint llm_id,
    uint256 llmInputTokens,
    uint256 llmOutputTokens,
    uint processingTime
) public {
    require(!blacklist[worker_id], "This worker is in blacklist");
    require(isApproved[worker_id], "This worker is not yet approved");

    LLM_meta memory lm = GetMetaLLM(llm_id);
    Pay_type pt_ = lm.pay_type_;
    CostDetails memory cd;

    if (pt_ == Pay_type.token) {
        cd.tokens_sum = llmInputTokens + llmOutputTokens;
        cd.hwpin = lm.hw_price_per_input_token;
        cd.hwpout = lm.hw_price_per_output_token;
        cd.cost_hw =
            cd.hwpin *
            llmInputTokens +
            cd.hwpout *
            llmOutputTokens;
        cd.a_cost = lm.author_royalty * cd.tokens_sum;
    } else {
        cd.hwpin = lm.hw_price_per_input_token;
        cd.awp = lm.author_royalty;
        cd.cost_hw = cd.hwpin;
        cd.a_cost = cd.awp;
    }

    address worker = GetWorkerAddress(worker_id);
    address author = lm.author_wallet;

    if (HFwhitelist[worker]) {
      hfswm.mint(worker, cd.cost_hw);
    } else {
        require(credit.transfer(worker, cd.cost_hw), "Credit to worker failed");
    }


    if (HFwhitelist[author]) {
      hfswm.mint(author, cd.a_cost);
    } else {
        require(credit.transfer(author, cd.a_cost), "Credit to author failed");
    }

    emit Response(
        request_id,
        llm_id,
        worker,
        llmInputTokens,
        llmOutputTokens,
        cd.cost_hw + cd.a_cost,
        processingTime,
        block.timestamp
    );

    emit Payout(worker, llm_id, cd.cost_hw);
}
}