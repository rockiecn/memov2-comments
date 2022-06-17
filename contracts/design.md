# design

## todo

1. getter: 精简get接口，特别是数组内容，用线下for循环获取
2. issu: 增长逻辑

## root contract

+ proxy

## contract type

+ 2 auth (set authrity) not upgrade
+ 3 access (erc control) can upgrade
+ 4 erc (erc data) non-upgrade
+ 5 poolp (pledge data) non-upgrade  
+ 6 role (role and group data) non-upgrade 
+ 7 token (token data) non-upgrade
+ 8 pledge (pledge control) can upgrade
+ 9 issue (issue control) can upgrade
+ 10 fs (fs data) non-upgrade
+ 11 group (group data) non-upgrade
+ 12 poolf (fs data) non-upgrade

+ 100 proxy
+ 101 control1 control role, group and pledge (interact control) can upgrade
+ 102 control2 control fs (interact control) can upgrade

+ 150 getter

kmanage (keeper profit data) has no contract type, deployed by Group.sol, get it by group.getKManage()


## control contracts

auth
instance
controlx

## base contracts

erc
token
pool
pledge
role
group
issue
fs
kmanage

## layer

root: instance, find contract address from inst 

proxy -> control1/2 -> base contracts 

getter -> base contracts 

## owner

control1    (proxy)
control2    (proxy)
token   (control1)
filesys (control2)
group   (control1)
kmanage (control2)
issuance    (control2)
poolp   (control1)
poolf   (control2)
pledge  (control1)
role    (control1)

## deploy

### deploy erc

+ deploy access => access address
+ deploy erc(access) => erc address

### deploy other

+ deploy auth => auth address
+ deploy inst(auth) => inst address 
+ deploy proxy(auth, inst)  => proxy address as root
+ deploy control1(auth, inst) => control1 address
+ deploy control2(auth, inst) => control2 address
+ deploy token(auth) => token address
+ deploy poolp(auth) => poolp address, deployed when createGroup by Role.sol
+ deploy kmanage(auth) => kmanage address, deployed when createGroup by Role.sol
+ deploy pledge(auth) => pledge address
+ deploy role(auth, foundation) => role address
+ deploy group(auth, inst) => group address
+ deploy poolf(auth) => poolf address
+ deploy fs(auth) => fs address
+ deploy issuance(auth) => issu address

### deploy getter

+ deploy getter(owner) => getter address

### set address to instance

### set owner

+ set control2 as token, pool, kmanage, pledge, role, group, fs, issuance's 202 owner
+ set token, pool, kmanage, pledge, role, group, fs, issuance to control1/2 address

### add token

+ addT(erc address)

### create group

+ createGroup(level, kpledge, ppledge, manageRatio) => create pool(control, auth), create kmanage(control, auth) => g.pool, g.kmanage
 
### register

+ register account => role index
+ register role => rtype
+ keeper/provider pledge
+ add to group(group should be active when user add) => user/provider is active
+ admin activate keeper => keeper is active

## proxy/control

+ function activate(uint64 _i, bytes[] memory signs) external

admin激活账户，用于keeper加入组后激活

+ function ban(uint64 _i, bool _ban, bytes[] memory signs) external;

admin禁止某账户

+ function addT(address _t, bool _ban, bytes[] memory signs) external;

admin添加代币合约地址

+ function banG(uint64 _gi, bool _ban, bytes[] memory signs) external;

admin禁止某组

+ function createGroup(uint16 _level, uint256 _kr, uint256 _pr, uint8 _mr) external;

创建一个新的组，组中会创建对应的kmanage，pool合约；kr: keeper pledge requirement, typically 1 memo; pr: provider pledge requirement, typically 1 memo; mr: manage rate, typically 4;

+ function reAcc(address _a) external; 

地址注册获得账户号

+ function reRole(address _a, uint8 _rtype, bytes memory _extra) external;

地址注册角色

+ function addToGroup(address _a, uint64 _gi) external;

加入组

+ function pledge(address _a, uint64 _i, uint256 _money) external;

给账户_i质押， 钱从调用者_a中转出

+ function unpledge(address _a, uint64 _i, uint8 _ti, uint256 _money) external;

从质押池取钱

+ function addOrder(address _a, OrderIn memory _oi) external;

验证账户，token合法性
验证user钱是否足够
验证订单合法


+ function subOrder(address _a, OrderIn memory _oi) external;

验证账户，token合法性
验证订单合法
分润给provider、kmanage

+ function recharge(address _a, uint64 _i, uint8 _ti, uint256 _money, bool isLock) external;

给账户_i充值， 钱从调用者_a中转出； isLock表示钱只能用于fs支付，不能提取出来

+ function withdraw(address _a, uint64 _i, uint8 _ti, uint256 _money) external

从fs池中取钱

+ function proWithdraw(address _a, PWIn memory _ps, uint64[] memory _kis, bytes[] memory ksigns) external;

pro取钱

+ function get(uint8 _type) external view returns(address); 

获取合约地址

## golang

+ go-mefs only need proxy.sol(include Owner.sol, which contains address of other contracts) and getter.sol; 

## upgrade

+ data sol(erc、pool、role、token、fs、kmanage): better no upgrade
+ control sol(control、access): can upgrade 
+ base control sol(pledge、issu): can upgrage
+ auth sol: can upgrade, be caution;

## addorder

order的duration介于100至1000天之间，且order.end需要是整天数。
control2合约地址需要有access权限去mint。