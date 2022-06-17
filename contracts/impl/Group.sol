//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IGroup.sol";
import "../interfaces/IInstance.sol";
import "./Owner.sol";
import "./Kmanage.sol";

/**
 *@author MemoLabs
 *@title Manage groups in the memo system.
 */
contract Group is IGroup, Owner {
    struct GroupInfo {
        bool isActive;      // true when keeper count >= level 
        bool isBan;         // 组是否已被禁用
        uint8 level;       // security level
        uint8 dc;           // storage strategy: data count
        uint8 pc;           // storage strategy: parity count
        address kManage;    // keeper manage contract address
        uint256 kpr;        // keeper pledge require 
        uint256 ppr;        // provider pledge require
    }

    uint16 public version = 2;

    address public inst;

    GroupInfo[] groups; // manage group

    // gas: 4299143
    constructor(address _a, address _inst) Owner(_a) {
        inst = _inst;
        GroupInfo memory g;
        g.isBan = true;  // group 0 is banned
        groups.push(g);
    }

    function ban(uint64 _gi, bool _isBan) external onlyOwner override {
        groups[_gi].isBan = _isBan;
    }

    function activate(uint64 _gi, uint8 _kc) external onlyOwner override {
        require(!groups[_gi].isBan, "act:g banned");
        if (_kc >= groups[_gi].level) {
            groups[_gi].isActive = true;
        }
    }

    function create(uint8 _level, uint8 mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) external onlyOwner override {
        uint64 _gIndex = uint64(groups.length);

        // create kmanage address; force each group has unique kmanage 
        IInstance insti = IInstance(inst); 
        Kmanage k = new Kmanage(insti.instances(2), mr);
        
        GroupInfo memory g;
        g.level = _level;
        g.kpr = _kr;
        g.ppr = _pr; 
        g.dc = _dc;
        g.pc = _pc;
        g.kManage = address(k);
        g.isBan = true;  // banned at first
        groups.push(g);

        emit Create(_gIndex, _level, mr, _dc, _pc, _kr, _pr);
    }

    function add(uint8 _rType, uint64 _gi, uint256 _pm) external override view {
        require(!groups[_gi].isBan, "add:g banned"); // group is banned
    
        if (_rType == 1) {
            require(groups[_gi].isActive, "addu:g not active");
            return;
        }

        if (_rType == 2) {
            require(groups[_gi].isActive, "addP:g not active");
            require(_pm >= groups[_gi].ppr, "p pledge not enough"); // pledge insuf
            return;
        }

        if (_rType == 3) {
            require(_pm >= groups[_gi].kpr, "k pledge not enough"); // pledge insuf
        }
    }

    // check
    function checkG(uint64 i) external view override {
        require(groups[i].isActive && !groups[i].isBan);
    }

    // ===================get===================

    function getGCnt() external view override returns (uint64) {
        return uint64(groups.length);
    }

    function getGInfo(uint64 i) external view override returns (bool, bool) {
        return (groups[i].isActive, groups[i].isBan);
    }

    function getLevel(uint64 i) external view override returns (uint8) {
        return groups[i].level;
    }

     function getPInfo(uint64 i) external view override returns (uint256, uint256) {
        return (groups[i].kpr, groups[i].ppr);
    }

    function getKManage(uint64 i) external view override returns (address) {
        return groups[i].kManage;
    }

    function getSStra(uint64 i) external view override returns (uint8, uint8) {
        return (groups[i].dc, groups[i].pc);
    }
}