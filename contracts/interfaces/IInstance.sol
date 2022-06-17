// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

interface IInstance {

    event Alter(uint8 _type, address from, address to);
    
    function instances(uint8 _type) external view returns(address);
}