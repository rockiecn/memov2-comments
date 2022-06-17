package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/erc"
)

var (
	keeperAddr = common.HexToAddress("0x70997970c51812dc3a010c7d01b50e0d17dc79c8")
	waitTime   = 30 * time.Second
)

// test check reRole in addToGroup
func main() {
	inputeth := flag.String("eth", "https://devchain.metamemo.one:8501", "eth api Address;") //dev net
	sk := flag.String("sk", "", "admin's sk")

	flag.Parse()

	eth := *inputeth
	hexSk := *sk

	if hexSk == "" {
		fmt.Println("Please input your sk!")
		return
	}

	fmt.Println()

	// make auth to send transaction
	txopts := &com.TxOpts{Nonce: nil, GasPrice: nil, GasLimit: 0, Money: nil} // gasLimit set 0, run estimateGasLimit to precheck tx
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

	// deploy erc20
	fmt.Println("deploy erc20..")
	erc20Addr, tx, _, err := erc.DeployERC20(txAuth, client, com.AccessAddr, "memo", "M")
	if err != nil {
		log.Fatal("deploy erc20 err:", err)
	}
	fmt.Println("erc20Addr:", erc20Addr.Hex())

	time.Sleep(waitTime)

	// get receipt
	receipt := com.GetTransactionReceipt(eth, tx.Hash())
	fmt.Println("receipt status:", receipt.Status)
	fmt.Println()
}
