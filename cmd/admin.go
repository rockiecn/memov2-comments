package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/auth"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
	"github.com/memoio/contractsv2/go_contracts/proxy"
	"github.com/urfave/cli/v2"
)

var (
	WaitTime = 30 * time.Second
)

// AdminCmd is about contracts functions called by admin
// admin value is optional. [completed]
var AdminCmd = &cli.Command{
	Name:  "admin",
	Usage: "Admin call contracts",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "adminAddr",
			Aliases: []string{"admin"},
			Value:   com.AdminAddr.Hex(),
			Usage:   "the admin account's address",
		},
		&cli.StringFlag{
			Name:    "adminSk",
			Aliases: []string{"as"},
			Value:   "", //默认值为空，需手动输入
			Usage:   "the admin account's secretkey",
		},
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   "dev", //默认值为devchain
			Usage:   "indicate which chain(dev, test, product), default dev",
		},
		&cli.StringSliceFlag{
			Name:    "5-authorizer-privateKey",
			Aliases: []string{"sks"},
			Value:   nil,
			Usage:   "5 authorizers'2 private key, input as: --sks xxx --sks xxx --sks xxx ...",
		},
	},
	Subcommands: []*cli.Command{
		activateCmd,
		banACmd,
		banGCmd,
		banTCmd,
		addTCmd,
	},
}

var activateCmd = &cli.Command{
	Name:      "actK",
	Usage:     "Admin activate keeper after keeper addToGroup.",
	ArgsUsage: "<keeper index and authorizers' sk>",
	Description: `
A function of Proxy contract.
Admin activate keeper after keeper add to group.

Arguments:
ki - the keeper index which admin will activate
	`,
	Action: func(cctx *cli.Context) error {
		kis := cctx.Args().Get(0)
		ki, _ := strconv.Atoi(kis)
		fmt.Println("ki:", ki)

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())

		sksSli := cctx.StringSlice("sks")
		fmt.Println("sks length:", len(sksSli))
		fmt.Println("5 sks:", sksSli)
		sks := [5]string{}
		copy(sks[:], sksSli)

		txopts := &com.TxOpts{
			Nonce:    nil,
			GasPrice: nil,
			GasLimit: 0,
		}

		txAuth, err := com.MakeAuth(cctx.Context, endPoint, sk, txopts)
		if err != nil {
			log.Fatal(err)
		}

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new proxyIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		proxyAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeProxy)
		if err != nil {
			log.Fatal(err)
		}
		proxyIns, _ := proxy.NewProxy(proxyAddr, client)

		// get control1Addr
		control1Addr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeControl1)
		if err != nil {
			log.Fatal(err)
		}

		// new authIns
		authAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeAuth)
		if err != nil {
			log.Fatal(err)
		}
		authIns, _ := auth.NewAuth(authAddr, client)

		// get nonce
		nonce, err := authIns.Nonce(&bind.CallOpts{From: addr})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("activate k nonce:", nonce)
		hash := com.GetActivateHash(control1Addr, authAddr, uint64(ki), nonce)
		signs := com.GetSigns(hash, sks)

		tx, err := proxyIns.Activate(txAuth, uint64(ki), signs)
		if err != nil {
			log.Fatal("activate keeper err:", err)
		}
		fmt.Println("tx hash:", tx.Hash())

		fmt.Println("waiting tx complete...")

		time.Sleep(WaitTime)
		receipt := com.GetTransactionReceipt(endPoint, tx.Hash())
		fmt.Println("tx status:", receipt.Status)

		return nil
	},
}

