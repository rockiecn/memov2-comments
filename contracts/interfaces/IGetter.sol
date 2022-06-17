// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IOwner.sol";
import "./ITokenGetter.sol";
import "./IFileSysGetter.sol";
import "./IGroupGetter.sol";
import "./IIssuance.sol";
import "./IRoleGetter.sol";
import "./IPool.sol";
import "./IPledge.sol";
import "./IKmanageGetter.sol";
import "./IInstance.sol";

interface IGetter {
    // ----------Token----------
    function getTA(uint8 tIndex) external view returns (address);
    function getTI(address t) external view returns (uint8, bool, bool);
    function getTCnt() external view returns (uint8);

    // -----------Filesys-----------
    function fsBalanceOf(uint64 i, uint8 ti) external view returns(uint256, uint256);
    function getGSInfo(uint64 gi, uint8 ti) external view returns (GroupOut memory);
    function getFsInfo(uint64 ui, uint64 pi) external view returns (FsOut memory);
    function getStoreInfo(uint64 ui, uint64 pi, uint8 ti) external view returns (StoreOut memory);
    function getSettleInfo(uint64 pi, uint8 ti) external view returns (SettleOut memory);

    // ------------Group----------
    function getGCnt() external view returns (uint64);
    function getGInfo(uint64 i) external view returns (bool, bool);
    function getLevel(uint64 i) external view returns (uint8);
    function getPInfo(uint64 i) external view returns (uint256, uint256);
    function getKManage(uint64 i) external view returns (address);
    function getSStra(uint64 i) external view returns (uint8, uint8);

    // ---------kmanage----------
    function manageRate(uint64 gi) external view returns (uint8);
    function getKCnt(uint64 gi) external view returns (uint8);
    function getK(uint64 i, uint64 gi) external view returns (uint64);
    function getPf(uint8 ti, uint64 gi) external view returns (uint64, uint256);
    function kmBalanceOf(uint64 ki, uint8 ti, uint64 gi) external view returns(uint256, uint256);

    // -----------owner---------
    function instances(uint8 _type) external view returns(address);

    // -----------pledge--------
    function pleBalanceOf(uint64 i, uint8 ti) external view returns (uint256);
    function getPledge(uint8 ti) external view returns (uint256);
    function getTotalPledge() external view returns (uint256);

    // ----------role----------
    function getAcc(address a) external view returns (uint64);
    function getAddr(uint64 i) external view returns (address);
    function getACnt() external view returns (uint64);
    function getRInfo(address a) external view returns (RoleOut memory);
}