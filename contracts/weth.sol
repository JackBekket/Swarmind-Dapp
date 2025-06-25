// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract WETH is ERC20 {
    constructor() ERC20("Wrapped ETH", "WETH") {}

    receive() external payable {
        _mint(msg.sender, msg.value);
    }

    function unwrap(uint256 amount) external {
        require(balanceOf(msg.sender) >= amount, "Not enough WETH");
        _burn(msg.sender, amount);
        payable(msg.sender).transfer(amount);
    }

    function unwrapTo(uint256 amount, address to) external {
        require(to != address(0), "Invalid recipient");
        require(balanceOf(msg.sender) >= amount, "Not enough WETH");
        _burn(msg.sender, amount);
        payable(to).transfer(amount);
    }
}