var banACmd = &cli.Command{
	Name:      "banA",
	Usage:     "Admin ban or not ban account.",
	ArgsUsage: "<account's index>",
	Description: `
A function of Proxy contract.
Admin ban account.

Arguments:
ri - the account's index
isBan - true: ban; false: not ban
	`,
	Action: func(cctx *cli.Context) error {
		riStr := cctx.Args().Get(0)
		ri, err := strconv.ParseUint(riStr, 10, 0)
		if err != nil {
			log.Fatal("parse ri to uint64 err:", err, ". Please input correct ri")
		}
		fmt.Println("ri:", ri)

		isBan := cctx.Args().Get(1)
		ban := uint8(1)
		banBool := true
		switch isBan {
		case "false":
			ban = 0
			banBool = false
		case "true":
		default:
			log.Fatal("isBan err, please input true or false")
		}

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())

		sksSli := cctx.StringSlice("sks")
		fmt.Println("sks length:", len(sksSli))
		fmt.Println("5 sks:", sksSli)
		sks := [5]string{}
		copy(sks[:], sksSli)

		txopts := &com.TxOpts{
			Nonce:    nil,
			GasPrice: nil,
			GasLimit: 0,
		}

		txAuth, err := com.MakeAuth(cctx.Context, endPoint, sk, txopts)
		if err != nil {
			log.Fatal(err)
		}

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new proxyIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		proxyAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeProxy)
		if err != nil {
			log.Fatal(err)
		}
		proxyIns, _ := proxy.NewProxy(proxyAddr, client)

		// get control1Addr
		control1Addr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeControl1)
		if err != nil {
			log.Fatal(err)
		}

		// new authIns
		authAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeAuth)
		if err != nil {
			log.Fatal(err)
		}
		authIns, _ := auth.NewAuth(authAddr, client)

		// get nonce
		nonce, err := authIns.Nonce(&bind.CallOpts{From: addr})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("ban account nonce:", nonce)
		hash := com.GetBanHash(control1Addr, authAddr, ri, ban, nonce)
		signs := com.GetSigns(hash, sks)

		tx, err := proxyIns.Ban(txAuth, ri, banBool, signs)
		if err != nil {
			log.Fatal("ban ri err:", err)
		}
		fmt.Println("tx hash:", tx.Hash())

		fmt.Println("waiting tx complete...")

		time.Sleep(WaitTime)
		receipt := com.GetTransactionReceipt(endPoint, tx.Hash())
		fmt.Println("tx status:", receipt.Status)

		return nil
	},
}

var banGCmd = &cli.Command{
	Name:      "banG",
	Usage:     "Admin ban or not ban group.",
	ArgsUsage: "<group's index>",
	Description: `
A function of Proxy contract.
Admin ban group.

Arguments:
gi - the group's index
isBan - true: ban; false: not ban
	`,
	Action: func(cctx *cli.Context) error {
		giStr := cctx.Args().Get(0)
		gi, err := strconv.ParseUint(giStr, 10, 0)
		if err != nil {
			log.Fatal("parse gi to uint64 err:", err, ". Please input correct gi")
		}
		fmt.Println("gi:", gi)

		isBan := cctx.Args().Get(1)
		ban := uint8(1)
		banBool := true
		switch isBan {
		case "false":
			ban = 0
			banBool = false
		case "true":
		default:
			log.Fatal("isBan err, please input true or false")
		}

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())

		sksSli := cctx.StringSlice("sks")
		fmt.Println("sks length:", len(sksSli))
		fmt.Println("5 sks:", sksSli)
		sks := [5]string{}
		copy(sks[:], sksSli)

		txopts := &com.TxOpts{
			Nonce:    nil,
			GasPrice: nil,
			GasLimit: 0,
		}

		txAuth, err := com.MakeAuth(cctx.Context, endPoint, sk, txopts)
		if err != nil {
			log.Fatal(err)
		}

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new proxyIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		proxyAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeProxy)
		if err != nil {
			log.Fatal(err)
		}
		proxyIns, _ := proxy.NewProxy(proxyAddr, client)

		// get control1Addr
		control1Addr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeControl1)
		if err != nil {
			log.Fatal(err)
		}

		// new authIns
		authAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeAuth)
		if err != nil {
			log.Fatal(err)
		}
		authIns, _ := auth.NewAuth(authAddr, client)

		// get nonce
		nonce, err := authIns.Nonce(&bind.CallOpts{From: addr})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("ban group nonce:", nonce)
		hash := com.GetBanGHash(control1Addr, authAddr, gi, ban, nonce)
		signs := com.GetSigns(hash, sks)

		tx, err := proxyIns.BanG(txAuth, gi, banBool, signs)
		if err != nil {
			log.Fatal("ban gi err:", err)
		}
		fmt.Println("tx hash:", tx.Hash())

		fmt.Println("waiting tx complete...")

		time.Sleep(WaitTime)
		receipt := com.GetTransactionReceipt(endPoint, tx.Hash())
		fmt.Println("tx status:", receipt.Status)

		return nil
	},
}

