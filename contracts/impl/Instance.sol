// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IInstance.sol";
import "../interfaces/IAuth.sol";

/**
 *@author MemoLabs
 *@title Manage owner and auth in memo system
 */
contract Instance is IInstance {
    // can support mask if using uint32 instead of uint8
    mapping(uint8=>address) public override instances;

    constructor(address _a) {
        instances[2] = _a;
    }

    function alter(address _a, uint8 _type, bytes[5] memory signs) external {
        uint size;
        assembly {
            size := extcodesize(_a)
        }
        require(size != 0,"need contract addr"); // need contract addr
        
        address au = instances[2];
        bytes32 h = keccak256(abi.encodePacked(address(this), "alter", _a, _type));
        IAuth(au).perm(h, signs);

        emit Alter(_type, instances[_type], _a);
        instances[_type] = _a;
    }
}