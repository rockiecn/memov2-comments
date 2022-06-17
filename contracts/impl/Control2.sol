// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IControl2.sol";
import "../interfaces/IRoleGetter.sol";
import "../interfaces/IGroupGetter.sol";
import "../interfaces/IFileSysSetter.sol";
import "../interfaces/IIssuance.sol";
import "../interfaces/ITokenGetter.sol";
import "../interfaces/IERC20.sol";
import "../interfaces/IPool.sol";
import "../interfaces/IKmanage.sol";
import "../interfaces/IInstance.sol";
import "./Owner.sol";
import "../Recover.sol";

/**
 *@author MemoLabs
 *@title controlling FileSys, Issue, erc20, PoolF, Kmanage in memo system
 */
contract Control2 is IControl2, Owner {
    using Recover for bytes32;
    uint16 public version = 2;
    address public inst;

    // deployed by admin
    // gas: 5305310
    constructor(address _a, address _inst) Owner(_a) {
        inst = _inst;
    }

    receive() external payable {}

    function checkOrder(OrderIn memory ps) internal view returns (bytes32) {
        require(ps.size > 0, "order size zero"); // size zero
        require(ps.sprice > 0, "order sprice zero" ); // sizePrice zero
        require(ps.end <= ps.start + 86400000, "order dur more than 1000d"); // end no more than start+1000 day
        require(ps.end%86400 == 0, "order end not align to d"); // end time error; align to day
        require(ps.start < block.timestamp, "order start more than now"); // start is less than now

        bytes32 h = keccak256(abi.encodePacked(
            ps.uIndex,
            ps.pIndex,
            ps.nonce,
            ps.start,
            ps.end,
            ps.size, 
            ps.tIndex, 
            ps.sprice)
        );
        return h;
    }

    function checkParam(OrderIn memory ps, bytes memory uSign, bytes memory pSign, bool sub) internal view returns (address, uint64) {
        // suborder can be shorter than 100days
        if (!sub) {
            require(ps.end >= ps.start + 8640000, "order dur less 100d"); // end more than start+100 day
        }

        bytes32 h = checkOrder(ps);

        IInstance insti = IInstance(inst);

        // verify provider sig
        (address _a,,uint64 _gi,) = IRoleGetter(insti.instances(6)).checkIG(ps.pIndex, 2);
        require(h.recover(pSign) == _a, "pro sign err"); // illegal pro sign

        uint64 _ugi;
        (_a,,_ugi,) = IRoleGetter(insti.instances(6)).checkIG(ps.uIndex, 1);

        // verify user sig
        require(h.recover(uSign) == _a, "user sign err"); // illegal user sign

        require(_gi == _ugi, "group differ"); // group different

        return (_a, _gi);
    }

    function getk(uint64 _gi) internal view returns (address) {
        IInstance insti = IInstance(inst);
        IGroupGetter g = IGroupGetter(insti.instances(11));
        g.checkG(_gi);
        return g.getKManage(_gi);
    }

    function checkSign(uint64 _gi ,bytes32 h ,uint64[] memory kIndexes, bytes[] memory ksigns) internal view returns (address)  {
        uint8 sigCnt;
        address _w;
        uint64 gIndex;
        IInstance insti = IInstance(inst);
        IRoleGetter r = IRoleGetter(insti.instances(6));
        for(uint64 i = 0; i < kIndexes.length; i++) {
            if (i != 0){
                require(kIndexes[i] > kIndexes[i-1], "kindexes err"); // larger than previous
            }
            (_w,, gIndex,) = r.checkIG(kIndexes[i], 3);
            if (gIndex == _gi && h.recover(ksigns[i]) == _w) {
                sigCnt += 1;
            }      
        }

        IGroupGetter g = IGroupGetter(insti.instances(11));
        g.checkG(_gi);

        _w = g.getKManage(_gi);

        // valid sig should not less than 2*(N+1)/3, N: kNum of group
        require(sigCnt >= g.getLevel(_gi), "ksign cnt less glevel");
        require(sigCnt >= 2 * (IKmanage(_w).getKCnt() + 1) / 3, "ksign cnt less 2/3"); // kSigns error

        return _w;
    }

    // --------proxy--------

    // use a's money, recharge for ui
    function recharge(address _a, uint64 _ui, uint8 _ti, bool isLock, uint256 money) external onlyOwner override {
        IInstance insti = IInstance(inst);
        (address _t,) = ITokenGetter(insti.instances(7)).checkT(_ti);
        
        IRoleGetter(insti.instances(6)).checkIG(_ui, 0);
        
        IPool(insti.instances(12)).inflow(_t, _a, money);
        
        IFileSysSetter(insti.instances(10)).recharge(_ui, _ti, isLock, money);

        emit Recharge(_a, _ui, isLock, _ti, money);
    }

    // called by Proxy.sol, tx.origin = a, need a = addr[i].payee;
    // foundation、user、keeper、provider(should call proWithdraw first) withdraw
    function withdraw(address _a, uint64 _i, uint8 _ti, uint256 money) external onlyOwner override {
        IInstance insti = IInstance(inst);
        (address _t,) = ITokenGetter(insti.instances(7)).checkT(_ti);
 
        (, address _re, uint64 _gi, ) = IRoleGetter(insti.instances(6)).checkIG(_i, 0);
        require(_a == _re, "withdraw:ill caller");

        uint256 _money = IFileSysSetter(insti.instances(10)).withdraw(_i, _ti, money); // reduce balances[i] in fs
        if (money > _money && _gi > 0) { // foundation's group is zero
            money -= _money;
            uint256 _m = IKmanage(getk(_gi)).withdraw(_i, _ti, money);
            _money += _m;
        } 
        IPool(insti.instances(12)).outflow(_t, _re, _money);

        emit Withdraw(_a, _i, _ti, money, _money);
    }

    // kIndexes is incremental, a = tx.origin, a should be ps.pIndex or ps.pIndex.payee
    function proWithdraw(address _a, PWIn memory ps, uint64[] memory kIndexes, bytes[] memory ksigns) external onlyOwner override {
        IInstance insti = IInstance(inst);
        ITokenGetter(insti.instances(7)).checkT(ps.tIndex);

        (address _w, address _re, uint64 _gi, ) = IRoleGetter(insti.instances(6)).checkIG(ps.pIndex, 2);
        require((_a == _w || _a == _re), "pwithdraw:ill caller"); // due to repair
  
        bytes32 h = keccak256(abi.encodePacked(ps.pIndex,ps.tIndex,ps.pay,ps.lost));
        _re = checkSign(_gi, h, kIndexes, ksigns);

        uint256 _money = IFileSysSetter(insti.instances(10)).proWithdraw(ps, _gi); // added to balance, not transfer out

        // linear pay to keepers
        uint256 _pr = _money * IKmanage(_re).manageRate() * 3 / 400; // 75% managePay for linearPaid
        IKmanage(_re).addProfit(ps.tIndex, _pr);

        emit ProWithdraw(_a, ps.pIndex, ps.tIndex, ps.pay, ps.lost);
    }

    // order duration >= 100 days & <= 1000 days, a = tx.origin, a is user
    function addOrder(address _a, OrderIn memory ps, bytes memory uSign, bytes memory pSign) external onlyOwner override {
        IInstance insti = IInstance(inst);
        (address _t, bool _isMint) = ITokenGetter(insti.instances(7)).checkT(ps.tIndex);
        
        (address _u, uint64 _gi) = checkParam(ps, uSign, pSign, false);
        require(_a == _u, "aOrder:ill caller"); // illegal caller

        uint8 _mr = IKmanage(getk(_gi)).manageRate();
        IFileSysSetter(insti.instances(10)).addOrder(ps, _mr, _gi);   
        
        // 产生奖励
        if (_isMint) {
            uint256 reward = IIssuance(insti.instances(9)).issu(ps.start, ps.end, ps.size, ps.sprice);
            if(reward>0){
                IERC20(_t).mint(insti.instances(5), reward); // need control2 has mint-access
            }
        } 
    }

    // _a = tx.origin; _a need to be ps.uaddr, or kaddr in same group as ps.uaddr
    function subOrder(address _a, OrderIn memory ps, bytes memory uSign, bytes memory pSign) external onlyOwner override {
        IInstance insti = IInstance(inst);
        ITokenGetter(insti.instances(7)).checkT(ps.tIndex);

        (address uAddr, uint64 _gi) = checkParam(ps, uSign, pSign, true);
        require(ps.end <= block.timestamp, "sOrder:end more than now"); // time error

        IFileSysSetter(insti.instances(10)).subOrder(ps, _gi);

        uint64 kIndex;
        if (_a != uAddr) {
            IRoleGetter r = IRoleGetter(insti.instances(6));
            kIndex = r.getAcc(_a);
            // not user, should be keeper
            (,,uint64 _ngi,) = r.checkIG(kIndex, 3);
            require(_gi == _ngi, "sOrder:g differ");
            
            IKmanage ikm =  IKmanage(getk(_gi));
            uint256 ep = ps.sprice * (ps.end-ps.start) * ikm.manageRate() / 400;

            ikm.addCnt(kIndex, 1);
            ikm.addProfit(ps.tIndex, ep);
        } 
    }

    // call by provider; signed by keepers
    // not charge manage and tax fee, not mint
    function addRepair(address _a, OrderIn memory ps, uint64[] memory kIndexes, bytes[] memory ksigns) external onlyOwner override {
        IInstance insti = IInstance(inst);
        ITokenGetter(insti.instances(7)).checkT(ps.tIndex);

        require(ps.uIndex == 0, "aRepair:ui not zero");

        (address _p,,uint64 _gi,) = IRoleGetter(insti.instances(6)).checkIG(ps.pIndex, 2);
        require(_a == _p, "aRepair:ill caller"); // illegal caller

        bytes32 h = checkOrder(ps);
        checkSign(_gi, h, kIndexes, ksigns);

        IFileSysSetter(insti.instances(10)).addRepair(ps, _gi);          
    }
}