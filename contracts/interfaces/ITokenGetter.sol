// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface ITokenGetter {
    function checkT(uint8 _ti) external view returns (address, bool);
    function getTA(uint8 _ti) external view returns (address);
    function getTI(address _t) external view returns (uint8, bool, bool);
    function getTCnt() external view returns (uint8);
}
