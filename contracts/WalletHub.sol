// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract WalletHub is Ownable {
    using SafeERC20 for IERC20;

    IERC20 public credit;
    address public treasury;
    mapping(address => uint256) public balances;
    mapping(address => bool) public pools; 

    event Deposit(address indexed user, uint256 amount);
    event Debit(address indexed user, uint256 amount, address indexed pool);
    event Credit(address indexed recipient, uint256 amount, address indexed pool);

    constructor(address credit_, address treasury_) Ownable(msg.sender) {
        credit = IERC20(credit_);
        treasury = treasury_;
    }

    modifier onlyPool() {
        require(pools[msg.sender], "Not authorized pool");
        _;
    }

    function addPool(address pool) external onlyOwner {
        pools[pool] = true;
    }

    function removePool(address pool) external onlyOwner {
        pools[pool] = false;
    }

    function deposit(uint256 amount) external {
        credit.safeTransferFrom(msg.sender, address(this), amount);
        balances[msg.sender] += amount;
        emit Deposit(msg.sender, amount);
    }


    function debit(address user, uint256 amount) external onlyPool {
        require(balances[user] >= amount, "Insufficient user balance");
        balances[user] -= amount;
        emit Debit(user, amount, msg.sender);
    }

    function pay(address to, uint256 amount) external onlyPool {
        if (to != address(0)) {
            credit.safeTransfer(to, amount);
            emit Credit(to, amount, msg.sender);
        }
    }

    function treasuryTransfer(uint256 amount) external onlyPool {
        credit.safeTransfer(treasury, amount);
    }
    function setTreasury(address newTreasury) external onlyOwner {
        require(newTreasury != address(0), "zero treasury");
        treasury = newTreasury;
    }
    function userBalance(address user) public view returns (uint256) {
    return balances[user];
    }
}