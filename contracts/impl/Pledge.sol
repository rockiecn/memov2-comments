//SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IPledge.sol";
import "../interfaces/IToken.sol";
import "../interfaces/IInstance.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Manage pledge in memo system
 */
contract Pledge is IPledge, Owner {

    struct RewardInfo {
        uint256 accu; // 本代币的accumulator, multi by 10^18
        uint256 last; // 上一次变更时的本代币奖励总量。等于余额加奖励
    }

    bytes4 private constant SELECTOR0 = bytes4(keccak256(bytes('balanceOf(address)')));

    uint16 public version = 2;

    address public inst;

    uint8 private cnt;  // token count     

    uint256 public override totalPledge; // 账户质押量

    mapping(uint8 => RewardInfo) public tInfo;  // 每种代币的信息, tokenIndex=>RewardInfo
    mapping(uint64 => mapping(uint8 => RewardInfo)) public rewards; // 所有质押的人的信息，(roleIndex => tokenIndex => RewardInfo)

    // gas: 2080678
    constructor(address _a, address _inst) Owner(_a){
        inst = _inst;
    }

    function _updateReward(uint256 amount, uint8 _ti, uint64 _index) internal {
        uint256 bal = _getBalance(_ti); // 获取质押池在该代币上的余额

        // 代币_ti有增发，需更新tInfo[_ti].accu
        // 需加上等于判断，一旦有人unpledge，就有可能出现tInfo[ti].last和bal相等的情况
        if (bal >= tInfo[_ti].last && totalPledge>0 ){
            tInfo[_ti].accu += (bal - tInfo[_ti].last) * 1e18 / totalPledge; // 此刻每质押主代币应分得的利润
            tInfo[_ti].last = bal; // update to latest

            // 将账户余额加上应得的分润值
            // res表示：距离上次更新分润值期间，每质押主代币应得的代币_tis分润值的累积总和
            uint256 res = tInfo[_ti].accu - rewards[_index][_ti].accu;
            res = res * amount / 1e18; // amount应大于1e18
            rewards[_index][_ti].last += res; // 添加分润值
            rewards[_index][_ti].accu = tInfo[_ti].accu; // 更新accu
        }
    }
    
    function pledge(uint64 _index, uint256 money) external onlyOwner override {
        RewardInfo memory reward = rewards[_index][0];
        uint256 amount = reward.last;

        // 更新结算奖励
        for(uint8 i=0; i<cnt; i++){
            reward = rewards[_index][i];
            if(reward.accu==0){
                // init 
                rewards[_index][i].accu = tInfo[i].accu;
            }
            _updateReward(amount, i, _index);
        }
        
        // update
        tInfo[0].last += money; // wrong if transfer first
        rewards[_index][0].last += money;
        totalPledge += money;
    }

    function withdraw(uint64 _index, uint8 _ti, uint256 money, uint256 _lock) external onlyOwner override returns (uint256) {
        RewardInfo memory reward = rewards[_index][0];
        uint256 amount = reward.last;

        // no pledge
        if(amount==0 || totalPledge == 0){
            return 0;
        }

        // 更新结算奖励
        if (_ti == 0) {
            // update all tokens due to pledge change
            for(uint8 i=0; i<cnt; i++) {
                _updateReward(amount, i, _index);
            }
        } else {
            // only update i
            _updateReward(amount, _ti, _index);
        }
    
        // 确定value
        uint256 rw = rewards[_index][_ti].last - _lock;
        if(money<rw){
            rw = money;
        }

        // update token
        tInfo[_ti].last -= rw;

        // update account value
        rewards[_index][_ti].last -= rw;

        // update totalPledge
        if(_ti==0){
            totalPledge -= rw;
        }

        return rw;
    }

    function addT(uint8 _ti) external onlyOwner override {
        require(_ti == cnt, "ille ti"); // illegal token
        tInfo[_ti].accu = 0;
        tInfo[_ti].last = _getBalance(_ti);
        cnt+=1;
    }

    // ========== get ===========
    // get balance of Pool-address with tIndex 
    function _getBalance(uint8 _ti) internal view returns (uint256) {
        IInstance insti = IInstance(inst);

        address tAddr = ITokenGetter(insti.instances(7)).getTA(_ti);
        (bool success, bytes memory data) = tAddr.staticcall(abi.encodeWithSelector(SELECTOR0, insti.instances(5)));
        require(success && data.length >= 32, "staticcall err"); // staticcall error
        return abi.decode(data, (uint256));
    }

    // pledge + mint, balance of pool
    function getPledge(uint8 _ti) external view override returns (uint256) {
        return _getBalance(_ti);
    }

    // get balance of index, only calculate the reward, don't update the reward
    function balanceOf(uint64 _index, uint8 _ti) external view override returns (uint256) {
        RewardInfo memory reward0 = rewards[_index][0];
        if (reward0.last==0) {
            return 0;
        }
        uint256 amount = reward0.last;

        RewardInfo memory rewardt = tInfo[_ti];
        uint256 val = rewardt.accu;

        uint256 bal = _getBalance(_ti);
        if(bal > rewardt.last && totalPledge > 0) {
            val += (bal - rewardt.last) * 1e18 / totalPledge;
        }

        RewardInfo memory rewardi = rewards[_index][_ti];
        val = rewardi.last + (val - rewardi.accu)*amount/1e18;

        return val;
    }
}