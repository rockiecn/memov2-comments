package com

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	alterStr    = "alter"
	permStr     = "perm"
	addTStr     = "addT"
	banGStr     = "banG"
	banStr      = "ban"
	banTStr     = "banT"
	activateStr = "activate"
	setStr      = "set"
	addOwnerStr = "add"
)

func GetAlterHash(instanAddr, authAddr, addr common.Address, t uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(instanAddr.Bytes(), []byte(alterStr), addr.Bytes(), []byte{t})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetAddTHash(ctl1Addr, authAddr, ercAddr common.Address, ismint uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(addTStr), ercAddr.Bytes(), []byte{ismint})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetBanGHash(ctl1Addr, authAddr common.Address, gi uint64, isBan uint8, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, gi)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(banGStr), tmp8, []byte{isBan})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func GetActivateHash(ctl1Addr, authAddr common.Address, ki uint64, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, ki)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(activateStr), tmp8)
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// abi.encodePacked(address(this), nonce, "set", account, isSet)
func GetSetHash(access, acc common.Address, nonce *big.Int, isSet bool) []byte {
	set := uint8(0)
	if isSet {
		set = 1
	}
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash := crypto.Keccak256(access.Bytes(), m, []byte(setStr), acc.Bytes(), []byte{set})
	return hash
}

// abi.encodePacked(uIndex,pIndex,nonce,start,end,size,tIndex,sprice)
func GetAddOrderHash(uIndex, pIndex, start, end, size, nonce uint64, tIndex uint8, sprice *big.Int) []byte {
	tmp1 := make([]byte, 8)
	tmp2 := make([]byte, 8)
	tmp3 := make([]byte, 8)
	tmp4 := make([]byte, 8)
	tmp5 := make([]byte, 8)
	tmp6 := make([]byte, 8)

	binary.BigEndian.PutUint64(tmp1, uIndex)
	binary.BigEndian.PutUint64(tmp2, pIndex)
	binary.BigEndian.PutUint64(tmp3, nonce)
	binary.BigEndian.PutUint64(tmp4, start)
	binary.BigEndian.PutUint64(tmp5, end)
	binary.BigEndian.PutUint64(tmp6, size)

	m := common.LeftPadBytes(sprice.Bytes(), 32)

	hash := crypto.Keccak256(tmp1, tmp2, tmp3, tmp4, tmp5, tmp6, []byte{tIndex}, m)

	return hash
}

// abi.encodePacked(address(this), "add", _a, isSet)
// abi.encodePacked(address(this), nonce, "perm", ha)
func GetAddOwnerHash(ownerContractAddr, addr, authAddr common.Address, isSet bool, nonce *big.Int) []byte {
	set := uint8(0)
	if isSet {
		set = 1
	}
	hash := crypto.Keccak256(ownerContractAddr.Bytes(), []byte(addOwnerStr), addr.Bytes(), []byte{set})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// abi.encodePacked(ps.pIndex,ps.tIndex,ps.pay,ps.lost)
func GetProWithdrawHash(pindex uint64, tindex uint8, pay, lost *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, pindex)

	m := common.LeftPadBytes(pay.Bytes(), 32)
	n := common.LeftPadBytes(lost.Bytes(), 32)

	hash := crypto.Keccak256(tmp8, []byte{tindex}, m, n)
	return hash
}

func GetBanHash(ctl1Addr, authAddr common.Address, ri uint64, isBan uint8, nonce *big.Int) []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, ri)
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(banStr), tmp8, []byte{isBan})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

// GetBanTHash tokenAddr: erc20-token address
func GetBanTHash(ctl1Addr, authAddr, tokenAddr common.Address, isBan uint8, nonce *big.Int) []byte {
	hash := crypto.Keccak256(ctl1Addr.Bytes(), []byte(banTStr), tokenAddr.Bytes(), []byte{isBan})
	m := common.LeftPadBytes(nonce.Bytes(), 32)
	hash = crypto.Keccak256(authAddr.Bytes(), m, []byte(permStr), hash)
	return hash
}

func Sign(hash []byte, sk string) ([]byte, error) {
	skEcdsa, err := crypto.HexToECDSA(sk)
	if err != nil {
		return nil, err
	}

	sig, err := crypto.Sign(hash, skEcdsa)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func GetSigns(hash []byte, sks [5]string) [5][]byte {
	res := [5][]byte{}

	for i := 0; i < 5; i++ {
		sign, err := Sign(hash, sks[i])
		if err != nil {
			fmt.Println("get signs i: ", i, " err:", err)
			res[i] = nil
		} else {
			res[i] = sign
		}
	}

	return res
}
