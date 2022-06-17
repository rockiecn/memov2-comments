// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IOwner.sol";
import "../interfaces/IAuth.sol";

/**
 *@author MemoLabs
 *@title Manage owner and auth in memo system
 */
contract Owner is IOwner {
    // can support mask if using uint32 instead of uint8
    mapping(address=>bool) public owners;
    address public auth; 

    constructor(address _a) {
        auth = _a;
    }

    //函数修饰符，判断是不是owner调用
    modifier onlyOwner(){
        require(owners[msg.sender], "not owner");
        _;
    }

    function add(address _a, bool isSet, bytes[5] memory signs) external {
        uint size;
        assembly {
            size := extcodesize(_a)
        }
        require(size != 0,"need contract addr"); // need contract addr

        bytes32 h = keccak256(abi.encodePacked(address(this), "add", _a, isSet));
        IAuth(auth).perm(h, signs);

        emit AddOwner(_a, isSet);
        owners[_a] = isSet;
    }
}