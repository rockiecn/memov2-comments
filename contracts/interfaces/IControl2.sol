// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IFileSysSetter.sol";

interface IControl2 {
    event Recharge(address indexed payer, uint64 indexed i, bool indexed isLock, uint8 ti, uint256 money);
    event Withdraw(address indexed addr, uint64 indexed i, uint8 ti, uint256 money, uint256 actualMoney);
    event ProWithdraw(address indexed a, uint64 indexed pi, uint8 ti, uint256 pay, uint256 lost);

    function addOrder(address _a, OrderIn memory _oi, bytes memory uSign, bytes memory pSign) external;
    function subOrder(address _a, OrderIn memory _oi, bytes memory uSign, bytes memory pSign) external;

    function addRepair(address _a, OrderIn memory _oi, uint64[] memory _kis, bytes[] memory ksigns) external;

    function recharge(address _a, uint64 _i, uint8 _ti, bool isLock, uint256 _money) external;
    function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) external;
    function proWithdraw(address _a, PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;
}