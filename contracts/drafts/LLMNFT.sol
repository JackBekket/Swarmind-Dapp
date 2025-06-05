//SPDX-License-Identifier: Business License
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
//import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";



/**
 * @title LLM as NFT
 * @author zer0_eX (JackBekket) 
 * @notice
 * 
 * Input data:
 *   model: 
 *   huggingface://TheDrummer/Tiger-Gemma-9B-v1-GGUF/Tiger-Gemma-9B-v1a-Q4_K_M.gguf
 *  # https://huggingface.co/TheDrummer/Tiger-Gemma-9B-v1-GGUF
 * 
 *  base url should be: https://huggingface.co/
 *  token_id should be masked as 'TheDrummer/Tiger-Gemma-9B-v1-GGUF/Tiger-Gemma-9B-v1a-'   --- there should be a variable for that?
 *  token_ids should be differ via quantization methods, so each token id is something like Q4_K_M.gguf
 *  there should be a price for a whole model, without diffirienting from a quantization type... ?
 */




// collection, singleton NFT(?)
contract LLMNFT is ERC721URIStorage{

    uint256 token_ids_count;

    struct LLM {
    string author;
    string hfid;    // huggingface id
    // model_config --- yaml config of a model for local ai
    uint256 Price;  //  // text could be charged per token or per 1m llm tokens, image, video, voice, etc -- per request.
    address author_wallet;
    llm_type model_type;
    }

    enum llm_type {text,image}  // text could be charged per token or per 1m llm tokens, image, video, voice, etc -- per request.
    
    // TODO: think how to set up and store quantisation type is neccessary. mostly used q4_k_m as default median; fp16 (floating point 16) is original file without quantisation
    //enum quantisation_type {q2_k,q3_k_m,q4_k_m ..... q8_0, fp16}

    mapping (uint256 => LLM) Models_metadata;
    mapping (string =>uint256) Hf_models; // from hfid to token_id




    constructor(string memory name_, string memory smbl_) ERC721(name_, smbl_) ERC721URIStorage() {
 
    }


    // base url. may be used as hugginface link
    function _baseURI() internal view virtual override returns (string memory) {
        return "https://huggingface.co/";
    }




    /**
     *  
     */
    /*
    function CreateItem(string memory file_id) public returns (uint256 token_id) {
        token_ids_count +=1;
        token_id = token_ids_count;
        _safeMint(msg.sender,token_id);
        _setTokenURI(token_id,file_id);

    }
    */


   function CreateLLM_NFT(string memory hf_id_, uint256 price, address wallet, llm_type model_type_) public returns (uint256) {

        LLM memory llm;
        //llm.author = author_;   // 'TheDrummer' -- hf user
        llm.hfid = hf_id_; // 'TheDrummer/Tiger-Gemma-9B-v1-GGUF' -- namespace/model
        llm.Price = price;          // price per token to go as royalty
        llm.author_wallet = wallet; // author wallet
        llm.model_type = model_type_;
        token_ids_count +=1;
        uint256 token_id = token_ids_count;
        _safeMint(msg.sender,token_id);
        Models_metadata[token_id] = llm;
        Hf_models[hf_id_] = token_id;
        string memory file_id = llm.hfid;
        _setTokenURI(token_id,file_id);
        return token_id;
   }

}