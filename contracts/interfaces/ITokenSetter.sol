// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface ITokenSetter {
    event AddT(address t, uint8 tIndex);
    event BanT(address t);

    function addT(address _t, bool _mint) external returns(uint8);
    function banT(address _t, bool _isBan) external; 
}