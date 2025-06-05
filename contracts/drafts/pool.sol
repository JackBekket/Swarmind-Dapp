//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;


import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol";
//import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Pool is Ownable {


    constructor() Ownable(msg.sender){
        
    }


    //global vars
    enum pool_type {privat_pool,public_pool} 

    //address owner;
    pool_type pt;
    IERC721Metadata LLMNFT_token;


    mapping (string => address) worker_wallets;     // local_ai public key -> worker wallet address
    mapping (address => string) wallets_workers;    // wallet address -> lai public key
    
    // TODO: add other metadata from edgevpn into worker struct and make update methods


    // LLM models
    struct LLM {
        uint256 token_id;
        uint256 author_royalty;
        uint256 hw_price_per_token; // hardware provider is setting price per 1m tokens, so there should be conversion here.
    }



    // TODO: think about possible security hacks
    // TODO: add approve mechanism for private federations
    function RegisterWorker(string memory lai_public_key) public {
        if (pt== pool_type.public_pool) {
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




}