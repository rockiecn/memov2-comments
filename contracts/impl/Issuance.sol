// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IIssuance.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title mint reward in memo system
 */
contract Issuance is IIssuance, Owner {
    uint16 public version = 2;

    // 1. 每个阶段的奖励上限是由本阶段数据量, ratio决定；
    // 2. 达到阶段奖励上限后无奖励，直到数据量变大解锁下一个阶段奖励；
    // 3. 在有很多高价订单时会出现早早用完阶段奖励的情况；
 
    // 1 nanoMemo / second*byte, about 3 memo/(month*TB), about 1500 EB*day 
    uint256 public targetReward = 30*10**25; // 300 million reward limit; 
    uint256 public rewardResid = targetReward; // residual reward 

    // init value of stage
    uint256 public stageRatio = 10000; // reward at most 1 nanoMemo per second*byte; 150% every stage
    uint256 public stageSize = 1024*1024*1024*1024; // 1TB, multi 32 every stage
    uint256 public stageReward = stageSize * stageRatio * 864; // set every stage 
    uint256 public stageRewardResid = stageReward; // residual reward at current stage

    uint64 public last;         // last mint time
    uint64 public size;
    uint256 public sprice;      // totalSizePrice
    uint256 public spaceTime;   // totalSpaceTime

    mapping(uint64 => uint256) public expire; // 保存sprice at each end

    // gas: 1555229
    constructor(address _a) Owner(_a) {
        last = uint64(block.timestamp); // now time
    }

    function issu(uint64 _start, uint64 _end, uint64 _size, uint256 _sPrice) external override onlyOwner returns (uint256) {
        // add sub sprice and size
        expire[_end] += _sPrice;

        uint64 nowTime = uint64(block.timestamp);
        if(nowTime > last + 86400){
            nowTime = last + 86400;
        }
        
        // unlock reward accord to sprice*duration
        uint256 reward = sprice * (nowTime - last);
        if(last/86400 < nowTime/86400){
            // 至少一天后调用
            uint64 midTime = nowTime/86400 * 86400; // 整天数对应时间
            uint256 sp = expire[midTime];
            sprice -= sp;
            reward -= sp * (nowTime - midTime);
        }
        last = nowTime;
  
        reward = reward * stageRatio / 10000;
        if (reward > stageRewardResid) {
            reward = stageRewardResid;    
        }

        if (reward > rewardResid) {
            reward = rewardResid;
        }

        stageRewardResid -= reward;
        rewardResid -= reward;

        // for reward stage
        spaceTime += _size * uint256(_end - _start);
        sprice += _sPrice;
        size += _size;

        // check stage, normalize to 100 day
        if (spaceTime/8640000 >= stageSize) {
            // go to next stage
            // after 1EB, double every stage
            if (stageSize >= 10**18) {
                stageSize *= 2;    
            } else {
                stageSize *= 32;
            }

            uint256 dur = spaceTime/(size*10000); // average duration

            stageRatio = stageRatio*15/10; // 150% of previous
            stageReward = stageSize * stageRatio * dur;
            if (stageReward >= rewardResid) {
                // 50% of rest as stageReward
                stageRatio = rewardResid / (stageSize*dur*2);
                stageReward = stageSize * stageRatio * dur;
            }
            stageRewardResid = stageReward;
        }

        return reward;
    }
}