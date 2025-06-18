// 
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

//TODO:
// should be mintable and burnable


contract Cred is ERC20 {
    constructor(uint256 initialSupply) ERC20("Swarmind", "SWM") {
        _mint(msg.sender, initialSupply);
    }

    // micro usd as base decimal
    function decimals() public view override returns (uint8) {
        return 9;
    }

}