//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IRole.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Manage account, roles in the memo system.
 */
contract Role is IRole, Owner {
    struct RoleInfo {
        bool isBan;     // 是否被admin禁止
        bool isActive;  // true when add to group; keeper need set by admin;
        uint8 rType;    // 0 account; 1 user; 2 provider; 3 keeper
        uint64 index;   // 账户序列号
        uint64 gIndex;  // 所属group
        address payee;  // acoount for receiving income
        address next;   // next payee
        bytes extra;    // 0 empty; 1 pdp verify key; 2 ; 3 bls public key
    }

    uint16 public version = 2;

    address[] accounts; // all roles 序列号即为index,从0开始, accounts[0]为foundation地址
    mapping(address => RoleInfo) info;

    // gas: 3359012
    constructor(address _a, address f) Owner(_a) {
        accounts.push(f);
        info[f].payee = f;
        emit ReAcc(f, 0);
    }

    function ban(uint64 _index, bool _banned) external onlyOwner override {
        address a = accounts[_index];
        info[a].isBan = _banned;
    }

    function setG(uint64 _index, uint64 _gi) external onlyOwner override {
        address a = accounts[_index];
        info[a].gIndex = _gi;
        if (info[a].rType != 3) {
            info[a].isActive = true;
        }
        emit SetG(_gi, _index);
    }

    function alterPayee(uint64 _index, address _p) external onlyOwner override {
        address a = accounts[_index];
        info[a].next = _p;
        
        emit AlterPayee(_p, a);
    }

    function confirmPayee(uint64 _index, address _p) external onlyOwner override {
        address a = accounts[_index];
        require(!info[a].isBan && info[a].next == _p, "caller err");
        info[a].payee = _p;

        emit ConfirmPayee(_p, a);
    }

    function activate(uint64 _index) external onlyOwner override returns (uint64) {
        address a = accounts[_index];
        // is keeper, not active, not ban, and been in some group
        require(!info[a].isActive && !info[a].isBan && info[a].rType == 3 && info[a].gIndex > 0, "acc err"); // account error
        
        info[a].isActive = true;

        return info[a].gIndex;
    }

    function reAcc(address a) external onlyOwner override {
        require(info[a].index == 0 && info[a].gIndex == 0, "has reAcc");
        
        uint64 len = uint64(accounts.length);
        info[a].index = len;
        info[a].payee = a;
        accounts.push(a);
        emit ReAcc(a, len);
    }

    function reRole(uint64 _index, uint8 _rType, bytes memory _extra) external onlyOwner override {
        address a = accounts[_index];
        require (_rType > 0, "has reRole"); // illegal type
        require(info[a].rType==0 && info[a].gIndex == 0 && !info[a].isBan && !info[a].isActive, "AE");  // role is not active, not in some group, not banned, empty type

        info[a].rType = _rType;
        info[a].extra = _extra;
        emit ReRole(_rType, _index);
    }

    // check and get role info
    function checkIG(uint64 _index, uint8 _rType) external view override returns (address, address, uint64, uint8) {
        address a = accounts[_index];
        require(!info[a].isBan, "acc banned"); // account error

        if (_rType > 0) {
            require(info[a].isActive && info[a].rType==_rType, "acc type err or not active"); // Type error
        }
        
        return (a,info[a].payee, info[a].gIndex, info[a].rType);
    }

    // ===================get===================
    function getACnt() external view override returns (uint64) {
        return uint64(accounts.length);
    }

    function getAddr(uint64 _i) external view override returns (address) {
        return accounts[_i];
    }

    function getAcc(address _a) external view override returns (uint64) {
        return info[_a].index;
    }

    function getRInfo(address acc) external view override returns (RoleOut memory) {
        RoleOut memory r;
        r = RoleOut(info[acc].isBan, info[acc].isActive, info[acc].rType, info[acc].index, info[acc].gIndex, info[acc].payee, info[acc].next, info[acc].extra);
        return r;
    }
}