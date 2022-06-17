// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./IFileSysSetter.sol";
import "./IOwner.sol";

interface IProxy {
    event Activate(uint64 i);
    event Ban(uint64 i, bool ban);
    event AddT(address t, bool mint);
    event BanT(address t, bool ban);
    event BanG(uint64 gi, bool ban);

    function activate(uint64 _i, bytes[5] memory signs) external;
    function ban(uint64 _i, bool _ban, bytes[5] memory signs) external;
    function addT(address _t, bool _mint, bytes[5] memory signs) external;
    function banT(address _t, bool _ban, bytes[5] memory signs) external;
    function banG(uint64 _gi, bool _ban, bytes[5] memory signs) external;

    function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) external;
    // register self to get index
    function reAcc() external; 
    function reRole(uint8 _rtype, bytes memory _extra) external;
    // add a user/keeper/provider to group
    function addToGroup(uint64 _gi) external;
    
    function pledge(uint64 _i, uint256 _money) external;
    function unpledge(uint64 _i, uint8 _ti, uint256 _money) external;

    function alterPayee(address _p) external;
    function confirmPayee(uint64 _i) external;

    function addOrder( OrderIn memory _oi, bytes memory uSign, bytes memory pSign) external;
    function subOrder(OrderIn memory _oi, bytes memory uSign, bytes memory pSign) external;

    function addRepair(OrderIn memory _oi, uint64[] memory _kis, bytes[] memory ksigns) external;

    function recharge(uint64 _i, uint8 _ti, bool isLock, uint256 _money) external;
    function withdraw(uint64 _i, uint8 _ti, uint256 _money) external;
    function proWithdraw(PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;
}