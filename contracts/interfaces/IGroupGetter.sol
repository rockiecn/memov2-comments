// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IGroupGetter {
    function add(uint8 _rType, uint64 _gi, uint256 _pm) external view;
    function checkG(uint64 i) external view;
    
    function getGCnt() external view returns (uint64);
    function getGInfo(uint64 i) external view returns (bool, bool);
    function getLevel(uint64 i) external view returns (uint8);
    function getPInfo(uint64 _i) external view returns (uint256, uint256);

    function getKManage(uint64 _i) external view returns (address);
    function getSStra(uint64 i) external view returns (uint8, uint8);
}