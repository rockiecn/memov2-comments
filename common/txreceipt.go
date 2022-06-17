package com

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TxOpts include basic parameters when sending a transaction
type TxOpts struct {
	Nonce    *big.Int // default nil
	GasPrice *big.Int // default nil, so use the recommended value for ethereum
	GasLimit uint64   // default recommended
	Money    *big.Int
}

func MakeAuth(ctx context.Context, endPoint string, hexSk string, txopts *TxOpts) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		fmt.Println("HexToECDSA err: ", err)
		return auth, err
	}

	client, err := ethclient.DialContext(ctx, endPoint)
	if err != nil {
		return auth, err
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("client.NetworkID error, use the default chainID")
		chainID = big.NewInt(666)
	}

	auth, err = bind.NewKeyedTransactorWithChainID(sk, chainID)
	if err != nil {
		return nil, err
	}
	auth.GasPrice = txopts.GasPrice
	auth.Value = txopts.Money // money to contract-account
	auth.Nonce = txopts.Nonce
	auth.GasLimit = txopts.GasLimit
	return auth, nil
}

//GetTransactionReceipt 通过交易hash获得交易详情
func GetTransactionReceipt(endPoint string, hash common.Hash) *types.Receipt {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		log.Fatal("rpc.Dial err", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Println("get transaction receipt: ", err)
	}
	return receipt
}
