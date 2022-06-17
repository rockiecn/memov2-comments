# memo-contracts-v2

`@author:Memolabs`

此仓库包含了 memo 项目中的智能合约文件。

## 环境

**语言** solidity

**编辑器** vscode

**语法及功能测试** solc、remix

**集成调试工具** Hardhat

## 文件

合约文件在 contracts 目录下。

Access.sol 以及 ERC20.sol 管理MEMO系统代币信息；
Auth.sol 以及 Owner.sol 管理MEMO系统中的权限信息；
Proxy.sol 是MEMO系统中的代理合约，也是与用户直接交互的合约，用户通过该合约发起交易；
Control1.sol 管理MEMO系统中的Role、Pledge、Token、PoolP（质押池，Pool合约）合约；
Control2.sol 管理MEMO系统中的FileSys、Issue、erc20、PoolF、Kmanage合约；
Getter.sol 是与用户直接交互的合约，用户通过该合约查询MEMO系统中其他合约的信息；
FileSys.sol 管理MEMO系统中文件存储订单、订单结算信息；
Group.sol 管理MEMO系统中组的信息；
Issuance.sol 管理MEMO系统中代币发行的信息；
Kmanage.sol 管理MEMO系统中Keepers的收益信息；
Pledge.sol 管理MEMO系统中账户质押的信息；
Pool.sol 用来存放MEMO系统账户质押资金、充值资金（会部署两个Pool合约，一个PoolP，一个PoolF）；
Role.sol 管理MEMO系统中账户注册信息和角色信息；
Token.sol 管理MEMO系统支持的代币信息；

## 注册角色流程

User: reAcc、reRole、addToGroup
Provider: reAcc、reRole、Pledge、addToGroup
Keeper: reAcc、reRole、Pledge、addToGroup、activate(by admin)

## 注意事项

由于交易gas限制，需要对group里面的keeper数量以及Token合约中的token数量进行限制。通过测试Control2.sol中的proWithdraw()函数最多能承受的keeper量，确定keeper上限。（根据for循环）

优化 gas 消耗，需考虑：

- 将 msg.sender、msg.value、tx.origin 赋值给局部变量，避免多次访问 msg、tx 信息
- 限制数组的使用
- 限制修改及访问 storage 存储

## 合约部署

替换instances中的control2合约地址后，还需将旧control2合约的access权限设为false;