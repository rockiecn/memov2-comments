// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IToken.sol";
import "./Owner.sol";

/**
 *@author MemoLabs
 *@title Manage token addresses that memo supported.
 */
contract Token is IToken, Owner {
    struct TokenInfo {
        bool mint;     // mint or not
        bool isBan; // 该币种是否被禁止
        uint8 index;   // 代币序列号
    }

    uint16 public version = 2;

    address[] tokens;
    mapping(address => TokenInfo) tInfo;

    // gas: 1437076
    constructor(address _a) Owner(_a) {
    }

    function addT(address _t, bool _mint) external override onlyOwner returns(uint8) {
        require(!_hasToken(_t), "has token"); // token exist 

        uint8 tIndex = uint8(tokens.length);
        tokens.push(_t);
        tInfo[_t].index = tIndex;
        tInfo[_t].mint = _mint;

        emit AddT(_t, tIndex);
        return tIndex;
    }

    function banT(address _t, bool _ban) external override onlyOwner {
        tInfo[_t].isBan = _ban;
        emit BanT(_t);
    }

    function _hasToken(address tAddr) internal view returns (bool) {
        if (tInfo[tAddr].index != 0) {
            return true;
        }

        if (tokens.length > 0 && tokens[0] == tAddr) {
            return true;
        }

        return false;
    }

    function checkT(uint8 tIndex) external view override returns (address, bool) {
        require(!tInfo[tokens[tIndex]].isBan, "token banned");
        return (tokens[tIndex], tInfo[tokens[tIndex]].mint);
    }

    // =========get==========
    function getTA(uint8 tIndex) external view override returns (address) {
        return tokens[tIndex];
    }

    function getTI(address t) external view override returns (uint8, bool, bool) {
        uint8 tIndex = tInfo[t].index;
        if(t==tokens[tIndex]){ // 注册过
            return (tIndex, tInfo[t].isBan, tInfo[t].mint);
        }
        return (tIndex, true, false);
    }

    function getTCnt() external view override returns (uint8) {
        return uint8(tokens.length);
    }
}