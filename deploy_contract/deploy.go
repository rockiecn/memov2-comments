package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/access"
	"github.com/memoio/contractsv2/go_contracts/auth"
	"github.com/memoio/contractsv2/go_contracts/control1"
	"github.com/memoio/contractsv2/go_contracts/control2"
	"github.com/memoio/contractsv2/go_contracts/erc"
	"github.com/memoio/contractsv2/go_contracts/fs"
	"github.com/memoio/contractsv2/go_contracts/getter"
	"github.com/memoio/contractsv2/go_contracts/group"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
	"github.com/memoio/contractsv2/go_contracts/issuance"
	"github.com/memoio/contractsv2/go_contracts/kmanage"
	"github.com/memoio/contractsv2/go_contracts/pledge"
	"github.com/memoio/contractsv2/go_contracts/pool"
	"github.com/memoio/contractsv2/go_contracts/proxy"
	"github.com/memoio/contractsv2/go_contracts/role"
	"github.com/memoio/contractsv2/go_contracts/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	eth        string
	hexSk      string
	sks        [5]string
	as         [5]common.Address
	foundation common.Address
	kp         *big.Int
	pp         *big.Int

	waitTime1 = 30 * time.Second
	waitTime2 = 40 * time.Second
	bigone    = big.NewInt(1)

	groupLevel = uint8(7)
	mr         = uint8(4)
	dc         = uint8(3)
	pc         = uint8(2)

	gi     = uint64(1)
	ti     = uint8(0)
	ismint = uint8(1) // true
	mint   = true
	isbanG = uint8(0) // false
	banG   = false
	set    = true
)

