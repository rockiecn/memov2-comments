package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/auth"
	"github.com/memoio/contractsv2/go_contracts/group"
	"github.com/memoio/contractsv2/go_contracts/kmanage"
	"github.com/memoio/contractsv2/go_contracts/proxy"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	eth   string
	hexSk string

	defaultGasLimit = uint64(8000000) // as small as possible
	pledgeAmount    = big.NewInt(int64(1e18))
	waitTime2       = 1 * time.Minute
	bigone          = big.NewInt(1)

	// AdminAddr admin address
	AdminAddr = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")

	groupLevel = uint8(7)
	mr         = uint8(4)
	dc         = uint8(3)
	pc         = uint8(2)

	ti     = uint8(0)
	isbanG = uint8(0) // false
	banG   = false
	set    = true
)

// 基于去掉control2中subOrder里面end需要与天对齐的限制进行测试
func main() {
	inputeth := flag.String("eth", "https://devchain.metamemo.one:8501", "eth api Address;") //dev net
	sk := flag.String("sk", "", "signature for sending transaction")
	flag.Parse()
	eth = *inputeth
	hexSk = *sk

	if hexSk == "" {
		fmt.Println("Please input your sk!")
		return
	}

	// get some value
	approveAmount := new(big.Int)
	rechargeAmount := new(big.Int)
	approveAmount, _ = approveAmount.SetString("1000000000000000000000", 10)
	rechargeAmount, _ = rechargeAmount.SetString("1000000000000000000000", 10)

	fmt.Println()

	// make auth to send transaction
	txopts := &com.TxOpts{Nonce: nil, GasPrice: nil, GasLimit: 0, Money: nil}
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

	// create group
	proxyIns, _ := proxy.NewProxy(com.ProxyAddr, client)
	tx, err := proxyIns.CreateGroup(txAuth, groupLevel, mr, dc, pc, pledgeAmount, pledgeAmount)
	if err != nil {
		log.Fatal("createGroup err:", err)
	}

	time.Sleep(waitTime2)

	// get create group receipt
	receipt := com.GetTransactionReceipt(eth, tx.Hash())
	fmt.Println("tx status: ", receipt.Status)
	if receipt.Status == 0 {
		log.Fatal("create group tx fail")
	}

	// get gcnt
	groupIns, _ := group.NewGroup(com.GroupAddr, client)
	gcnt, err := groupIns.GetGCnt(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal("get ginfo err:", err)
	}
	fmt.Println("after creategroup, gcnt:", gcnt)

	// get gi
	gi := gcnt - 1

	// ban group
	authIns, _ := auth.NewAuth(com.AuthAddr, client)
	nonce, err := authIns.Nonce(&bind.CallOpts{From: AdminAddr})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ban group nonce:", nonce)
	hash := com.GetBanGHash(com.Control1Addr, com.AuthAddr, gi, isbanG, nonce)
	signs, err := com.Sign(hash, hexSk)
	if err != nil {
		log.Fatal("sign banG err:", err)
	}
	tx, err = proxyIns.BanG(txAuth, gi, banG, [5][]byte{signs, signs, signs, signs, signs})
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
	hash = com.GetAddOwnerHash(kma, com.Control1Addr, com.AuthAddr, set, nonce)
	signs, err = com.Sign(hash, hexSk)
	if err != nil {
		log.Fatal("sign err:", err)
	}
	tx, err = kmanageIns.Add(txAuth, com.Control1Addr, set, [5][]byte{signs, signs, signs, signs, signs})
	if err != nil {
		log.Fatal("add kmanage's owner err:", err)
	}

	// add kmanage's owner=> control2
	fmt.Println("add kmanage's owner=>control2 nonce:", nonce.Add(nonce, bigone))
	hash = com.GetAddOwnerHash(kma, com.Control2Addr, com.AuthAddr, set, nonce)
	signs, err = com.Sign(hash, hexSk)
	if err != nil {
		log.Fatal("sign err:", err)
	}
	tx, err = kmanageIns.Add(txAuth, com.Control2Addr, set, [5][]byte{signs, signs, signs, signs, signs})
	if err != nil {
		log.Fatal("add kmanage's owner err:", err)
	}

	time.Sleep(waitTime2)

	fmt.Println()
}
