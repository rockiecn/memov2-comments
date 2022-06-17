// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IControl1.sol";
import "../interfaces/IRole.sol";
import "../interfaces/IGroup.sol";
import "../interfaces/IToken.sol";
import "../interfaces/IAuth.sol";
import "../interfaces/IPledge.sol";
import "../interfaces/IPool.sol";
import "../interfaces/IKmanage.sol";
import "../interfaces/IInstance.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title controlling Role, Pledge, Token, PoolP, Kmanage in memo system
 */
contract Control1 is IControl1, Owner {
    

    uint16 public version = 2;
    address public inst;

    // deployed by admin
    // gas: 4200255
    constructor(address _a, address _inst) Owner(_a) {
        inst = _inst;
    }

    // ---------admin-----------

    // add erc-token
    function addT(address _t, bool _mint, bytes[5] memory signs) external override {
        bytes32 h = keccak256(abi.encodePacked(address(this), "addT", _t, _mint));
        IInstance insti = IInstance(inst);
        IAuth(auth).perm(h, signs);
   
        uint8 ti = ITokenSetter(insti.instances(7)).addT(_t, _mint);
        IPledgeSetter(insti.instances(8)).addT(ti);
    }

    // ban or unban token
    function banT(address _t, bool _ban, bytes[5] memory signs) external override {
        bytes32 h = keccak256(abi.encodePacked(address(this), "banT", _t, _ban));
        IAuth(auth).perm(h, signs);

        ITokenSetter(IInstance(inst).instances(7)).banT(_t, _ban);
    }

    // ban or unban account
    function ban(uint64 _i, bool _ban, bytes[5] memory signs) external override {
        bytes32 h = keccak256(abi.encodePacked(address(this), "ban", _i, _ban));
        IAuth(auth).perm(h, signs);

        IRoleSetter(IInstance(inst).instances(6)).ban(_i, _ban);
    }

    // ban or unban group
    function banG(uint64 _gi, bool _isBan, bytes[5] memory signs) external override {
        bytes32 h = keccak256(abi.encodePacked(address(this), "banG", _gi, _isBan));
        IAuth(auth).perm(h, signs);

        IGroupSetter(IInstance(inst).instances(11)).ban(_gi, _isBan);
    } 

    // activate keeper
    function activate(uint64 _i, bytes[5] memory signs) external override {
        bytes32 h = keccak256(abi.encodePacked(address(this), "activate", _i));
        IAuth(auth).perm(h, signs);

        IInstance insti = IInstance(inst);
        uint64 _gi = IRoleSetter(insti.instances(6)).activate(_i);

        IGroup g = IGroup(insti.instances(11));
        
        IKmanage ikm =  IKmanage(g.getKManage(_gi));
        ikm.addKeeper(_i); // judge here
        
        uint8 kcnt = ikm.getKCnt();
        g.activate(_gi, kcnt);
    }

    // ----------anyone---------

    // anyone can create group, then need admin to unban this group
    function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _k, uint256 _p) external override {
        require(_mr > 1 && _mr % 4 == 0, "managerate err");
        IGroupSetter(IInstance(inst).instances(11)).create(_level, _mr, _dc, _pc, _k, _p);
    }

    // --------proxy--------

    // get account index of _a
    function reAcc(address _a) external onlyOwner override {
        return IRoleSetter(IInstance(inst).instances(6)).reAcc(_a);
    }

    // _a roletype is set to be _rtype, extra is verified offline
    function reRole(address _a, uint8 _rtype, bytes memory extra) external onlyOwner override {
        IInstance insti = IInstance(inst);
        uint64 _i = IRoleGetter(insti.instances(6)).getAcc(_a);
        require(_i > 0, "reRole:not reAcc first"); // illegal account
        return IRoleSetter(insti.instances(6)).reRole(_i, _rtype, extra);
    }

    // _a is add to group _gi; need admin to activate keeper
    function addToGroup(address _a, uint64 _gi) external onlyOwner override {
        IInstance insti = IInstance(inst);
        RoleOut memory ro = IRoleGetter(insti.instances(6)).getRInfo(_a);
        require(ro.rType>0, "hasn't reRole");
        uint64 _i = IRoleGetter(insti.instances(6)).getAcc(_a);
        (,,uint64 _ogi, uint8 _rType) = IRoleGetter(insti.instances(6)).checkIG(_i, 0);
        require(_ogi==0,"addtoG:has in group"); //group exists

        uint256 bal = IPledgeGetter(insti.instances(8)).balanceOf(_i, 0);
        IGroupGetter(insti.instances(11)).add(_rType, _gi, bal); // check bal > pledge
        IRoleSetter(insti.instances(6)).setG(_i, _gi);
    }

    // anyone can pledge using its money, use a's money, pledge for i
    function pledge(address _a ,uint64 _i, uint256 money) external onlyOwner override {
        IInstance insti = IInstance(inst);
        (address _t,) = ITokenGetter(insti.instances(7)).checkT(0);

        // cal first then transfer
        IPledgeSetter(insti.instances(8)).pledge(_i, money);
        IPool(insti.instances(5)).inflow(_t, _a, money);

        emit Pledge(_a, _i, money);
    }

    // cancle pledge; unpledge only by payee
    function unpledge(address _a, uint64 _i, uint8 _ti, uint256 money) external onlyOwner override {
        IInstance insti = IInstance(inst);
        (address _t,) = ITokenGetter(insti.instances(7)).checkT(_ti);

        (, address _re, uint64 _gi, uint8 _rType) = IRoleGetter(insti.instances(6)).checkIG(_i, 0);
        require(_re == _a, "unpledge:ill caller");

        uint256 _lock;
       
        if (_ti == 0 && _gi > 0) {
            (uint256 _kp, uint256 _pp) = IGroupGetter(insti.instances(11)).getPInfo(_gi);
            if (_rType == 2) {
                _lock = _pp;
            }

            if (_rType == 3) {
                _lock = _kp;
            }
        }

        // todo: need lock more from fs

        uint256 _money = IPledgeSetter(insti.instances(8)).withdraw(_i, _ti, money, _lock);
        IPool(insti.instances(5)).outflow(_t, _re, _money);

        emit Unpledge(_a, _i, _ti, _money, _lock);
    }

    // old payee propose alter to new one; then new confirm it
    function alterPayee(address _a, address _p) external onlyOwner override {
        IInstance insti = IInstance(inst);
        uint64 _i = IRoleGetter(insti.instances(6)).getAcc(_a);
        require(_i > 0, "alterPayee:ill acc"); // illegal account
        (,address _payee,,) = IRoleGetter(insti.instances(6)).checkIG(_i, 0);
        require(_a == _payee, "payee has changed"); // illegal caller
        return IRoleSetter(insti.instances(6)).alterPayee(_i, _p);
    }

    // new payee confirm
    function confirmPayee(address _a, uint64 _i) external onlyOwner override {
        return IRoleSetter(IInstance(inst).instances(6)).confirmPayee(_i, _a);
    }
}