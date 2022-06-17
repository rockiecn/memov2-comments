// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IFileSys.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title manage fs in memo system
 */
contract FileSys is IFileSys, Owner {

    struct StoreInfo {
        uint64 start;  // last start 
        uint64 end;    // 什么时刻的状态，last end time
        uint64 size;   // 在该存储节点上的存储总量，byte
        uint256 sprice; 
    }

    struct GroupInfo {
        uint64 size;   
        uint256 sprice;
        uint256 lost; 
    }

    struct AggOrder {
        uint64 nonce;    // 防止order重复提交
        uint64 subNonce; // 用于订单到期
        mapping(uint8 => StoreInfo) sInfo; // 不同代币的支付信息，tokenIndex => StoreInfo
    }

    struct FsInfo {
        mapping(uint64 => AggOrder) ao; // 该User对每个Provider的订单信息
    }

    // Settlement indicates billing information
    struct Settlement {
        uint64 time;    // store状态改变或支付的时间, 与 epoch 对齐
        uint64 size;    // 在该存储节点上的存储总量
        uint256 sprice; // 累积的sprice(即sizePrice)

        uint256 maxPay;  // 对此provider所有user聚合总额度； expected 加和
        uint256 hasPaid; // 已经支付
        uint256 canPay;  // 最近一次store/pay时刻，可以支付的金额
        uint256 lost;    // lost due to unable response to challenge
    }

    uint8 public constant taxRate = 1;

    uint16 public version = 2;

    mapping(uint64 => mapping(uint8 => GroupInfo)) groups;     // total lost per group, gi=>ti=>GroupInfo
    mapping(uint64 => mapping(uint8 => uint256)) balances; // avail per account
    mapping(uint64 => mapping(uint8 => uint256)) lock;     // cannot withdraw 

    mapping(uint64 => FsInfo) fs; // user => FsInfo; user 0 is repair fs

    mapping(uint64 => mapping(uint8 => Settlement)) proInfo; // pro => token => income

    // gas: 4598363
    constructor(address _a) Owner(_a) {}

    function _settlementAdd(uint64 _pIndex, uint8 _ti, uint64 start, uint64 size, uint256 sprice, uint256 pay) internal {
        // update canPay
        uint64 t = proInfo[_pIndex][_ti].time;
        if(t <= start){ 
            if(t!=0){ // 非首次addOrder
                proInfo[_pIndex][_ti].canPay += (start-t) * proInfo[_pIndex][_ti].sprice;
            }
            proInfo[_pIndex][_ti].time = start;
        }else{
            proInfo[_pIndex][_ti].canPay += uint256(t - start)*sprice;
        }

        proInfo[_pIndex][_ti].sprice += sprice;
        proInfo[_pIndex][_ti].size += size;
        proInfo[_pIndex][_ti].maxPay += pay; // update maxPay; hardlimit
    }

    function _settlementSub(uint64 _pIndex, uint8 _ti, uint64 end, uint64 size, uint256 sprice) internal {
        // update canPay
        uint64 t = proInfo[_pIndex][_ti].time;

        if(t <= end){
            proInfo[_pIndex][_ti].canPay += (end - t) * proInfo[_pIndex][_ti].sprice;
            proInfo[_pIndex][_ti].time = end;
        } else {
            // should sub it
            proInfo[_pIndex][_ti].canPay -= (t - end) * sprice;
        }

        // update size and price
        proInfo[_pIndex][_ti].sprice -= sprice;
        proInfo[_pIndex][_ti].size -= size;
    }

    function _settlementCal(uint64 _pIndex, uint8 _ti, uint256 pay, uint256 _lost) internal returns (uint256, uint256) {
        Settlement memory se = proInfo[_pIndex][_ti];
        // 'has paid', or 'lost' is not right
        if(se.hasPaid > pay || se.lost > _lost){
            return (0, 0);
        }

        // max pay is not enough
        if(se.maxPay - _lost < pay){
            return (0, 0);
        }

        proInfo[_pIndex][_ti].lost = _lost;
        _lost = _lost - se.lost;

        uint64 ntime = uint64(block.timestamp);
        if(se.time < ntime){
            proInfo[_pIndex][_ti].canPay += (ntime - se.time) * se.sprice;
            proInfo[_pIndex][_ti].time = ntime;
        }

        // can pay is not right
        if(proInfo[_pIndex][_ti].canPay<pay){
            return (0,0);
        }

        uint256 res = pay - se.hasPaid;
        proInfo[_pIndex][_ti].hasPaid = pay;
        return (res, _lost);
    }

    function addOrder(OrderIn memory ps, uint256 mr, uint64 _gi) external override onlyOwner {
        uint256 pay = (ps.end-ps.start) * ps.sprice;
        uint256 manage = pay / 100 * uint256(mr);
        uint256 tax = pay / 100 * uint256(taxRate);
        uint256 payAndTax = pay + manage + tax;
        require(balances[ps.uIndex][ps.tIndex] + lock[ps.uIndex][ps.tIndex] >= payAndTax, "u fs-bal not enough"); // balance not enough

        // 验证nonce
        require(fs[ps.uIndex].ao[ps.pIndex].nonce == ps.nonce, "nonce err"); // nonce error
        // start不减, end不减
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start <= ps.start, "start err"); // start error, start shouldn't less than last order's start
        require(fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end <= ps.end, "end err"); // end error, end shouldn't less than last order's end

        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].sprice += ps.sprice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size += ps.size;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].start = ps.start;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].end = ps.end;

        _settlementAdd(ps.pIndex, ps.tIndex, ps.start, ps.size, ps.sprice, pay);

        fs[ps.uIndex].ao[ps.pIndex].nonce++;

        groups[_gi][ps.tIndex].size += ps.size;
        groups[_gi][ps.tIndex].sprice += ps.sprice;

        // add to foundation
        balances[0][ps.tIndex] += tax;

        // pay from lock first
        if (lock[ps.uIndex][ps.tIndex] >= payAndTax) {
            lock[ps.uIndex][ps.tIndex] -= payAndTax;
        } else {
            payAndTax -= lock[ps.uIndex][ps.tIndex]; 
            lock[ps.uIndex][ps.tIndex] = 0;
            balances[ps.uIndex][ps.tIndex] -= payAndTax;
        }

        emit AddOrder(ps.tIndex, ps.uIndex, ps.pIndex, ps.start, ps.end, ps.size, ps.sprice);
    }

    function subOrder(OrderIn memory ps, uint64 _gi) external override onlyOwner {
        require(fs[ps.uIndex].ao[ps.pIndex].subNonce == ps.nonce, "nonce not equal fs.subnonce"); // nonce error
        require(fs[ps.uIndex].ao[ps.pIndex].nonce > ps.nonce, "nonce not less fs.nonce"); // nonce error; add then sub

        // update size and sprice
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].sprice -= ps.sprice;
        fs[ps.uIndex].ao[ps.pIndex].sInfo[ps.tIndex].size -= ps.size;

        // update settlement
        _settlementSub(ps.pIndex, ps.tIndex, ps.end, ps.size, ps.sprice);

        fs[ps.uIndex].ao[ps.pIndex].subNonce++;

        groups[_gi][ps.tIndex].size -= ps.size;
        groups[_gi][ps.tIndex].sprice -= ps.sprice;

        emit SubOrder(ps.tIndex, ps.uIndex, ps.pIndex, ps.start, ps.end, ps.size, ps.sprice);
    }

    function addRepair(OrderIn memory ps, uint64 _gi) external override onlyOwner {
        require(fs[ps.uIndex].ao[ps.pIndex].nonce == ps.nonce, "nonce not equal fs.nonce");

        uint256 pay = (ps.end-ps.start) * ps.sprice;
        groups[_gi][ps.tIndex].lost -= pay;

        _settlementAdd(ps.pIndex, ps.tIndex, ps.end, ps.size, ps.sprice, pay);

        fs[ps.uIndex].ao[ps.pIndex].nonce += 1;

        emit AddRepair(ps.tIndex, ps.uIndex, ps.pIndex, ps.start, ps.end, ps.size, ps.sprice);
    }

    function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) external override onlyOwner {
        if (isLock) {
            lock[_i][_ti] += money;
        } else {
            balances[_i][_ti] += money;
        }  
    }

    // provider withdraw money, called by owner
    function proWithdraw(PWIn memory ps, uint64 _gi) external override onlyOwner returns(uint256) {
        // pay to provider
        (uint256 thisPay, uint256 _lost) = _settlementCal(ps.pIndex, ps.tIndex, ps.pay, ps.lost);

        groups[_gi][ps.tIndex].lost += _lost;

        // add to balance, then withdraw from it
        balances[ps.pIndex][ps.tIndex] += thisPay;
        
        return thisPay;
    }

    function withdraw(uint64 _i, uint8 _ti, uint256 amount) external override onlyOwner returns(uint256){
        uint256 bal = balances[_i][_ti];
        if(amount>bal){
            amount=bal;
        }

        balances[_i][_ti] -= amount;
        return amount;
    }

    // ================get=============
    // get groupinfo
    function getGInfo(uint64 _gi, uint8 _ti) external view override returns (GroupOut memory){
        GroupOut memory g;
        g = GroupOut(groups[_gi][_ti].size,groups[_gi][_ti].sprice,groups[_gi][_ti].lost);
        return g;
    }

    function getFsInfo(uint64 _ui, uint64 _pi) external view override returns (FsOut memory){
        FsOut memory f;
        f = FsOut(fs[_ui].ao[_pi].nonce,fs[_ui].ao[_pi].subNonce);
        return f;
    }

    // get storeInfo in fs
    function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) external view override returns (StoreOut memory){
        StoreOut memory s;
        s = StoreOut(fs[_ui].ao[_pi].sInfo[_ti].start,fs[_ui].ao[_pi].sInfo[_ti].end,fs[_ui].ao[_pi].sInfo[_ti].size,fs[_ui].ao[_pi].sInfo[_ti].sprice);
        return s;
    }

    // 获得支付计费相关的信息
    function getSettleInfo(uint64 _pi, uint8 _ti) external view override returns (SettleOut memory){
        SettleOut memory s;
        s = SettleOut(proInfo[_pi][_ti].time,proInfo[_pi][_ti].size,proInfo[_pi][_ti].sprice,proInfo[_pi][_ti].maxPay,proInfo[_pi][_ti].hasPaid,proInfo[_pi][_ti].canPay,proInfo[_pi][_ti].lost);
        return s;
    }

    function balanceOf(uint64 _i, uint8 _ti) external view override returns(uint256, uint256){
        uint256 avail = balances[_i][_ti];
        uint256 _lock = lock[_i][_ti];
        Settlement memory se = proInfo[_i][_ti];
        uint256 canPay = se.canPay;
        if (se.time > 0) {
            if (block.timestamp > se.time) {
                canPay += uint256(block.timestamp - se.time) * se.sprice;
            }

            uint256 tmp = se.maxPay - se.lost;
            if(canPay > tmp){
                canPay = tmp;
            }

            _lock += (canPay - se.hasPaid);
        }
        
        return (avail, _lock);
    }
}