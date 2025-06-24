// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract Cred is ERC20, Ownable {
    IERC20 public weth;
    uint256 public tokensPerWETH;

    constructor(address _weth, uint256 initialSupply) ERC20("Swarmind", "SWM") Ownable(msg.sender) {
        weth = IERC20(_weth);
        tokensPerWETH = 1000 * 10 ** decimals();
        _mint(msg.sender, initialSupply);
    }

    function decimals() public pure override returns (uint8) {
        return 9;
    }

    function buy(uint256 wethAmount) external {
        require(wethAmount > 0, "Zero WETH");
        require(weth.transferFrom(msg.sender, address(this), wethAmount), "Transfer failed");

        uint256 tokensToMint = (wethAmount * tokensPerWETH) / 1 ether;
        _mint(msg.sender, tokensToMint);
    }

    function sell(uint256 tokenAmount) external {
        require(balanceOf(msg.sender) >= tokenAmount, "Not enough tokens");
        uint256 wethToReturn = (tokenAmount * 1 ether) / tokensPerWETH;
        require(weth.balanceOf(address(this)) >= wethToReturn, "Contract lacks WETH");

        _burn(msg.sender, tokenAmount);
        require(weth.transfer(msg.sender, wethToReturn), "Transfer failed");
    }

    function setTokensPerWETH(uint256 newRate) external onlyOwner {
        require(newRate > 0, "Rate must be > 0");
        tokensPerWETH = newRate;
    }
}