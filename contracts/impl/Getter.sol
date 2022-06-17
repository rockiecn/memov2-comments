// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./Owner.sol";
import "../interfaces/ITokenGetter.sol";
import "../interfaces/IFileSysGetter.sol";
import "../interfaces/IGroupGetter.sol";
import "../interfaces/IIssuance.sol";
import "../interfaces/IRoleGetter.sol";
import "../interfaces/IPool.sol";
import "../interfaces/IPledge.sol";
import "../interfaces/IKmanageGetter.sol";
import "../interfaces/IInstance.sol";
import "../interfaces/IGetter.sol";

/**
 *@author MemoLabs
 *@title get interface in memo system
 */

contract Getter is IGetter, Owner {
    uint16 public version = 2;
    address public inst; // instance-contract address

    // gas: 830484
    constructor(address _a,address _inst) Owner(_a) {
        inst = _inst;
    }

    // ----------Token----------
    function getTA(uint8 tIndex) external view override returns (address) {
        IInstance insti = IInstance(inst);
        return ITokenGetter(insti.instances(7)).getTA(tIndex);
    }
    function getTI(address t) external view override returns (uint8, bool, bool) {
        IInstance insti = IInstance(inst);
        return ITokenGetter(insti.instances(7)).getTI(t);
    }
    function getTCnt() external view override returns (uint8) {
        IInstance insti = IInstance(inst);
        return ITokenGetter(insti.instances(7)).getTCnt();
    }

    // -----------Filesys-----------
    function fsBalanceOf(uint64 i, uint8 ti) external view override returns(uint256, uint256) {
        IInstance insti = IInstance(inst);
        return IFileSysGetter(insti.instances(10)).balanceOf(i, ti);
    }
    function getGSInfo(uint64 gi, uint8 ti) external view override returns (GroupOut memory) {
        IInstance insti = IInstance(inst);
        return IFileSysGetter(insti.instances(10)).getGInfo(gi, ti);
    }
    function getFsInfo(uint64 ui, uint64 pi) external view override returns (FsOut memory){
        IInstance insti = IInstance(inst);
        return IFileSysGetter(insti.instances(10)).getFsInfo(ui, pi);
    }
    function getStoreInfo(uint64 ui, uint64 pi, uint8 ti) external view override returns (StoreOut memory){
        IInstance insti = IInstance(inst);
        return IFileSysGetter(insti.instances(10)).getStoreInfo(ui, pi, ti);
    }
    function getSettleInfo(uint64 pi, uint8 ti) external view override returns (SettleOut memory){
        IInstance insti = IInstance(inst);
        return IFileSysGetter(insti.instances(10)).getSettleInfo(pi, ti);
    }

    // ------------Group----------
    function getGCnt() external view override returns (uint64) {
        IInstance insti = IInstance(inst);
        return IGroupGetter(insti.instances(11)).getGCnt();
    }
    function getGInfo(uint64 i) external view override returns (bool, bool) {
        IInstance insti = IInstance(inst);
        return IGroupGetter(insti.instances(11)).getGInfo(i);
    }

    function getLevel(uint64 i) external view override returns (uint8) {
        IInstance insti = IInstance(inst);
        return IGroupGetter(insti.instances(11)).getLevel(i);
    }
    function getPInfo(uint64 i) external view override returns (uint256, uint256) {
        IInstance insti = IInstance(inst);
        return IGroupGetter(insti.instances(11)).getPInfo(i);
    }
    function getKManage(uint64 i) external view override returns (address) {
        IInstance insti = IInstance(inst);
        return IGroupGetter(insti.instances(11)).getKManage(i);
    }
    function getSStra(uint64 i) external view override returns (uint8, uint8) {
        IInstance insti = IInstance(inst);
        return IGroupGetter(insti.instances(11)).getSStra(i);
    }

    // ----------Issuance------------
    

    // ---------kmanage----------
    function manageRate(uint64 gi) external view override returns (uint8) {
        IInstance insti = IInstance(inst);
        address km =IGroupGetter(insti.instances(11)).getKManage(gi);
        return IKmanageGetter(km).manageRate();
    }
    function getKCnt(uint64 gi) external view override returns (uint8) {
        IInstance insti = IInstance(inst);
        address km = IGroupGetter(insti.instances(11)).getKManage(gi);
        return IKmanageGetter(km).getKCnt();
    }
    function getK(uint64 i, uint64 gi) external view override returns (uint64) {
        IInstance insti = IInstance(inst);
        address km = IGroupGetter(insti.instances(11)).getKManage(gi);
        return IKmanageGetter(km).getK(i);
    }
    function getPf(uint8 ti, uint64 gi) external view override returns (uint64, uint256) {
        IInstance insti = IInstance(inst);
        address km = IGroupGetter(insti.instances(11)).getKManage(gi);
        return IKmanageGetter(km).getPf(ti);
    } 
    function kmBalanceOf(uint64 ki, uint8 ti, uint64 gi) external view override returns(uint256, uint256) {
        IInstance insti = IInstance(inst);
        address km = IGroupGetter(insti.instances(11)).getKManage(gi);
        return IKmanageGetter(km).balanceOf(ki, ti);
    }

    // -----------owner---------
    function instances(uint8 _type) external view override returns(address) {
        IInstance insti = IInstance(inst);
        return insti.instances(_type);
    }

    // -----------pledge--------
    function pleBalanceOf(uint64 i, uint8 ti) external view override returns (uint256) {
        IInstance insti = IInstance(inst);
        IPledge ple = IPledge(insti.instances(8));
        return ple.balanceOf(i, ti);
    }
    function getPledge(uint8 ti) external view override returns (uint256) {
        IInstance insti = IInstance(inst);
        return IPledge(insti.instances(8)).getPledge(ti);
    }
    function getTotalPledge() external view override returns (uint256) {
        IInstance insti = IInstance(inst);
        return IPledge(insti.instances(8)).totalPledge();
    }

    // ----------pool---------


    // ----------role----------
    function getAcc(address a) external view override returns (uint64) {
        IInstance insti = IInstance(inst);
        return IRoleGetter(insti.instances(6)).getAcc(a);
    }
    function getAddr(uint64 i) external view override returns (address) {
        IInstance insti = IInstance(inst);
        return IRoleGetter(insti.instances(6)).getAddr(i);
    }
    function getACnt() external view override returns (uint64) {
        IInstance insti = IInstance(inst);
        return IRoleGetter(insti.instances(6)).getACnt();
    }
    function getRInfo(address a) external view override returns (RoleOut memory) {
        IInstance insti = IInstance(inst);
        return IRoleGetter(insti.instances(6)).getRInfo(a);
    }
}