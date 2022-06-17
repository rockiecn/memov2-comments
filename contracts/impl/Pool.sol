//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IPool.sol";
import "../interfaces/IERC20.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Pool for receiving and sending tokens in the memo system.
 */
contract Pool is IPool, Owner {
    uint16 public version = 2;

    // gas: 966350
    constructor(address _a) Owner(_a) {
    }

    receive() external payable {}

    // from should approve address(this) before call 'inflow'
    function inflow(address tAddr, address from, uint256 money) external onlyOwner override  payable {
        IERC20(tAddr).transferFrom(from, address(this), money);
        emit Inflow(from, tAddr, money);
    }

    function outflow(address tAddr, address to, uint256 money) external onlyOwner override payable {
        IERC20(tAddr).transfer(to, money);
        emit Outflow(to, tAddr, money);
    }
}