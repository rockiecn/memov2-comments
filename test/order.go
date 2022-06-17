package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/access"
	"github.com/memoio/contractsv2/go_contracts/auth"
	control "github.com/memoio/contractsv2/go_contracts/control1"
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
	eth                                    string
	hexSk, hexUsk, hexKsk, hexPsk, hexP2sk string
	sks                                    [5]string

	defaultGasLimit = uint64(8000000) // as small as possible
	pledgeAmount    = big.NewInt(int64(1e18))
	unpledgeAmount  = big.NewInt(int64(1e6))
	payAmount       = big.NewInt(1e4)
	waitTime1       = 30 * time.Second
	waitTime2       = 30 * time.Second
	bigone          = big.NewInt(1)

	// AdminAddr admin address
	AdminAddr = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
	// Foundation foundation address
	Foundation = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")

	// five admin addrs for multi-signature
	addrs = [5]common.Address{common.HexToAddress("0xB651b3ce6C7c712E41b03993a11DA656eF3C95A8"), common.HexToAddress("0x54a18833beaa8A22672ffF4e5b6714eCebC9368b"), common.HexToAddress("0x6f7bc6032278753c8A13bd430b13f1965407633f"), common.HexToAddress("0x09d186ce74f9092f5e6F20390d730943170bDbBA"), common.HexToAddress("0x8AadBf87Dc8Ae4A1161DEF8185D2b98C1904a3c3")}

	erc20Name   = "memo"
	erc20Symbol = "M"

	userAddr      = common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
	keeperAddr    = common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8")
	providerAddr  = common.HexToAddress("0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc")
	provider2Addr = common.HexToAddress("0x90f79bf6eb2c4f870365e785982e1f101e93b906")

	groupLevel = uint8(1)
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
	isLock = false

	start  = uint64(time.Now().Unix()) - 86400
	end    = start + 100*86400 + 86400 - start%86400 // 与天对齐
	size   = uint64(10 * 1024 * 1024 * 1024)         // 10GB
	sprice = big.NewInt(int64(size * 1))             // 1attomemo/(Byte*Second), just for test

	start2 = start + 10
	end2   = end

	subEnd = start + 86400 - start%86400

	lostAmount = new(big.Int).Mul(big.NewInt(int64(subEnd-start2)), sprice) // need more than (subEnd-start2)*sprice
)

