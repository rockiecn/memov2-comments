// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

// re:register; ac: account; t: token; p: provider; k: keeper; u: user; pp: pledgepool
interface IGroupSetter {
    event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr);
    
    // called by admin
    function ban(uint64 _gi, bool _isBan) external;

    function activate(uint64 _gi, uint8 _kc) external;
    function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) external;
}