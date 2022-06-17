// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;


interface IControl1 {
    event Pledge(address indexed payer, uint64 indexed i, uint256 money);
    event Unpledge(address indexed payee, uint64 indexed i, uint8 indexed ti, uint256 money, uint256 lock);

    // called by admin 
    function activate(uint64 _i, bytes[5] memory signs) external;          // 6, 11;11
    function ban(uint64 _i, bool _ban, bytes[5] memory signs) external;    //6
    function addT(address _t, bool _mint, bytes[5] memory signs) external;  //7,8
    function banT(address _t, bool _ban, bytes[5] memory signs) external;
    function banG(uint64 _gi, bool _ban, bytes[5] memory signs) external;  //11 

    function createGroup(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) external; //11
    function reAcc(address _a) external;  // 6
    function reRole(address _a, uint8 _rtype, bytes memory _extra) external; // 6;6
    function addToGroup(address _a, uint64 _gi) external; // 6; 6,8,11
    
    function pledge(address _a, uint64 _i, uint256 _money) external; // 8; 7,5
    function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) external; // 5,6,7,8,11

    function alterPayee(address _a, address _p) external;
    function confirmPayee(address _a, uint64 _i) external;
}