func main() {
	inputeth := flag.String("eth", "http://119.147.213.220:8191", "eth api Address;") //dev net; http://119.147.213.220:8191; https://devchain.metamemo.one:8501
	// TODO: read sk from local config
	sk := flag.String("sk", "", "signature for sending transaction")
	usk := flag.String("usk", "", "user's sk")
	ksk := flag.String("ksk", "", "keeper's sk")
	psk := flag.String("psk", "", "provider's sk")
	p2sk := flag.String("p2sk", "", "provider2's sk")

	sk1 := flag.String("sk1", "", "addr1's sk")
	sk2 := flag.String("sk2", "", "addr2's sk")
	sk3 := flag.String("sk3", "", "addr3's sk")
	sk4 := flag.String("sk4", "", "addr4's sk")
	sk5 := flag.String("sk5", "", "addr5's sk")

	flag.Parse()
	eth = *inputeth
	hexSk, hexUsk, hexKsk, hexPsk, hexP2sk = *sk, *usk, *ksk, *psk, *p2sk

	sks = [5]string{*sk1, *sk2, *sk3, *sk4, *sk5}

	if hexSk == "" || hexUsk == "" || hexKsk == "" || hexPsk == "" || hexP2sk == "" {
		fmt.Println("Please input your sk!")
		return
	}

	// get some value
	approveAmount := new(big.Int)
	rechargeAmount := new(big.Int)
	approveAmount, _ = approveAmount.SetString("1000000000000000000000", 10)
	rechargeAmount, _ = rechargeAmount.SetString("1000000000000000000000", 10)

	fmt.Println()

	fmt.Println("start:", start, "\t", "end:", end, "\t", "duration:", end-start, "\t", "end align to day:", end%86400)
	ur := new(big.Int).Mul(big.NewInt(int64(end-start)), sprice)
	ur.Mul(ur, big.NewInt(105))
	ur.Div(ur, big.NewInt(100))
	fmt.Println("need pay:", ur)
	fmt.Println("user bal:", rechargeAmount)
	fmt.Println()

	// make auth to send transaction
	txopts := &com.TxOpts{Nonce: nil, GasPrice: nil, GasLimit: defaultGasLimit, Money: nil}
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
	accessAddr, tx, accessIns, err := access.DeployAccess(txAuth, client, addrs)
	if err != nil {
		log.Fatal("deploy access err:", err)
	}
	fmt.Println("accessAddr: ", accessAddr.Hex())
	_ = tx
	_ = accessIns

	// deploy ERC20
	erc20Addr, tx, erc20Ins, err := erc.DeployERC20(txAuth, client, accessAddr, erc20Name, erc20Symbol)
	if err != nil {
		log.Fatal("deploy erc20 err:", err)
	}
	fmt.Println("erc20Addr: ", erc20Addr.Hex())
	_ = erc20Ins

	// deploy Auth
	authAddr, tx, authIns, err := auth.DeployAuth(txAuth, client, addrs)
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
	control1Addr, tx, control1Ins, err := control.DeployControl1(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy control1 err:", err)
	}
	fmt.Println("control1Addr: ", control1Addr.Hex())
	_ = control1Ins

	time.Sleep(waitTime2)

	// deploy Control2
	control2Addr, tx, control2Ins, err := control2.DeployControl2(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy control2 err:", err)
	}
	fmt.Println("control2Addr: ", control2Addr.Hex())
	_ = control2Ins

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
	roleAddr, tx, roleIns, err := role.DeployRole(txAuth, client, authAddr, Foundation)
	if err != nil {
		log.Fatal("deploy role err:", err)
	}
	fmt.Println("roleAddr: ", roleAddr.Hex())
	_ = roleIns

	// deploy getter
	getterAddr, tx, getterIns, err := getter.DeployGetter(txAuth, client, authAddr, instanceAddr)
	if err != nil {
		log.Fatal("deploy getter err:", err)
	}
	fmt.Println("getterAddr: ", getterAddr.Hex())

	fmt.Println()

	// add Control1's owner: proxy; use 2 valid signature(should fail)
	nonce, err := authIns.Nonce(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	hash := com.GetAddOwnerHash(control1Addr, proxyAddr, authAddr, set, nonce)
	sks[1], sks[2], sks[3] = "", "", ""
	signs := com.GetSigns(hash, sks)
	tx, err = control1Ins.Add(txAuth, proxyAddr, set, signs)
	if err != nil {
		log.Fatal("add control1's owner err:", err)
	}
	time.Sleep(waitTime2)
	receipt := com.GetTransactionReceipt(eth, tx.Hash())
	if receipt.Status != 0 {
		log.Fatal("use 2 valid signature to add control1's owner should fail, but not")
	}

	// add Control1's owner: proxy; use 3 valid signature(should pass)
	sks[1] = *sk2
	fmt.Println("add Control1's owner=>proxy nonce:", nonce)
	hash = com.GetAddOwnerHash(control1Addr, proxyAddr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = control1Ins.Add(txAuth, proxyAddr, set, signs)
	if err != nil {
		log.Fatal("add control1's owner err:", err)
	}
	time.Sleep(waitTime2)
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt.Status != 1 {
		log.Fatal("use 3 valid signature to add control1's owner should pass, but fail")
	}

	// add Control2's owner: proxy
	sks[2], sks[3] = *sk3, *sk4
	fmt.Println("add Control2's owner=>proxy nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(control2Addr, proxyAddr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = control2Ins.Add(txAuth, proxyAddr, set, signs)
	if err != nil {
		log.Fatal("add control2's owner err:", err)
	}

	// add token's owner=> control1
	fmt.Println("add token's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(tokenAddr, control1Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = tokenIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add token's owner err:", err)
	}

	// add fs's owner=> control2
	fmt.Println("add fs's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(filesysAddr, control2Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = filesysIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add fs's owner err:", err)
	}

	// add group's owner=> control1
	fmt.Println("add group's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(groupAddr, control1Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = groupIns.Add0(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add group's owner err:", err)
	}

	// add issuance's owner=> control2
	fmt.Println("add issuance's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(issuanceAddr, control2Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = issuanceIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add issuance's owner err:", err)
	}

	// add poolp's owner=> control1
	fmt.Println("add poolp's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(poolpAddr, control1Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = poolpIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add poolp's owner err:", err)
	}

	// add pledge's owner=> control1
	fmt.Println("add pledge's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(pledgeAddr, control1Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = pledgeIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add pledge's owner err:", err)
	}

	// add poolf's owner=> control2
	fmt.Println("add poolf's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(poolfAddr, control2Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = poolfIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add poolf's owner err:", err)
	}

	// add role's owner=> control1
	fmt.Println("add role's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(roleAddr, control1Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
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
	nonce, err = authIns.Nonce(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("set instances[100] nonce:", nonce)
	hash = com.GetAlterHash(instanceAddr, authAddr, proxyAddr, com.TypeProxy, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, proxyAddr, com.TypeProxy, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[101]=> control1
	fmt.Println("set instances[101] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, control1Addr, com.TypeControl1, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, control1Addr, com.TypeControl1, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[102]=> control2
	fmt.Println("set instances[102] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, control2Addr, com.TypeControl2, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, control2Addr, com.TypeControl2, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[150]=> getter
	fmt.Println("set instances[150] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, getterAddr, com.TypeGetter, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, getterAddr, com.TypeGetter, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[5]=>poolp
	fmt.Println("set instances[5] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, poolpAddr, com.TypePoolP, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, poolpAddr, com.TypePoolP, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[6]=>role
	fmt.Println("set instances[6] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, roleAddr, com.TypeRole, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, roleAddr, com.TypeRole, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[7]=>token
	fmt.Println("set instances[7] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, tokenAddr, com.TypeToken, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, tokenAddr, com.TypeToken, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[8]=>pledge
	fmt.Println("set instances[8] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, pledgeAddr, com.TypePledge, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, pledgeAddr, com.TypePledge, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[9]=>issue
	fmt.Println("set instances[9] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, issuanceAddr, com.TypeIssu, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, issuanceAddr, com.TypeIssu, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[10]=>fs
	fmt.Println("set instances[10] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, filesysAddr, com.Typefs, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, filesysAddr, com.Typefs, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[11]=>group
	fmt.Println("set instances[11] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, groupAddr, com.TypeGroup, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, groupAddr, com.TypeGroup, signs)
	if err != nil {
		log.Fatal(err)
	}

	// set instances[12]=>poolf
	fmt.Println("set instances[12] nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAlterHash(instanceAddr, authAddr, poolfAddr, com.TypePoolF, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = instanceIns.Alter(txAuth, poolfAddr, com.TypePoolF, signs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	time.Sleep(waitTime2)

	fmt.Println("*****************after set instance****************")

	// get instance
	tmp, _ := instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypePoolP)
	fmt.Println("instance[5] poolp:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeRole)
	fmt.Println("instance[6] role:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeToken)
	fmt.Println("instance[7] token:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypePledge)
	fmt.Println("instance[8] pledge:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeIssu)
	fmt.Println("instance[9] issuance:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.Typefs)
	fmt.Println("instance[10] filesys:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeGroup)
	fmt.Println("instance[11] group:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypePoolF)
	fmt.Println("instance[12] poolf:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeControl1)
	fmt.Println("instance[101] control1:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeControl2)
	fmt.Println("instance[102] control2:", tmp.Hex())
	tmp, _ = instanceIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeGetter)
	fmt.Println("instance[150] getter:", tmp.Hex())

	fmt.Println()

	// get balance of admin
	erc20Ins, err = erc.NewERC20(erc20Addr, client)
	bal, err := erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, AdminAddr)
	if err != nil {
		log.Fatal("get bal err:", err)
	}
	fmt.Println("admin balance: ", bal)

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

	// user reAcc
	time.Sleep(waitTime2)
	txAuth, err = com.MakeAuth(ctx, eth, hexUsk, txopts)
	if err != nil {
		log.Fatal(" com.MakeAuth err:", err)
	}
	tx, err = proxyIns.ReAcc(txAuth)
	if err != nil {
		log.Fatal("user reacc err:", err)
	}
	// provider reAcc
	txAuth, err = com.MakeAuth(ctx, eth, hexPsk, txopts)
	if err != nil {
		log.Fatal(" com.MakeAuth err:", err)
	}
	tx, err = proxyIns.ReAcc(txAuth)
	if err != nil {
		log.Fatal("provider reacc err:", err)
	}
	// keeper reAcc
	txAuth, err = com.MakeAuth(ctx, eth, hexKsk, txopts)
	if err != nil {
		log.Fatal(" com.MakeAuth err:", err)
	}
	tx, err = proxyIns.ReAcc(txAuth)
	if err != nil {
		log.Fatal("keeper reacc err:", err)
	}

	fmt.Println("*****************after reAcc****************")

	// get user roleInfo
	time.Sleep(waitTime2)
	uri, err := roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, userAddr)
	if err != nil {
		log.Fatal("getRInfo err:", err)
	}
	roleout, _ := json.Marshal(uri)
	var out bytes.Buffer
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("user=%v\n", out.String())
	// get provider roleInfo
	out.Reset()
	pri, err := roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, providerAddr)
	if err != nil {
		log.Fatal("getRInfo err:", err)
	}
	roleout, _ = json.Marshal(pri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("provider=%v\n", out.String())
	// get keeper roleInfo
	out.Reset()
	kri, err := roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, keeperAddr)
	if err != nil {
		log.Fatal("getRInfo err:", err)
	}
	roleout, _ = json.Marshal(kri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("keeper=%v\n", out.String())

	fmt.Println()

	// user reRole
	txAuth, _ = com.MakeAuth(ctx, eth, hexUsk, txopts)
	tx, err = proxyIns.ReRole(txAuth, com.UserType, nil)
	if err != nil {
		log.Fatal("user reRole err:", err)
	}

	// keeper reRole
	txAuth, _ = com.MakeAuth(ctx, eth, hexKsk, txopts)
	tx, err = proxyIns.ReRole(txAuth, com.KeeperType, nil)
	if err != nil {
		log.Fatal("keeper reRole err:", err)
	}

	// provider reRole
	txAuth, _ = com.MakeAuth(ctx, eth, hexPsk, txopts)
	tx, err = proxyIns.ReRole(txAuth, com.ProviderType, nil)
	if err != nil {
		log.Fatal("provider reRole err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after reRole****************")

	// get user roleInfo
	uri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, userAddr)
	if err != nil {
		log.Fatal("getRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(uri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("user=%v\n", out.String())
	// get provider roleInfo
	pri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, providerAddr)
	if err != nil {
		log.Fatal("getRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(pri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("provider=%v\n", out.String())
	// get keeper roleInfo
	kri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, keeperAddr)
	if err != nil {
		log.Fatal("getRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(kri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("keeper=%v\n", out.String())

	fmt.Println()

	// add token
	nonce, err = authIns.Nonce(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("addT nonce:", nonce)
	hash = com.GetAddTHash(control1Addr, authAddr, erc20Addr, ismint, nonce)
	signs = com.GetSigns(hash, sks)
	txAuth, _ = com.MakeAuth(ctx, eth, hexSk, txopts)
	tx, err = proxyIns.AddT(txAuth, erc20Addr, mint, signs)
	if err != nil {
		log.Fatal("addt err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after add token****************")
	fmt.Println("addT txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("add token tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get t cnt
	tcnt, err := tokenIns.GetTCnt(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal("get tcnt err:", err)
	}
	fmt.Println("tcnt: ", tcnt)
	// get tinfo
	ta, err := tokenIns.GetTA(&bind.CallOpts{From: AdminAddr}, ti)
	fmt.Println("tokenAddr:", ta.Hex())
	fmt.Println("token right:", ta.Hex() == erc20Addr.Hex())

	fmt.Println()

	// keeper pledge
	tx, err = proxyIns.Pledge(txAuth, kri.Index, pledgeAmount)
	if err != nil {
		log.Fatal("k pledge err:", err)
	}

	// provider pledge
	tx, err = proxyIns.Pledge(txAuth, pri.Index, pledgeAmount)
	if err != nil {
		log.Fatal("p pledge err:", err)
	}

	// user pledge
	tx, err = proxyIns.Pledge(txAuth, uri.Index, pledgeAmount)
	if err != nil {
		log.Fatal("u pledge err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after pledge kp****************")

	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("pledge k bal:", bal)
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("pledge p bal:", bal)
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("pledge u bal:", bal)
	// get reward info
	trewardinfo, _ := pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ := pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ := pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())
	// get totalpledge
	tp, _ := pledgeIns.TotalPledge(&bind.CallOpts{From: AdminAddr})
	fmt.Println("totalPledge:", tp)
	ureward, _ := pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(ureward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("uReward=%v\n", out.String())

	fmt.Println()

	// create group
	tx, err = proxyIns.CreateGroup(txAuth, groupLevel, mr, dc, pc, pledgeAmount, pledgeAmount)
	if err != nil {
		log.Fatal("createGroup err:", err)
	}

	time.Sleep(waitTime2)

	// get gcnt
	gcnt, err := groupIns.GetGCnt(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal("get ginfo err:", err)
	}
	fmt.Println("after creategroup, gcnt:", gcnt)

	// ban group
	nonce, err = authIns.Nonce(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ban group nonce:", nonce)
	hash = com.GetBanGHash(control1Addr, authAddr, gi, isbanG, nonce)
	signs = com.GetSigns(hash, sks)
	if err != nil {
		log.Fatal("sign banG err:", err)
	}
	tx, err = proxyIns.BanG(txAuth, gi, banG, signs)
	if err != nil {
		log.Fatal("banG err:", err)
	}

	time.Sleep(waitTime2)

	// get ginfo
	isActive, isBan, err := groupIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi)
	if err != nil {
		log.Fatal("get ginfo err:", err)
	}
	fmt.Println("after notban group, group isActive:", isActive, "\t", "isBan:", isBan)

	fmt.Println()

	// get kmanage
	kma, _ := groupIns.GetKManage(&bind.CallOpts{From: AdminAddr}, gi)
	fmt.Println("kmanage:", kma.Hex())
	kmanageIns, _ := kmanage.NewKmanage(kma, client)

	// add kmanage's owner=> control1
	fmt.Println("add kmanage's owner=>control1 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(kma, control1Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = kmanageIns.Add(txAuth, control1Addr, set, signs)
	if err != nil {
		log.Fatal("add kmanage's owner err:", err)
	}

	// add kmanage's owner=> control2
	fmt.Println("add kmanage's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(kma, control2Addr, authAddr, set, nonce)
	signs = com.GetSigns(hash, sks)
	tx, err = kmanageIns.Add(txAuth, control2Addr, set, signs)
	if err != nil {
		log.Fatal("add kmanage's owner err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println()

	// keeper addToGroup
	txAuth, _ = com.MakeAuth(ctx, eth, hexKsk, txopts)
	tx, err = proxyIns.AddToGroup(txAuth, gi)
	if err != nil {
		log.Fatal("addToGroup err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after addToGroup****************")
	// get kinfo
	kri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, keeperAddr)
	if err != nil {
		log.Fatal("getkRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(kri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("keeper=%v\n", out.String())

	fmt.Println()

	// activate keeper
	nonce, err = authIns.Nonce(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("activate k nonce:", nonce)
	hash = com.GetActivateHash(control1Addr, authAddr, kri.Index, nonce)
	signs = com.GetSigns(hash, sks)
	txAuth, _ = com.MakeAuth(ctx, eth, hexSk, txopts)
	tx, err = proxyIns.Activate(txAuth, kri.Index, signs)
	if err != nil {
		log.Fatal("activate err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after activate****************")

	// get kinfo
	kri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, keeperAddr)
	if err != nil {
		log.Fatal("getkRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(kri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("keeper=%v\n", out.String())

	// get ginfo
	isActive, isBan, err = groupIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi)
	if err != nil {
		log.Fatal("get ginfo err:", err)
	}
	fmt.Println("group isActive:", isActive, "\t", "isBan:", isBan)

	fmt.Println()

	// user addToGroup
	txAuth, _ = com.MakeAuth(ctx, eth, hexUsk, txopts)
	tx, err = proxyIns.AddToGroup(txAuth, gi)
	if err != nil {
		log.Fatal("user addToGroup err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println("*****************after addToGroup****************")

	// get uinfo
	uri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, userAddr)
	if err != nil {
		log.Fatal("getuRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(uri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("user=%v\n", out.String())

	fmt.Println()

	// provider addToGroup
	txAuth, _ = com.MakeAuth(ctx, eth, hexPsk, txopts)
	tx, err = proxyIns.AddToGroup(txAuth, gi)
	if err != nil {
		log.Fatal("provider addToGroup err:", err)
	}

	// get pinfo
	time.Sleep(waitTime2)
	pri, err = roleIns.GetRInfo(&bind.CallOpts{From: AdminAddr}, providerAddr)
	if err != nil {
		log.Fatal("getpRInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(pri)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("provider=%v\n", out.String())

	fmt.Println()

	// get admin approve poolf、balance
	allo, err := erc20Ins.Allowance(&bind.CallOpts{From: AdminAddr}, AdminAddr, poolfAddr)
	fmt.Println("admin approve poolf: ", allo)
	bal, err = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, AdminAddr)
	fmt.Println("admin       balance: ", bal)
	fmt.Println("admin recharge user: ", rechargeAmount)

	// get tinfo[ti] isBan==false, uinfo[ui] isBan==false
	tmp, err = tokenIns.GetTA(&bind.CallOpts{From: AdminAddr}, ti)
	fmt.Println("token equal:", tmp.Hex() == erc20Addr.Hex())
	tindex, _isban, _m, err := tokenIns.GetTI(&bind.CallOpts{From: AdminAddr}, erc20Addr)
	fmt.Println("ti:", tindex, " isBan:", _isban, " mint:", _m)

	// user recharge
	txAuth, _ = com.MakeAuth(ctx, eth, hexSk, txopts)
	tx, err = proxyIns.Recharge(txAuth, uri.Index, ti, isLock, rechargeAmount)
	if err != nil {
		log.Fatal("user recharge err:", err)
	}

	fmt.Println("*****************after recharge****************")
	time.Sleep(waitTime2)
	time.Sleep(waitTime1)
	// get user balance in fs
	avail, lock, err := filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	if err != nil {
		log.Fatal("get fs user bal err:", err)
	}
	fmt.Println("get fs user avail:", avail, " lock:", lock)

	// get keeper bal in fs
	avail, lock, err = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	if err != nil {
		log.Fatal("get fs keeper bal err:", err)
	}
	fmt.Println("get fs keeper avail:", avail, " lock:", lock)

	// get foundation bal
	avail, lock, err = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, 0, ti)
	if err != nil {
		log.Fatal("get fs foundation bal err:", err)
	}
	fmt.Println("get fs foundation avail:", avail, " lock:", lock)

	fmt.Println()

	// set ctl2 access to mint
	nonce, err = accessIns.Nonce(&bind.CallOpts{From: AdminAddr})
	fmt.Println("set access nonce:", nonce)
	hash = com.GetSetHash(accessAddr, control2Addr, nonce, set)
	signs = com.GetSigns(hash, sks)
	if err != nil {
		log.Fatal("set access err:", err)
	}
	txAuth, _ = com.MakeAuth(ctx, eth, hexSk, txopts)
	accessIns.Set(txAuth, control2Addr, set, signs)

	fmt.Println("*****************after set ctl2 access****************")

	time.Sleep(waitTime2)
	can, _ := accessIns.Access(&bind.CallOpts{From: AdminAddr}, control2Addr)
	fmt.Println("control2 access can:", can)

	fmt.Println()

	// user get fs nonce
	fsinfo, err := filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index)
	if err != nil {
		log.Fatal("get fsinfo err:", err)
	}
	fmt.Println("fsinfo nonce:", fsinfo.Nonce, "\t", "subNonce:", fsinfo.SubNonce)

	// user add order
	txAuth, _ = com.MakeAuth(ctx, eth, hexUsk, txopts)
	oi := proxy.OrderIn{
		UIndex: uri.Index,
		PIndex: pri.Index,
		Start:  start,
		End:    end,
		Size:   size,
		Nonce:  fsinfo.Nonce,
		TIndex: ti,
		Sprice: sprice,
	}
	// abi.encodePacked(uindex,pindex,nonce,start,end,size,tindex,sprice)
	hash = com.GetAddOrderHash(oi.UIndex, oi.PIndex, start, end, size, fsinfo.Nonce, ti, sprice)
	usign, _ := com.Sign(hash, hexUsk)
	pSign, _ := com.Sign(hash, hexPsk)
	tx, err = proxyIns.AddOrder(txAuth, oi, usign, pSign)
	if err != nil {
		log.Fatal("addOrder err:", err)
	}
	fmt.Println()

	fmt.Println("*****************after add order****************")

	time.Sleep(waitTime2)
	fmt.Println("addOrder txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("add order tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get fs storeinfo
	fsStore, err := filesysIns.GetStoreInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index, ti)
	if err != nil {
		log.Fatal("get fa storeInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(fsStore)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("fsStore=%v\n", out.String())

	// get fs groupinfo
	ginfo, _ := filesysIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi, ti)
	out.Reset()
	roleout, _ = json.Marshal(ginfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("ginfo=%v\n", out.String())

	// get fs info
	finfo, _ := filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index)
	out.Reset()
	roleout, _ = json.Marshal(finfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("finfo=%v\n", out.String())

	// get settleinfo
	seinfo, _ := filesysIns.GetSettleInfo(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(seinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("seinfo=%v\n", out.String())

	// get balance of user、foundation、provider
	ubal, ulock, _ := filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("user bal in fs:", ubal, ulock)
	fbal, flock, _ := filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, 0, ti)
	fmt.Println("foundation bal in fs:", fbal, flock)
	pbal, plock, _ := filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("provider bal in fs:", pbal, plock)

	// get issuance rewardResid、rewardTarget、and so on
	rewardresid, _ := issuanceIns.RewardResid(&bind.CallOpts{From: AdminAddr})
	tr, _ := issuanceIns.TargetReward(&bind.CallOpts{From: AdminAddr})
	fmt.Println("issuance rewardresid:", rewardresid, "\t", "targetreward:", tr, "\t", "reward:", tr.Sub(tr, rewardresid))
	s, _ := issuanceIns.Size(&bind.CallOpts{From: AdminAddr})
	fmt.Println("issuance size:", s)

	// get poolp balance
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, poolpAddr)
	fmt.Println("poolp balance:", bal)

	// get reward info
	trewardinfo, _ = pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())

	fmt.Println()

	// user addorder again
	// user get fs nonce
	fsinfo, err = filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index)
	if err != nil {
		log.Fatal("get fsinfo err:", err)
	}
	fmt.Println("fsinfo nonce:", fsinfo.Nonce, "\t", "subNonce:", fsinfo.SubNonce)

	// user add order again
	txAuth, _ = com.MakeAuth(ctx, eth, hexUsk, txopts)
	oi = proxy.OrderIn{
		UIndex: uri.Index,
		PIndex: pri.Index,
		Start:  start2,
		End:    end2,
		Size:   size,
		Nonce:  fsinfo.Nonce,
		TIndex: ti,
		Sprice: sprice,
	}
	// abi.encodePacked(uindex,pindex,nonce,start,end,size,tindex,sprice)
	hash = com.GetAddOrderHash(oi.UIndex, oi.PIndex, oi.Start, oi.End, oi.Size, oi.Nonce, oi.TIndex, oi.Sprice)
	usign, _ = com.Sign(hash, hexUsk)
	pSign, _ = com.Sign(hash, hexPsk)
	tx, err = proxyIns.AddOrder(txAuth, oi, usign, pSign)
	if err != nil {
		log.Fatal("addOrder err:", err)
	}
	fmt.Println()

	fmt.Println("*****************after add order again****************")

	time.Sleep(waitTime2)
	fmt.Println("addOrder txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("add order tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get fs storeinfo
	fsStore, err = filesysIns.GetStoreInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index, ti)
	if err != nil {
		log.Fatal("get fa storeInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(fsStore)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("fsStore=%v\n", out.String())

	// get fs groupinfo
	ginfo, _ = filesysIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi, ti)
	out.Reset()
	roleout, _ = json.Marshal(ginfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("ginfo=%v\n", out.String())

	// get fs info
	finfo, _ = filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index)
	out.Reset()
	roleout, _ = json.Marshal(finfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("finfo=%v\n", out.String())

	// get settleinfo
	seinfo, _ = filesysIns.GetSettleInfo(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(seinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("seinfo=%v\n", out.String())

	// get balance of user、foundation、provider
	ubal, ulock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("user bal in fs:", ubal, ulock)
	fbal, flock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, 0, ti)
	fmt.Println("foundation bal in fs:", fbal, flock)
	pbal, plock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("provider bal in fs:", pbal, plock)

	// get issuance rewardResid、rewardTarget、and so on
	rewardresid, _ = issuanceIns.RewardResid(&bind.CallOpts{From: AdminAddr})
	tr, _ = issuanceIns.TargetReward(&bind.CallOpts{From: AdminAddr})
	fmt.Println("issuance rewardresid:", rewardresid, "\t", "targetreward:", tr, "\t", "reward:", tr.Sub(tr, rewardresid))
	s, _ = issuanceIns.Size(&bind.CallOpts{From: AdminAddr})
	fmt.Println("issuance size:", s)

	// get poolp balance
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, poolpAddr)
	fmt.Println("poolp balance:", bal)

	// get reward info
	trewardinfo, _ = pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())
	ureward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(ureward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("uReward=%v\n", out.String())

	fmt.Println()

	// get fs info
	fsinfo, err = filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index)
	if err != nil {
		log.Fatal("get fsinfo err:", err)
	}
	fmt.Println("fsinfo nonce:", fsinfo.Nonce, "\t", "subNonce:", fsinfo.SubNonce)
	// user subOrder
	txAuth, _ = com.MakeAuth(ctx, eth, hexKsk, txopts)
	oi = proxy.OrderIn{
		UIndex: uri.Index,
		PIndex: pri.Index,
		Start:  start,
		End:    subEnd,
		Size:   size,
		Nonce:  fsinfo.SubNonce,
		TIndex: ti,
		Sprice: sprice,
	}
	// abi.encodePacked(uindex,pindex,nonce,start,end,size,tindex,sprice)
	hash = com.GetAddOrderHash(oi.UIndex, oi.PIndex, oi.Start, oi.End, oi.Size, oi.Nonce, oi.TIndex, oi.Sprice)
	usign, _ = com.Sign(hash, hexUsk)
	pSign, _ = com.Sign(hash, hexPsk)
	tx, err = proxyIns.SubOrder(txAuth, oi, usign, pSign)
	if err != nil {
		log.Fatal("subOrder err:", err)
	}

	fmt.Println()

	fmt.Println("*****************after sub order****************")

	time.Sleep(waitTime2)
	fmt.Println("subOrder txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("sub order tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get fs storeinfo
	fsStore, err = filesysIns.GetStoreInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index, ti)
	if err != nil {
		log.Fatal("get fa storeInfo err:", err)
	}
	out.Reset()
	roleout, _ = json.Marshal(fsStore)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("fsStore=%v\n", out.String())

	// get fs groupinfo
	ginfo, _ = filesysIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi, ti)
	out.Reset()
	roleout, _ = json.Marshal(ginfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("ginfo=%v\n", out.String())

	// get fs info
	finfo, _ = filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, uri.Index, pri.Index)
	out.Reset()
	roleout, _ = json.Marshal(finfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("finfo=%v\n", out.String())

	// get settleinfo
	seinfo, _ = filesysIns.GetSettleInfo(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(seinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("seinfo=%v\n", out.String())

	// get balance of user、foundation、provider、keeper
	ubal, ulock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("user bal in fs:", ubal, ulock)
	fbal, flock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, 0, ti)
	fmt.Println("foundation bal in fs:", fbal, flock)
	pbal, plock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("provider bal in fs:", pbal, plock)
	kbal, klock, _ := filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("keeper bal in fs:", kbal, klock)

	// get issuance rewardResid、rewardTarget、and so on
	rewardresid, _ = issuanceIns.RewardResid(&bind.CallOpts{From: AdminAddr})
	tr, _ = issuanceIns.TargetReward(&bind.CallOpts{From: AdminAddr})
	fmt.Println("issuance rewardresid:", rewardresid, "\t", "targetreward:", tr, "\t", "reward:", tr.Sub(tr, rewardresid))
	s, _ = issuanceIns.Size(&bind.CallOpts{From: AdminAddr})
	fmt.Println("issuance size:", s)

	// get poolp balance
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, poolpAddr)
	fmt.Println("poolp balance:", bal)

	// get kmanage info
	last, profit, _ := kmanageIns.GetPf(&bind.CallOpts{From: AdminAddr}, ti)
	fmt.Println("kmanage last get reward:", last, "\t", "total reward acculumate:", profit)
	avail, tmpBal, _ := kmanageIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("kmanage keeper avail:", avail, "\t", "tmp:", tmpBal)

	fmt.Println()

	// get reward info
	trewardinfo, _ = pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())

	fmt.Println()

	// get provider balance in pledge
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("p balance in pledge:", bal)
	// get totalpledge
	tp, _ = pledgeIns.TotalPledge(&bind.CallOpts{From: AdminAddr})
	fmt.Println("totalPledge:", tp)
	// p withdraw from pledge
	txAuth, _ = com.MakeAuth(ctx, eth, hexPsk, txopts)
	tx, err = proxyIns.Unpledge(txAuth, pri.Index, ti, unpledgeAmount)
	if err != nil {
		log.Fatal("unpledge err:", err)
	}

	fmt.Println()

	fmt.Println("*****************after p unpledge****************")

	time.Sleep(waitTime2)
	fmt.Println("p unpledge txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("p unpledge tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get p bal
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, providerAddr)
	fmt.Println("p bal:", bal)
	// get provider balance in pledge
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("p balance in pledge:", bal)
	// get poolp bal
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, poolpAddr)
	fmt.Println("poolp bal:", bal)

	// get reward info
	trewardinfo, _ = pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())

	// get totalpledge
	tp, _ = pledgeIns.TotalPledge(&bind.CallOpts{From: AdminAddr})
	fmt.Println("totalPledge:", tp)

	fmt.Println()

	// get keeper balance in pledge
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("keeper balance in pledge:", bal)
	// keeper withdraw from pledge
	txAuth, _ = com.MakeAuth(ctx, eth, hexKsk, txopts)
	tx, err = proxyIns.Unpledge(txAuth, kri.Index, ti, unpledgeAmount)
	if err != nil {
		log.Fatal("unpledge err:", err)
	}

	fmt.Println()

	fmt.Println("*****************after keeper unpledge****************")

	time.Sleep(waitTime2)
	fmt.Println("k unpledge txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("k unpledge tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get keeper bal
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, keeperAddr)
	fmt.Println("keeper bal:", bal)
	// get keeper balance in pledge
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("keeper balance in pledge:", bal)
	// get poolp bal
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, poolpAddr)
	fmt.Println("poolp bal:", bal)

	// get reward info
	trewardinfo, _ = pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())

	// get totalpledge
	tp, _ = pledgeIns.TotalPledge(&bind.CallOpts{From: AdminAddr})
	fmt.Println("totalPledge:", tp)

	fmt.Println()

	// get user balance in pledge
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("user balance in pledge:", bal)
	// user withdraw from pledge
	txAuth, _ = com.MakeAuth(ctx, eth, hexUsk, txopts)
	tx, err = proxyIns.Unpledge(txAuth, uri.Index, ti, unpledgeAmount)
	if err != nil {
		log.Fatal("unpledge err:", err)
	}

	fmt.Println()

	fmt.Println("*****************after user unpledge****************")

	time.Sleep(waitTime2)
	fmt.Println("u unpledge txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("u unpledge tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get user bal
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, userAddr)
	fmt.Println("user bal:", bal)
	// get user balance in pledge
	bal, _ = pledgeIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("user balance in pledge:", bal)
	// get poolp bal
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, poolpAddr)
	fmt.Println("poolp bal:", bal)

	// get reward info
	trewardinfo, _ = pledgeIns.TInfo(&bind.CallOpts{From: AdminAddr}, ti)
	out.Reset()
	roleout, _ = json.Marshal(trewardinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("tRewardInfo=%v\n", out.String())
	preward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(preward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("pReward=%v\n", out.String())
	kreward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(kreward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("kReward=%v\n", out.String())
	ureward, _ = pledgeIns.Rewards(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(ureward)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("uReward=%v\n", out.String())

	// get totalpledge
	tp, _ = pledgeIns.TotalPledge(&bind.CallOpts{From: AdminAddr})
	fmt.Println("totalPledge:", tp)

	fmt.Println()

	// provider prowithdraw from fs
	txAuth, _ = com.MakeAuth(ctx, eth, hexPsk, txopts)
	pw := proxy.PWIn{
		PIndex: pri.Index,
		TIndex: ti,
		Pay:    payAmount,
		Lost:   lostAmount,
	}
	hash = com.GetProWithdrawHash(pw.PIndex, pw.TIndex, pw.Pay, pw.Lost)
	ksign, _ := com.Sign(hash, hexKsk)
	tx, err = proxyIns.ProWithdraw(txAuth, pw, []uint64{kri.Index}, [][]byte{ksign})
	if err != nil {
		log.Fatal("p prowithdraw err:", err)
	}

	fmt.Println("*****************after prowithdraw****************")

	time.Sleep(waitTime2)
	fmt.Println("prowithdraw txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("prowithdraw tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get kmanage profit
	last, acc, _ := kmanageIns.GetPf(&bind.CallOpts{From: AdminAddr}, ti)
	fmt.Println("last release profit:", last, "\t", "acc:", acc)
	// get k bal in kmanage
	avail, tmpb, _ := kmanageIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("k avail:", avail, "\t", "tmp:", tmpb)
	// seinfo
	seinfo, _ = filesysIns.GetSettleInfo(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	roleout, _ = json.Marshal(seinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("seinfo=%v\n", out.String())
	// ginfo
	ginfo, _ = filesysIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi, ti)
	out.Reset()
	roleout, _ = json.Marshal(ginfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("ginfo=%v\n", out.String())
	// balance[pi]
	avail, lock, _ = filesysIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	fmt.Println("provider avail:", avail, "\t", "provider lock:", lock)

	fmt.Println()

	// provider withdraw
	tx, err = proxyIns.Withdraw(txAuth, pri.Index, ti, payAmount)
	if err != nil {
		log.Fatal("provider withdraw err:", err)
	}

	fmt.Println("*****************after provider withdraw****************")

	time.Sleep(waitTime2)
	fmt.Println("provider withdraw txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("provider withdraw tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get provider balance
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, providerAddr)
	fmt.Println("provider balance in erc20:", bal)

	fmt.Println()

	// keeper withdraw from fs
	txAuth, _ = com.MakeAuth(ctx, eth, hexKsk, txopts)
	tx, err = proxyIns.Withdraw(txAuth, kri.Index, ti, payAmount)
	if err != nil {
		log.Fatal("keeper withdraw err:", err)
	}

	fmt.Println("*****************after keeper withdraw****************")

	time.Sleep(waitTime2)
	fmt.Println("keeper withdraw txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("keeper withdraw tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)
	// get kmanage profit
	last, acc, _ = kmanageIns.GetPf(&bind.CallOpts{From: AdminAddr}, ti)
	fmt.Println("last release profit:", last, "\t", "acc:", acc)
	// get k bal in kmanage
	avail, tmpb, _ = kmanageIns.BalanceOf(&bind.CallOpts{From: AdminAddr}, kri.Index, ti)
	fmt.Println("k avail:", avail, "\t", "tmp:", tmpb)
	// get keeper balance in fs
	bal, _ = erc20Ins.BalanceOf(&bind.CallOpts{From: AdminAddr}, keeperAddr)
	fmt.Println("keeper balance in erc20:", bal)

	fmt.Println()

	// get repair fs nonce
	fsinfo, err = filesysIns.GetFsInfo(&bind.CallOpts{From: AdminAddr}, 0, pri.Index)
	if err != nil {
		log.Fatal("get fsinfo err:", err)
	}
	fmt.Println("fsinfo nonce:", fsinfo.Nonce, "\t", "subNonce:", fsinfo.SubNonce)

	// provider add repair
	txAuth, _ = com.MakeAuth(ctx, eth, hexPsk, txopts)
	oi = proxy.OrderIn{
		UIndex: 0,
		PIndex: pri.Index,
		Start:  start2,
		End:    subEnd,
		Size:   size,
		Nonce:  fsinfo.Nonce,
		TIndex: ti,
		Sprice: sprice,
	}
	hash = com.GetAddOrderHash(oi.UIndex, oi.PIndex, oi.Start, oi.End, oi.Size, oi.Nonce, oi.TIndex, oi.Sprice)
	ksign, _ = com.Sign(hash, hexKsk)
	tx, err = proxyIns.AddRepair(txAuth, oi, []uint64{kri.Index}, [][]byte{ksign})
	if err != nil {
		log.Fatal("keeper withdraw err:", err)
	}

	fmt.Println("*****************after provider addRepair****************")

	time.Sleep(waitTime2)
	fmt.Println("provider addrepair txHash:", tx.Hash())
	receipt = com.GetTransactionReceipt(eth, tx.Hash())
	if receipt == nil {
		log.Fatal("provider addrepair tx receipt is nil")
	}
	fmt.Println("tx status:", receipt.Status)
	fmt.Println("tx gas:", receipt.GasUsed)

	// get ginfo
	ginfo, _ = filesysIns.GetGInfo(&bind.CallOpts{From: AdminAddr}, gi, ti)
	out.Reset()
	roleout, _ = json.Marshal(ginfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("ginfo=%v\n", out.String())

	// get seinfo
	seinfo, _ = filesysIns.GetSettleInfo(&bind.CallOpts{From: AdminAddr}, pri.Index, ti)
	out.Reset()
	roleout, _ = json.Marshal(seinfo)
	json.Indent(&out, roleout, "", "\t")
	fmt.Printf("seinfo=%v\n", out.String())

	fmt.Println()

	// --------test getter--------
	tAddr, _ := getterIns.GetTA(&bind.CallOpts{From: AdminAddr}, ti)
	fmt.Println("ta:", tAddr, "\t", "right:", tAddr.Hex() == erc20Addr.Hex())
	avail, lock, _ = getterIns.FsBalanceOf(&bind.CallOpts{From: AdminAddr}, uri.Index, ti)
	fmt.Println("user fs avail:", avail, "\t", "lock:", lock)
	gcnt, _ = getterIns.GetGCnt(&bind.CallOpts{From: AdminAddr})
	fmt.Println("gcnt:", gcnt)
	manageRate, _ := getterIns.ManageRate(&bind.CallOpts{From: AdminAddr}, gi)
	fmt.Println("kmanage manageRate:", manageRate)
	ia, _ := getterIns.Instances(&bind.CallOpts{From: AdminAddr}, com.TypeIssu)
	fmt.Println("issuance addr:", ia.Hex(), "\t", "right:", ia.Hex() == issuanceAddr.Hex())
	tp, _ = getterIns.GetTotalPledge(&bind.CallOpts{From: AdminAddr})
	fmt.Println("pledge totalPledge:", tp)
	ui, _ := getterIns.GetAcc(&bind.CallOpts{From: AdminAddr}, userAddr)
	fmt.Println("user index:", ui, "\t", "right:", ui == uri.Index)
}
