//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;


//import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
//import "@openzeppelin/contracts/access/AccessControl.sol";
import "./LLMNFT.sol";
import "./FeesCalculator.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";



/**
 * @title LLM inference Pool
 * @author Sergey Ponomarev (aka JackBekket)
 * @notice This is draft for inference (LLM) provider pool
 */



contract Pool is Ownable {


    constructor(address llm_nft, address credit_) Ownable(msg.sender){
        nft = LLMNFT(llm_nft);
        credit = IERC20(credit_);
    }


    //global vars
    enum Pool_type {private_pool,public_pool} 
    enum Pay_type {token,request}

    //address owner;
    Pool_type pt;
    //IERC721Metadata LLMNFT_token;
    LLMNFT nft;

    IERC20 credit;



    uint service_fee_percent = 1; // 1%
    // TODO: add constant service address where we will collect fee
    //SUGGESTION @daseinsucks: use multisig?


    mapping (string => address) worker_wallets;     // local_ai public key -> worker wallet address
    mapping (address => string) wallets_workers;    // wallet address -> lai public key
    mapping (address => uint) user_deposits;
    mapping (string => bool) public blacklist; //address => banned (true/false)
    mapping (string => bool) public isApproved; //for private pools
    //TODO: events


    event Deposit(address user, uint amount);
    event Payout(address provider,uint256 llm_id, uint amount);
    event Response(
        uint256 id,     
        uint256 llm_id,
        address worker,
        uint256 llmTokens,
        uint256 cost,
        uint256 processingTime,
        uint256 timestamp
    );
    event NewWorker (address worker, bool isApproved);


    // TODO: add other metadata from edgevpn into worker struct and make update methods


    // LLM models
    struct LLM_meta {
        uint256 token_id;
        uint256 author_royalty;
        address author_wallet;
        uint256 hw_price_per_token; // hardware provider is setting price per 1m tokens, so there should be conversion here.
        Pay_type pay_type_;
    }


    //LLM_meta[] llm_list;    // TODO: idk if we should use mapping or array.
    //SUGGESTION @daseinsucks: I say we use map, since for crud methods it'll be cheaper to fetch from map than iterate through array
    mapping (uint => LLM_meta) llm_list;


    function GetMetaLLM(uint llm_id) public view returns (LLM_meta memory) {
        return llm_list[llm_id];
    }

    function GetTotalPrice(uint llm_id) public view returns (uint) {
        LLM_meta memory lm = GetMetaLLM(llm_id);
        uint ap = lm.author_royalty;
        uint hwp = lm.hw_price_per_token;
        uint tp = ap + hwp;
        return tp;
    }


    // TODO: think about possible security hacks
    function RegisterWorker(string memory lai_public_key) public {
        require(!blacklist[lai_public_key], "This key is in blacklist");
            bytes memory pub_key = bytes(wallets_workers[msg.sender]);
            require (pub_key.length == 0);
            wallets_workers[msg.sender] = lai_public_key;
            worker_wallets[lai_public_key] = msg.sender;
         if (pt== Pool_type.public_pool) {
            isApproved[lai_public_key] = true;
        } else {
            isApproved[lai_public_key] = false;
        }
        emit NewWorker(msg.sender, isApproved);
    }

    function GetWorkerAddress(string memory lai_id) public view returns (address) {
        address worker = worker_wallets[lai_id];
        return worker;
    }

    function ApproveWorker(string memory lai_public_key) public onlyOwner {
        isApproved[lai_public_key] = true;
    }

    function Unregister() public  {
        bytes memory pub_key = bytes(wallets_workers[msg.sender]);
        require(pub_key.length!=0);
        string memory lai_pub_key = wallets_workers[msg.sender];
        delete wallets_workers[msg.sender];
        delete worker_wallets[lai_pub_key];
    }



    // TODO: IMPORTANT -- change visibility to only Owner or only factory(?). idk what would be tipical deployment process prolly just onlyOwner works fine.
    function Ban(string memory lai_pub_key) public onlyOwner {
     blacklist[lai_pub_key] = true;
    }

     function Unban(string memory lai_pub_key) public onlyOwner {
     blacklist[lai_pub_key] = false;
    }


    function AddModel(uint token_id, uint hw_price) public onlyOwner  {
       LLMNFT.LLM memory llm_struct = nft.GetLLM(token_id);
       LLM_meta memory lm;
       Pay_type pt_;
       LLMNFT.Llm_type lm_type = llm_struct.model_type;
       if (lm_type == LLMNFT.Llm_type.text) {
            pt_ = Pay_type.token;    // tarification per token
       } else {
        pt_ = Pay_type.request;
       }
       lm.pay_type_ = pt_;
       lm.author_royalty = llm_struct.royalty_price;
       lm.author_wallet = llm_struct.author_wallet;
       lm.token_id = token_id;
       lm.hw_price_per_token = hw_price;
       //llm_list.push(lm);
       llm_list[token_id] = lm;
    }

    function GetModel (uint256 llm_id) public view returns (Pool.LLM_meta memory) {
        return llm_list[llm_id];
    }

    function UpdateModel (uint256 llm_id, uint new_royalty, address new_address) public onlyOwner  {
       Pool.LLM_meta memory lm = llm_list[llm_id];
       lm.author_royalty = new_royalty;
       lm.author_wallet = new_address;
       llm_list[llm_id] = lm;
    }

    function DeleteModel (uint256 llm_id) public onlyOwner  {
    delete llm_list[llm_id];
    }
        

    function DepositCredit(uint amount)  public { 
       require (credit.transferFrom(msg.sender,address(this),amount));
       user_deposits[msg.sender] += amount;
       emit Deposit(msg.sender,amount);
    }


    // Check that user have enough money before calling request
    function Pre_request(uint llm_id, address user) public view {
        LLMNFT.LLM memory llms = nft.GetLLM(llm_id);
        LLMNFT.Llm_type lm_type = llms.model_type;
        Pay_type pt_;
        uint max_context = llms.max_context_window;
        LLM_meta memory meta = GetMetaLLM(llm_id);
        uint hw_price = meta.hw_price_per_token;
        //uint royalty = meta.author_royalty;
        uint price;
       if (lm_type == LLMNFT.Llm_type.text) {
            pt_ = Pay_type.token;    // tarification per token
            price = max_context * hw_price;  
       } else {
        pt_ = Pay_type.request;
        price = hw_price * 1;
       }
        uint balance = user_deposits[user];
        require (balance >= price, "user balance below max context window");
    }


    function CalculateServiceFee(uint hw_cost) public view returns (uint) {
        uint service_fee = FeesCalculator.CalculateAbstractFee(hw_cost,100,service_fee_percent);
        return service_fee;
    }

    
    function ProcessResponse(uint request_id, string memory worker_id ,uint llm_id, uint256 llmTokens, uint processingTime) public  {
        require(!blacklist[worker_id], "This worker is in blacklist"); 
         require(isApproved[worker_id], "This worker is not yet approved"); 
        uint tprice = GetTotalPrice(llm_id);
        LLM_meta memory lm = GetMetaLLM(llm_id);
        Pay_type pt_ = lm.pay_type_;
        uint hwp;
        uint cost_hw;
        uint awp;
        uint a_cost;
        
        if (pt_ == Pay_type.token) {
            hwp = lm.hw_price_per_token;
            cost_hw = hwp * llmTokens;
            a_cost = lm.author_royalty * llmTokens;
        } else {
            hwp = lm.hw_price_per_token;
            awp = lm.author_royalty;
            cost_hw = hwp; // per request
            a_cost = awp;
        }

        address worker = GetWorkerAddress(worker_id);
        address author = lm.author_wallet;
        
        require (credit.transfer(worker,cost_hw));
        require(credit.transfer(author,a_cost));

        uint timestamp = block.timestamp;

        emit Response(request_id,llm_id,worker,llmTokens,tprice,processingTime,timestamp);
        emit Payout(worker,llm_id,cost_hw);

    }
    


}