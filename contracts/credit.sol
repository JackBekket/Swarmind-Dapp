// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

 /// @notice made barebones mint and burn functions, TBU 
contract Cred is ERC20, Ownable {
    uint256 public tokensPerEth;

    constructor(uint256 initialSupply) ERC20("Swarmind", "SWM") Ownable(msg.sender) {
        tokensPerEth = 1000 * 10 ** decimals();
        _mint(msg.sender, initialSupply);
    }

    function decimals() public view override returns (uint8) {
        return 9;
    }

    receive() external payable {
        require(msg.value > 0, "No ETH sent");
        uint256 tokensToMint = (msg.value * tokensPerEth) / 1 ether;
        _mint(msg.sender, tokensToMint);
    }

    function burn(uint256 amount) external {
        require(balanceOf(msg.sender) >= amount, "Not enough tokens");
        uint256 ethToReturn = (amount * 1 ether) / tokensPerEth;
        require(address(this).balance >= ethToReturn, "Contract has not enough ETH");

        _burn(msg.sender, amount);
        payable(msg.sender).transfer(ethToReturn);
    }

   
    function setTokensPerEth(uint256 newRate) external onlyOwner {
        require(newRate > 0, "Rate must be > 0");
        tokensPerEth = newRate;
    }
}
