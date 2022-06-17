// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IIssuance {
    function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) external returns (uint256);
}