var banTCmd = &cli.Command{
	Name:      "banT",
	Usage:     "Admin ban or not ban erc20-token.",
	ArgsUsage: "<erc20-token address>",
	Description: `
A function of Proxy contract.
Admin ban erc20-token.

Arguments:
ercAddr - the erc20-token address
isBan - true: ban; false: not ban
	`,
	Action: func(cctx *cli.Context) error {
		ercAddrStr := cctx.Args().Get(0)
		ercAddr := common.HexToAddress(ercAddrStr)
		fmt.Println("ercAddr:", ercAddr.Hex())

		isBan := cctx.Args().Get(1)
		ban := uint8(1)
		banBool := true
		switch isBan {
		case "false":
			ban = 0
			banBool = false
		case "true":
		default:
			log.Fatal("isBan err, please input true or false")
		}

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())

		sksSli := cctx.StringSlice("sks")
		fmt.Println("sks length:", len(sksSli))
		fmt.Println("5 sks:", sksSli)
		sks := [5]string{}
		copy(sks[:], sksSli)

		txopts := &com.TxOpts{
			Nonce:    nil,
			GasPrice: nil,
			GasLimit: 0,
		}

		txAuth, err := com.MakeAuth(cctx.Context, endPoint, sk, txopts)
		if err != nil {
			log.Fatal(err)
		}

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new proxyIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		proxyAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeProxy)
		if err != nil {
			log.Fatal(err)
		}
		proxyIns, _ := proxy.NewProxy(proxyAddr, client)

		// get control1Addr
		control1Addr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeControl1)
		if err != nil {
			log.Fatal(err)
		}

		// new authIns
		authAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeAuth)
		if err != nil {
			log.Fatal(err)
		}
		authIns, _ := auth.NewAuth(authAddr, client)

		// get nonce
		nonce, err := authIns.Nonce(&bind.CallOpts{From: addr})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("ban erc20-token nonce:", nonce)
		hash := com.GetBanTHash(control1Addr, authAddr, ercAddr, ban, nonce)
		signs := com.GetSigns(hash, sks)

		tx, err := proxyIns.BanT(txAuth, ercAddr, banBool, signs)
		if err != nil {
			log.Fatal("ban erc20-token err:", err)
		}
		fmt.Println("tx hash:", tx.Hash())

		fmt.Println("waiting tx complete...")

		time.Sleep(WaitTime)
		receipt := com.GetTransactionReceipt(endPoint, tx.Hash())
		fmt.Println("tx status:", receipt.Status)

		return nil
	},
}

var addTCmd = &cli.Command{
	Name:      "addT",
	Usage:     "Admin add erc20-token.",
	ArgsUsage: "<erc20-token address>",
	Description: `
A function of Proxy contract.
Admin add erc20-token.

Arguments:
ercAddr - the erc20-token address
isMint - true: will add reward when addOrder; false: will not add reward when addOrder
	`,
	Action: func(cctx *cli.Context) error {
		ercAddrStr := cctx.Args().Get(0)
		ercAddr := common.HexToAddress(ercAddrStr)
		fmt.Println("ercAddr:", ercAddr.Hex())

		mintStr := cctx.Args().Get(1)
		mint := uint8(1)
		isMint := true
		switch mintStr {
		case "false":
			mint = 0
			isMint = false
		case "true":
		default:
			log.Fatal("mintStr err, please input true or false")
		}

		addr := common.HexToAddress(cctx.String("adminAddr"))
		fmt.Println("adminAddr:", addr.Hex())
		sk := cctx.String("adminSk")
		fmt.Println("adminSk:", sk)
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())

		sksSli := cctx.StringSlice("sks")
		fmt.Println("sks length:", len(sksSli))
		fmt.Println("5 sks:", sksSli)
		sks := [5]string{}
		copy(sks[:], sksSli)

		txopts := &com.TxOpts{
			Nonce:    nil,
			GasPrice: nil,
			GasLimit: 0,
		}

		txAuth, err := com.MakeAuth(cctx.Context, endPoint, sk, txopts)
		if err != nil {
			log.Fatal(err)
		}

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new proxyIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		proxyAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeProxy)
		if err != nil {
			log.Fatal(err)
		}
		proxyIns, _ := proxy.NewProxy(proxyAddr, client)

		// get control1Addr
		control1Addr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeControl1)
		if err != nil {
			log.Fatal(err)
		}

		// new authIns
		authAddr, err := instanceIns.Instances(&bind.CallOpts{From: addr}, com.TypeAuth)
		if err != nil {
			log.Fatal(err)
		}
		authIns, _ := auth.NewAuth(authAddr, client)

		// get nonce
		nonce, err := authIns.Nonce(&bind.CallOpts{From: addr})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("add erc20-token nonce:", nonce)
		hash := com.GetAddTHash(control1Addr, authAddr, ercAddr, mint, nonce)
		signs := com.GetSigns(hash, sks)

		tx, err := proxyIns.AddT(txAuth, ercAddr, isMint, signs)
		if err != nil {
			log.Fatal("add erc20-token err:", err)
		}
		fmt.Println("tx hash:", tx.Hash())

		fmt.Println("waiting tx complete...")

		time.Sleep(WaitTime)
		receipt := com.GetTransactionReceipt(endPoint, tx.Hash())
		fmt.Println("tx status:", receipt.Status)

		return nil
	},
}
