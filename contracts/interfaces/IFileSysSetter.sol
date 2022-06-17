// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;


// struct for addOrder params
struct OrderIn {
    uint64 uIndex;
    uint64 pIndex;
    uint64 start;
    uint64 end;
    uint64 size;
    uint64 nonce;
    uint8 tIndex;
    uint256 sprice;
}

// struct for proWithdraw params
struct PWIn {
    uint64 pIndex;
    uint8 tIndex;
    uint256 pay;
    uint256 lost;
}

interface IFileSysSetter {
    event AddOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice);
    event SubOrder(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice);
    event AddRepair(uint8 indexed ti, uint64 indexed ui, uint64 indexed pi, uint64 start, uint64 end, uint64 size, uint256 sprice);

    function addOrder(OrderIn memory ps, uint256 _mr, uint64 _gi) external;
    function subOrder(OrderIn memory ps, uint64 _gi) external;

    function addRepair(OrderIn memory ps, uint64 _gi) external;

    function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 money) external;
    function withdraw(uint64 _i, uint8 _ti, uint256 money) external returns (uint256);
    function proWithdraw(PWIn memory ps, uint64 _gi) external returns(uint256);
}