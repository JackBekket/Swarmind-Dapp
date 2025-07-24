// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {ERC20Permit} from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract Cred is ERC20, ERC20Permit, Ownable {
    address public poolAddress;
    
    constructor(uint256 initialSupply)
        ERC20("HFSwarmind", "HFSWM")
        ERC20Permit("HFSwarmind")
        Ownable(msg.sender)
    {
        _mint(msg.sender, initialSupply);
    }

    function decimals() public pure override returns (uint8) {
        return 9;
    }

    function burnFrom(address account, uint256 amount) external onlyOwner {
        _burn(account, amount);
    }
    
    function setPoolAddress(address _poolAddress) external onlyOwner {
        poolAddress = _poolAddress;
    }
    
    modifier onlyPool() {
        require(msg.sender == poolAddress, "Caller is not the pool");
        _;
    }
    
    function mint(address to, uint256 amount) external onlyPool {
        _mint(to, amount);
    }
}