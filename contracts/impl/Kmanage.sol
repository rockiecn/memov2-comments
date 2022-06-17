// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IKmanage.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title keeper management in memo system, per group
 */
contract Kmanage is IKmanage, Owner {

    struct ProfitInfo {
        uint64 last; // 上次分利润时间
        uint256 accu; // 记录分润值
    }

    uint16 public version = 2;
    uint8 public override manageRate = 4; // default 4%, 3% linear and 1% at end;

    uint64 period = 7*86400;     // keeper根据比例获取收益的时间间隔, one week           

    uint64 tc; // 记录所有keeper触发order相关函数的总次数
    uint64[] keepers;
    mapping(uint64 => uint64) counter; // 记录keeper触发Order相关函数的次数，用于分润

    mapping(uint8 => ProfitInfo) pfinfo; // 记录分润的金额和时间

    mapping(uint64 => mapping(uint8 => uint256)) balances; // 账户可用的余额
    
    // owner: Control-contract-address
    constructor(address _a, uint8 mr) Owner(_a) {
        manageRate = mr;
    }

    // after add keeper to group, called by Control.sol
    function addKeeper(uint64 _ki) external override onlyOwner {
        require(counter[_ki] == 0, "k exist"); // keeper exists
        keepers.push(_ki);
        counter[_ki] = 1;
        tc++;
    }

    // after sub order, called by Control.sol
    function addCnt(uint64 _ki, uint64 cnt) external override onlyOwner {
        require(counter[_ki] > 0, "sOrder:ille caller");
        counter[_ki] += cnt;
        tc += cnt;      
        emit AddCnt(_ki, cnt);
    }

    // after pro withdraw and subOrder, called by Control.sol
    function addProfit(uint8 _ti, uint256 _money) external override onlyOwner {    
        pfinfo[_ti].accu += _money;
        emit AddProfit(_ti, _money); 
    }

    // after withdraw, called by Control.sol
    function withdraw(uint64 _ki, uint8 _ti, uint256 amount) external override onlyOwner returns(uint256){
        if (counter[_ki] == 0) {
            return 0;
        }
        uint64 ntime = uint64(block.timestamp);
        if(ntime-pfinfo[_ti].last > period){
            pfinfo[_ti].last = ntime; 
            ntime = 0;
            uint256 per = pfinfo[_ti].accu / uint256(tc);
            // reduce counter every period; like divided by 2;
            for(uint j =0; j<keepers.length; j++){
                uint256 pro = per * counter[keepers[j]];
                counter[keepers[j]] = counter[keepers[j]]/2 + 1; // assure > 0
                ntime += counter[keepers[j]];
                balances[keepers[j]][_ti] += pro;
                pfinfo[_ti].accu -= pro; // accu is not zero if not divisible
            }
            tc = ntime;
        }
    
        uint256 bal = balances[_ki][_ti];
        if(amount>bal){
            amount=bal;
        }

        balances[_ki][_ti] -= amount;
        return amount;
    }   

    // ================get=============

    function getPf(uint8 _ti) external view override returns (uint64, uint256) {
        return (pfinfo[_ti].last, pfinfo[_ti].accu);
    }

    function getKCnt() external view override returns (uint8) {
        return uint8(keepers.length);
    }

    function getK(uint64 _i) external view override returns (uint64) {
        return keepers[_i];
    }

    function balanceOf(uint64 _ki, uint8 _ti) external view override returns(uint256, uint256){
        uint256 avail = balances[_ki][_ti];
        uint256 tmp = 0;

        uint256 pt = pfinfo[_ti].accu / tc * counter[_ki]; // 和上面计算保持一致
        if((block.timestamp-pfinfo[_ti].last)>period){
            avail += pt;
        }else{
            tmp += pt;
        }
    
        return (avail, tmp);
    }
}