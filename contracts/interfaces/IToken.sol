// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "./ITokenSetter.sol";
import "./ITokenGetter.sol";

interface IToken is ITokenSetter, ITokenGetter {}