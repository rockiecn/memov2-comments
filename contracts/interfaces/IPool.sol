// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// receive 
interface IPool {
    event Inflow(address indexed from, address tAddr, uint256 money);
    event Outflow(address indexed to, address tAddr, uint256 money);
    
    function inflow(address tAddr, address from, uint256 money) external payable;
    function outflow(address tAddr, address to, uint256 money) external payable;
}