// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IProxy.sol";
import "../interfaces/IControl1.sol";
import "../interfaces/IControl2.sol";
import "../interfaces/IInstance.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Operate Interface in the memo system.
 */
contract Proxy is IProxy, Owner {
    uint16 public version = 2;
    address public inst;

    // gas: 2791349
    constructor(address _a, address _inst) Owner(_a) {
        inst = _inst;
    }

    receive() external payable {}

    // gas: 196592
    function activate(uint64 _i, bytes[5] memory signs) external override {
        IControl1(IInstance(inst).instances(101)).activate(_i, signs);
        emit Activate(_i);
    }

    // gas: 83258
    function ban(uint64 _i, bool _ban, bytes[5] memory signs) external override {
        IControl1(IInstance(inst).instances(101)).ban(_i, _ban, signs);
        emit Ban(_i, _ban);
    }

    // gas: 153354
    function addT(address _t, bool _mint, bytes[5] memory signs) external override {
        IControl1(IInstance(inst).instances(101)).addT(_t, _mint, signs);
        emit AddT(_t, _mint);
    }

    // gas: 80220
    function banT(address _t, bool _isBan, bytes[5] memory signs) external override {
        IControl1(IInstance(inst).instances(101)).banT(_t, _isBan, signs);
        emit BanT(_t, _isBan);
    }

    // gas: 80776
    function banG(uint64 _gi, bool _ban, bytes[5] memory signs) external override {
        IControl1(IInstance(inst).instances(101)).banG(_gi, _ban, signs);
        emit BanG(_gi, _ban);
    }

    // gas: 2133781
    function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) external override {
        IControl1(IInstance(inst).instances(101)).createGroup(_level, _mr, _dc, _pc, _kr, _pr);
    }

    // gas: 128450
    function reAcc() external override {
        IControl1(IInstance(inst).instances(101)).reAcc(msg.sender);
    }
    
    // gas: 90703
    function reRole(uint8 _rtype, bytes memory _extra) external override {
        IControl1(IInstance(inst).instances(101)).reRole(msg.sender, _rtype, _extra);
    }

    // gas: 122224(keeper)/91054(user)
    function addToGroup(uint64 _gi) external override {
        IControl1(IInstance(inst).instances(101)).addToGroup(msg.sender, _gi);
    }
    
    // gas: 203609
    function pledge(uint64 _i, uint256 _money) external override {
        IControl1(IInstance(inst).instances(101)).pledge(msg.sender, _i, _money);
    }

    // gas: 180084
    function unpledge(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl1(IInstance(inst).instances(101)).unpledge(msg.sender, _i, _ti, _money);
    }

    // gas: 87262
    function alterPayee(address _p) external override {
        IControl1(IInstance(inst).instances(101)).alterPayee(msg.sender, _p);
    }

    // gas: 60295
    function confirmPayee(uint64 _i) external override {
        IControl1(IInstance(inst).instances(101)).confirmPayee(msg.sender, _i);
    }

    // called by user
    // gas: 399046、210710
    function addOrder(OrderIn memory _oi, bytes memory uSign, bytes memory pSign) external override {
        IControl2(IInstance(inst).instances(102)).addOrder(msg.sender, _oi, uSign, pSign);
    }

    // called by user or keeper
    // gas: 213248
    function subOrder(OrderIn memory _oi, bytes memory uSign, bytes memory pSign) external override {
        IControl2(IInstance(inst).instances(102)).subOrder(msg.sender, _oi, uSign, pSign);
    }

    // gas: 150257
    function addRepair(OrderIn memory _oi, uint64[] memory _kis, bytes[] memory ksigns) external override {
        IControl2(IInstance(inst).instances(102)).addRepair(msg.sender, _oi, _kis, ksigns);
    }

    // gas: 170922
    function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) external override {
       IControl2(IInstance(inst).instances(102)).recharge(msg.sender, _i, _ti, isLock, _money);
    }

    // gas: 62108(p)、143506(k)
    function withdraw(uint64 _i, uint8 _ti, uint256 _money) external override {
        IControl2(IInstance(inst).instances(102)).withdraw(msg.sender, _i, _ti, _money);
    }

    // gas: 204825
    function proWithdraw(PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external override {
        IControl2(IInstance(inst).instances(102)).proWithdraw(msg.sender, _ps, _kis, ksigns);
    }
}