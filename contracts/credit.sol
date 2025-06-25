// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract Cred is ERC20, ERC20Permit, Ownable {
    constructor(uint256 initialSupply)
        ERC20("Swarmind", "SWM")
        ERC20Permit("Swarmind")
        Ownable(msg.sender)
    {
        _mint(msg.sender, initialSupply);
    }

    function decimals() public pure override returns (uint8) {
        return 9;
    }
}