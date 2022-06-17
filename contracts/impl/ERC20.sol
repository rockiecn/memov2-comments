// SPDX-License-Identifier:UNLICENSED
pragma solidity ^0.8.0;

import "../interfaces/IERC20.sol";
import "../interfaces/IAccess.sol";

/**
 *@author MemoLabs
 *@title erc20 token in memo system
 */
contract ERC20 is IERC20 {
    uint16 public version = 2;

    address public access;

    uint256 public override totalSupply; // 上限6亿；初始发行3亿
    uint256 public constant initialSupply = 3 * 10**26;
    uint256 public constant maxSupply = 6 * 10**26;

    string public override name;
    string public override symbol;

    mapping(address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowances;

    /// @dev created by admin. gas: 1204960
    constructor(
        address _access,
        string memory _name,
        string memory _symbol
    ) {
        name = _name;
        symbol = _symbol;
        access = _access;
        totalSupply = initialSupply;
        balances[msg.sender] += initialSupply;

        emit Transfer(address(0), msg.sender, initialSupply);
    }

    function balanceOf(address acc) external view override returns (uint256) {
        return balances[acc];
    }

    function allowance(address owner, address spender)
        external
        view
        override
        returns (uint256)
    {
        return allowances[owner][spender];
    }

    // gas: 59884
    function transfer(address recipient, uint256 amount)
        external
        override
        returns (bool)
    {
        _transfer(msg.sender, recipient, amount);
        return true;
    }

    // gas: 53518
    function approve(address spender, uint256 amount)
        external
        override
        returns (bool)
    {
        _approve(msg.sender, spender, amount);
        return true;
    }

    // gas: 69266
    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external override returns (bool) {
        uint256 ca = allowances[sender][msg.sender];
        require(ca >= amount, "exce bal"); // transfer amount exceeds balance
        _transfer(sender, recipient, amount);
        _approve(sender, msg.sender, ca - amount);
        return true;
    }

    function _transfer(
        address sender,
        address recipient,
        uint256 amount
    ) internal {
        uint256 sb = balances[sender];
        require(sb >= amount, "exce bal"); // amount exceeds balance

        balances[sender] = sb - amount;
        balances[recipient] += amount;
        emit Transfer(sender, recipient, amount);
    }

    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) internal {
        allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    // use a mint contract if need
    // gas: 56659(called by other contract)
    function mint(address target, uint256 mintedAmount) external override {
        require(IAccess(access).access(msg.sender), "caller no mintAccess"); // can not mint; only contract address can mint

        totalSupply += mintedAmount;
        require(totalSupply <= maxSupply, "mint exce limit"); // exceed the limit

        balances[target] += mintedAmount;
        emit Transfer(address(0), target, mintedAmount);
    }

    // everyone can burn its value; need it or add it in transfer
    // gas: 38435
    function burn(uint256 burnAmount) external override {
        uint256 sb = balances[msg.sender];
        require(sb >= burnAmount, "burnAmount exce bal"); // burn amount exceeds balance
        balances[msg.sender] = sb - burnAmount;
        totalSupply -= burnAmount;
        emit Transfer(msg.sender, address(0), burnAmount);
    }
}
