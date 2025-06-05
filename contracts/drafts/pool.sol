//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;


//import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
//import "@openzeppelin/contracts/access/AccessControl.sol";
import "./LLMNFT.sol";
import "@openzeppelin/contracts/access/Ownable.sol";



/**
 * @title LLM inference Pool
 * @author Sergey Ponomarev (aka JackBekket)
 * @notice This is draft for inference (LLM) provider pool
 */



contract Pool is Ownable {


    constructor(address llm_nft) Ownable(msg.sender){
        nft = LLMNFT(llm_nft);
    }


    //global vars
    enum Pool_type {privat_pool,public_pool} 
    enum Pay_type {token,request}

    //address owner;
    Pool_type pt;
    //IERC721Metadata LLMNFT_token;
    LLMNFT nft;


    mapping (string => address) worker_wallets;     // local_ai public key -> worker wallet address
    mapping (address => string) wallets_workers;    // wallet address -> lai public key
    
    // TODO: add other metadata from edgevpn into worker struct and make update methods


    // LLM models
    struct LLM_meta {
        uint256 token_id;
        uint256 author_royalty;
        address author_wallet;
        uint256 hw_price_per_token; // hardware provider is setting price per 1m tokens, so there should be conversion here.
        Pay_type pay_type_;
    }

    LLM_meta[] llm_list;    // TODO: idk if we should use mapping or array.



    // TODO: think about possible security hacks
    // TODO: add approve mechanism for private federations
    function RegisterWorker(string memory lai_public_key) public {
        if (pt== Pool_type.public_pool) {
            bytes memory pub_key = bytes(wallets_workers[msg.sender]);
            require (pub_key.length == 0);
            wallets_workers[msg.sender] = lai_public_key;
            worker_wallets[lai_public_key] = msg.sender;
           // return true;
        }
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
        address worker_address = worker_wallets[lai_pub_key];
        delete worker_wallets[lai_pub_key];
        delete wallets_workers[worker_address];

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
       llm_list.push(lm);
    }

    // TODO: add Read (getter), Update, Delete funcs for llm models.


}