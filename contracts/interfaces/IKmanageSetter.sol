// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IPool.sol";

interface IKmanageSetter {
    event AddProfit(uint8 indexed ti, uint256 money);
    event AddCnt(uint64 indexed ki, uint64 cnt);
    
    function addKeeper(uint64 _ki) external;
    function addCnt(uint64 _ki, uint64 _cnt) external;
    function addProfit(uint8 _ti, uint256 _money) external;
    function withdraw(uint64 _ki, uint8 _ti, uint256 money) external returns (uint256);
}