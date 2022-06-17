// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

struct StoreOut {
    uint64 start; // last start 
    uint64 end;   // 什么时刻的状态，last end time
    uint64 size;   // 在该存储节点上的存储总量，byte
    uint256 sprice; // 按周期计费; per cycle
}

struct FsOut {
    uint64 nonce;    // 防止order重复提交
    uint64 subNonce; // 用于订单到期
}

struct GroupOut {
    uint64 size;   
    uint256 sprice;
    uint256 lost;
}

// Settlement indicates billing information
struct SettleOut {
    uint64 time; // store状态改变或支付的时间, 与 epoch 对齐
    uint64 size; // 在该存储节点上的存储总量
    uint256 price; // 累积的sprice(即sizePrice)

    uint256 maxPay;  // 对此provider所有user聚合总额度； expected 加和
    uint256 hasPaid; // 已经支付
    uint256 canPay;  // 最近一次store/pay时刻，可以支付的金额
    uint256 lost;    // lost due to unable response to challenge
}

interface IFileSysGetter {
    function balanceOf(uint64 _i, uint8 _ti) external view returns(uint256, uint256);
    function getGInfo(uint64 _gi, uint8 _ti) external view returns (GroupOut memory);
    function getFsInfo(uint64 _ui, uint64 _pi) external view returns (FsOut memory);
    function getStoreInfo(uint64 _ui, uint64 _pi, uint8 _ti) external view returns (StoreOut memory);
    function getSettleInfo(uint64 _pi, uint8 _ti) external view returns (SettleOut memory);
}