// deploy all contracts by admin with specified chain
func main() {
	inputeth := flag.String("eth", "https://devchain.metamemo.one:8501", "eth api Address;") //dev net
	sk := flag.String("sk", "", "signature for sending transaction")
	sk1 := flag.String("sk1", "", "addr1's sk")
	sk2 := flag.String("sk2", "", "addr2's sk")
	sk3 := flag.String("sk3", "", "addr3's sk")
	sk4 := flag.String("sk4", "", "addr4's sk")
	sk5 := flag.String("sk5", "", "addr5's sk")
	addrs := flag.String("5addrs", "0xB651b3ce6C7c712E41b03993a11DA656eF3C95A8,0x54a18833beaa8A22672ffF4e5b6714eCebC9368b,0x6f7bc6032278753c8A13bd430b13f1965407633f,0x09d186ce74f9092f5e6F20390d730943170bDbBA,0x8AadBf87Dc8Ae4A1161DEF8185D2b98C1904a3c3", "5 addrs to authorize")
	foun := flag.String("f", com.Foundation.Hex(), "foundation")
	kpStr := flag.String("kp", com.PledgeAmount, "keeper pledge amount")
	ppStr := flag.String("pp", com.PledgeAmount, "provider pledge amount")

	flag.Parse()

	eth = *inputeth
	hexSk = *sk
	sks = [5]string{*sk1, *sk2, *sk3, *sk4, *sk5}
	foundation = common.HexToAddress(*foun)

	as = [5]common.Address{}
	addrsSlice := strings.Split(*addrs, ",")
	fmt.Println("addrs num:", len(addrsSlice), " addrs:", addrsSlice)
	if len(addrsSlice) != 5 {
		fmt.Println("please input 5 addrs to authorize!")
		return
	}
	for i := 0; i < 5; i++ {
		as[i] = common.HexToAddress(addrsSlice[i])
	}

	if hexSk == "" || sks[0] == "" || sks[1] == "" || sks[2] == "" || sks[3] == "" || sks[4] == "" {
		fmt.Println("Please input your sk!")
		return
	}

	kp, err := com.ParseMemo(*kpStr)
	if err != nil {
		fmt.Println("Please input keeper pledge amount as required!")
		fmt.Println(err)
		return
	}
	pp, err := com.ParseMemo(*ppStr)
	if err != nil {
		fmt.Println("Please input provider pledge amount as required!")
		fmt.Println(err)
		return
	}

	// get some value
	approveAmount := new(big.Int)
	approveAmount.Add(kp, pp)

	fmt.Println()

	// make auth to send transaction
	txopts := &com.TxOpts{Nonce: nil, GasPrice: nil, GasLimit: com.DefaultGasLimit, Money: nil}
	ctx := context.Background()
	txAuth, err := com.MakeAuth(ctx, eth, hexSk, txopts)
	if err != nil {
		log.Fatal(err)
	}

	// get client
	client, err := ethclient.DialContext(ctx, eth)
	if err != nil {
		log.Fatal(err)
	}

	// deploy Access
	accessAddr, tx, accessIns, err := access.DeployAccess(txAuth, client, as)
	if err != nil {
		log.Fatal("deploy access err:", err)
	}
	fmt.Println("accessAddr: ", accessAddr.Hex())
	_ = tx
	_ = accessIns

	// deploy ERC20
	erc20Addr, tx, erc20Ins, err := erc.DeployERC20(txAuth, client, accessAddr, com.Erc20Name, com.Erc20Symbol)
	if err != nil {
		log.Fatal("deploy erc20 err:", err)
	}
	fmt.Println("erc20Addr: ", erc20Addr.Hex())
	_ = erc20Ins

	// deploy Auth
	authAddr, tx, authIns, err := auth.DeployAuth(txAuth, client, as)
	if err != nil {
		log.Fatal("deploy auth err:", err)
	}
	fmt.Println("authAddr: ", authAddr.Hex())
	_ = authIns

	// deploy Instance
	instanceAddr, tx, instanceIns, err := inst.DeployInstance(txAuth, client, authAddr)
	if err != nil {
		log.Fatal("deploy instance err:", err)
	}
	fmt.Println("instanceAddr: ", instanceAddr.Hex())
	_ = instanceIns

	// deploy Proxy
	proxyAddr, tx, proxyIns, err := proxy.DeployProxy(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy proxy err:", err)
	}
	fmt.Println("proxyAddr: ", proxyAddr.Hex())
	_ = proxyIns

	// deploy Control1
	control1Addr, tx, control1Ins, err := control1.DeployControl1(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy control1 err:", err)
	}
	fmt.Println("control1Addr: ", control1Addr.Hex())
	_ = control1Ins

	// deploy Control2
	control2Addr, tx, control2Ins, err := control2.DeployControl2(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy control2 err:", err)
	}
	fmt.Println("control2Addr: ", control2Addr.Hex())
	_ = control2Ins

	time.Sleep(waitTime2)

	// deploy Token
	tokenAddr, tx, tokenIns, err := token.DeployToken(txAuth, client, authAddr)
	if err != nil {
		log.Fatal("deploy token err:", err)
	}
	fmt.Println("tokenAddr: ", tokenAddr.Hex())
	_ = tokenIns

	// deploy Filesys
	filesysAddr, tx, filesysIns, err := fs.DeployFileSys(txAuth, client, authAddr)
	if err != nil {
		log.Fatal("deploy filesys err:", err)
	}
	fmt.Println("filesysAddr: ", filesysAddr.Hex())
	_ = filesysIns

	// deploy Group
	groupAddr, tx, groupIns, err := group.DeployGroup(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy group err:", err)
	}
	fmt.Println("groupAddr: ", groupAddr.Hex())
	_ = groupIns

	// deploy Issuance
	issuanceAddr, tx, issuanceIns, err := issuance.DeployIssuance(txAuth, client, authAddr)
	if err != nil {
		log.Fatal("deploy issuance err:", err)
	}
	fmt.Println("issuanceAddr: ", issuanceAddr.Hex())
	_ = issuanceIns

	// deploy PoolP
	poolpAddr, tx, poolpIns, err := pool.DeployPool(txAuth, client, authAddr)
	if err != nil {
		log.Fatal("deploy poolp err:", err)
	}
	fmt.Println("poolpAddr: ", poolpAddr.Hex())
	_ = poolpIns

	// deploy Pledge
	pledgeAddr, tx, pledgeIns, err := pledge.DeployPledge(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy pledge err:", err)
	}
	fmt.Println("pledgeAddr: ", pledgeAddr.Hex())
	_ = pledgeIns

	// deploy PoolF
	poolfAddr, tx, poolfIns, err := pool.DeployPool(txAuth, client, authAddr)
	if err != nil {
		log.Fatal("deploy poolf err:", err)
	}
	fmt.Println("poolfAddr: ", poolfAddr.Hex())
	_ = poolfIns

	// deploy Role
	roleAddr, tx, roleIns, err := role.DeployRole(txAuth, client, authAddr, foundation)
	if err != nil {
		log.Fatal("deploy role err:", err)
	}
	fmt.Println("roleAddr: ", roleAddr.Hex())
	_ = roleIns

	// deploy getter
	getterAddr, tx, _, err := getter.DeployGetter(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy getter err:", err)
	}
	fmt.Println("getterAddr: ", getterAddr.Hex())

	time.Sleep(waitTime2)

	fmt.Println()

	// add Control1's owner: proxy
	nonce, err := authIns.Nonce(&bind.CallOpts{From: com.AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("add Control1's owner=>proxy nonce:", nonce)
	hash := com.GetAddOwnerHash(control1Addr, proxyAddr, authAddr, set, nonce)
	signs := getSigns(hash, sks)
	tx, err = control1Ins.Add(txAuth, proxyAddr, set, signs)
	if err != nil {
		log.Fatal("add control1's owner err:", err)
	}

	// add Control2's owner: proxy
	fmt.Println("add Control2's owner=>proxy nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(control2Addr, proxyAddr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = control2Ins.Add(txAuth, proxyAddr, set, signs) // 基于对已有状态检查（即已有状态，nonce为0），所以需要赋值gasLimit
	if err != nil {
		log.Fatal("add control2's owner err:", err)
	}

	// add token's owner=> control1
	fmt.Println("add token's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(tokenAddr, control1Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = tokenIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add token's owner err:", err)
	}

	// add fs's owner=> control2
	fmt.Println("add fs's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(filesysAddr, control2Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = filesysIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add fs's owner err:", err)
	}

	// add group's owner=> control1
	fmt.Println("add group's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(groupAddr, control1Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = groupIns.Add0(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add group's owner err:", err)
	}

	// add issuance's owner=> control2
	fmt.Println("add issuance's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(issuanceAddr, control2Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = issuanceIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add issuance's owner err:", err)
	}

	// add poolp's owner=> control1
	fmt.Println("add poolp's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(poolpAddr, control1Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = poolpIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add poolp's owner err:", err)
	}

	// add pledge's owner=> control1
	fmt.Println("add pledge's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(pledgeAddr, control1Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = pledgeIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add pledge's owner err:", err)
	}

	// add poolf's owner=> control2
	fmt.Println("add poolf's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(poolfAddr, control2Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = poolfIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add poolf's owner err:", err)
	}

	// add role's owner=> control1
	fmt.Println("add role's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(roleAddr, control1Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = roleIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add role's owner err:", err)
	}

	// wait for the deployment contract transaction to complete
	time.Sleep(waitTime1)
	time.Sleep(waitTime2)

	fmt.Println()

	// set instances
	// set instances[100]=> proxy
	nonce, err = authIns.Nonce(&bind.CallOpts{From: com.AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("set instances[100] nonce:", nonce)
	hash = com.GetAlterHash(instanceAddr, authAddr, proxyAddr, com.TypeProxy, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, proxyAddr, com.TypeProxy, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[101]=> control1
	fmt.Println("set instances[101] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, control1Addr, com.TypeControl1, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, control1Addr, com.TypeControl1, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[102]=> control2
	fmt.Println("set instances[102] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, control2Addr, com.TypeControl2, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, control2Addr, com.TypeControl2, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[150]=> getter
	fmt.Println("set instances[150] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, getterAddr, com.TypeGetter, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, getterAddr, com.TypeGetter, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[3]=> access
	fmt.Println("set instances[3] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, accessAddr, com.TypeAccess, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, accessAddr, com.TypeAccess, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[4]=> erc20
	fmt.Println("set instances[4] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, erc20Addr, com.TypeERC20, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, erc20Addr, com.TypeERC20, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[5]=>poolp
	fmt.Println("set instances[5] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, poolpAddr, com.TypePoolP, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, poolpAddr, com.TypePoolP, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[6]=>role
	fmt.Println("set instances[6] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, roleAddr, com.TypeRole, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, roleAddr, com.TypeRole, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[7]=>token
	fmt.Println("set instances[7] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, tokenAddr, com.TypeToken, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, tokenAddr, com.TypeToken, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[8]=>pledge
	fmt.Println("set instances[8] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, pledgeAddr, com.TypePledge, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, pledgeAddr, com.TypePledge, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[9]=>issue
	fmt.Println("set instances[9] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, issuanceAddr, com.TypeIssu, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, issuanceAddr, com.TypeIssu, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[10]=>fs
	fmt.Println("set instances[10] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, filesysAddr, com.Typefs, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, filesysAddr, com.Typefs, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[11]=>group
	fmt.Println("set instances[11] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, groupAddr, com.TypeGroup, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, groupAddr, com.TypeGroup, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[12]=>poolf
	fmt.Println("set instances[12] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, poolfAddr, com.TypePoolF, nonce)
	signs = getSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, poolfAddr, com.TypePoolF, signs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	time.Sleep(waitTime1)
	time.Sleep(waitTime2)

	fmt.Println("*****************after set instance****************")

	// get instance
	tmp, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeAccess)
	fmt.Println("instance[3] access:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeERC20)
	fmt.Println("instance[4] erc20:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePoolP)
	fmt.Println("instance[5] poolp:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeRole)
	fmt.Println("instance[6] role:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeToken)
	fmt.Println("instance[7] token:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePledge)
	fmt.Println("instance[8] pledge:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeIssu)
	fmt.Println("instance[9] issuance:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.Typefs)
	fmt.Println("instance[10] filesys:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGroup)
	fmt.Println("instance[11] group:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePoolF)
	fmt.Println("instance[12] poolf:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeProxy)
	fmt.Println("instance[100] proxy:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeControl1)
	fmt.Println("instance[101] control1:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeControl2)
	fmt.Println("instance[102] control2:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
	fmt.Println("instance[150] getter:", tmp.Hex())

	fmt.Println()

	// admin approve poolp
	tx, err = erc20Ins.Approve(txAuth, poolpAddr, approveAmount)
	if err != nil {
		log.Fatal("approve poolp err:", err)
	}

	// admin approve poolf
	tx, err = erc20Ins.Approve(txAuth, poolfAddr, approveAmount)
	if err != nil {
		log.Fatal("approve poolf err:", err)
	}

	fmt.Println()

	// add token
	nonce, err = authIns.Nonce(&bind.CallOpts{From: com.AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("addT nonce:", nonce)
	hash = com.GetAddTHash(control1Addr, authAddr, erc20Addr, ismint, nonce)
	signs = getSigns(hash, sks)
	txAuth, _ = com.MakeAuth(ctx, eth, hexSk, txopts)
	tx, err = proxyIns.AddT(txAuth, erc20Addr, mint, signs)
	if err != nil {
		log.Fatal("addt err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after add token****************")
	fmt.Println("addT txHash:", tx.Hash())
	receipt := com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("add token tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get t cnt
	tcnt, err := tokenIns.GetTCnt(&bind.CallOpts{From: com.AdminAddr})
	if err != nil {
		log.Fatal("get tcnt err:", err)
	}
	fmt.Println("tcnt: ", tcnt)
	// get tinfo
	ta, err := tokenIns.GetTA(&bind.CallOpts{From: com.AdminAddr}, ti)
	fmt.Println("tokenAddr:", ta.Hex())
	fmt.Println("token right:", ta.Hex() == erc20Addr.Hex())

	fmt.Println()

	// create group
	tx, err = proxyIns.CreateGroup(txAuth, groupLevel, mr, dc, pc, kp, pp)
	if err != nil {
		log.Fatal("createGroup err:", err)
	}

	time.Sleep(waitTime2)

	// get gcnt
	gcnt, err := groupIns.GetGCnt(&bind.CallOpts{From: com.AdminAddr})
	if err != nil {
		log.Fatal("get ginfo err:", err)
	}
	fmt.Println("after creategroup, gcnt:", gcnt)

	// ban group
	nonce, err = authIns.Nonce(&bind.CallOpts{From: com.AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ban group nonce:", nonce)
	hash = com.GetBanGHash(control1Addr, authAddr, gi, isbanG, nonce)
	signs = getSigns(hash, sks)
	tx, err = proxyIns.BanG(txAuth, gi, banG, signs)
	if err != nil {
		log.Fatal("banG err:", err)
	}

	time.Sleep(waitTime2)

	// get ginfo
	isActive, isBan, err := groupIns.GetGInfo(&bind.CallOpts{From: com.AdminAddr}, gi)
	if err != nil {
		log.Fatal("get ginfo err:", err)
	}
	fmt.Println("after notban group, group isActive:", isActive, "\t", "isBan:", isBan)

	fmt.Println()

	// get kmanage
	kma, _ := groupIns.GetKManage(&bind.CallOpts{From: com.AdminAddr}, gi)
	fmt.Println("kmanage:", kma.Hex())
	kmanageIns, _ := kmanage.NewKmanage(kma, client)

	// add kmanage's owner=> control1
	fmt.Println("add kmanage's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(kma, control1Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = kmanageIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add kmanage's owner err:", err)
	}

	// add kmanage's owner=> control2
	fmt.Println("add kmanage's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(kma, control2Addr, authAddr, set, nonce)
	signs = getSigns(hash, sks)
	tx, err = kmanageIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add kmanage's owner err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println()

	// set ctl2 access to mint
	nonce, err = accessIns.Nonce(&bind.CallOpts{From: com.AdminAddr})
	fmt.Println("set access nonce:", nonce)
	hash = com.GetSetHash(accessAddr, control2Addr, nonce, set)
	signs = getSigns(hash, sks)
	txAuth, _ = com.MakeAuth(ctx, eth, hexSk, txopts)
	accessIns.Set(txAuth, control2Addr, set, signs)

	fmt.Println("*****************after set ctl2 access****************")

	time.Sleep(waitTime2)
	can, _ := accessIns.Access(&bind.CallOpts{From: com.AdminAddr}, control2Addr)
	fmt.Println("control2 access can:", can)

	fmt.Println()
}

func getSigns(hash []byte, sks [5]string) [5][]byte {
	res := [5][]byte{}

	for i := 0; i < 5; i++ {
		sign, err := com.Sign(hash, sks[i])
		if err != nil {
			fmt.Println("get signs err: ", err, "i: ", i)
			res[i] = nil
		} else {
			res[i] = sign
		}
	}

	return res
}
