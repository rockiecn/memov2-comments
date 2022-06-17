package com

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/xerrors"
)

var (
	DevChain     = "https://devchain.metamemo.one:8501"
	TestChain    = "http://119.147.213.220:8191"
	ProductChain = "https://chain.metamemo.one:8501"

	DefaultGasLimit = uint64(8000000) // as small as possible
	PledgeAmount    = "1 memo"        // 1 memo = 1e9 nanomemo = 1e18 attomemo
	nano            = big.NewInt(1e9)
	memo            = big.NewInt(1e18)

	// AdminAddr admin address
	AdminAddr = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
	// Foundation foundation address
	Foundation = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")

	// five admin addrs for multi-signature
	Addrs = [5]common.Address{common.HexToAddress("0xB651b3ce6C7c712E41b03993a11DA656eF3C95A8"), common.HexToAddress("0x54a18833beaa8A22672ffF4e5b6714eCebC9368b"), common.HexToAddress("0x6f7bc6032278753c8A13bd430b13f1965407633f"), common.HexToAddress("0x09d186ce74f9092f5e6F20390d730943170bDbBA"), common.HexToAddress("0x8AadBf87Dc8Ae4A1161DEF8185D2b98C1904a3c3")}

	Erc20Name   = "memo"
	Erc20Symbol = "M"

	TypeProxy    = uint8(100)
	TypeControl1 = uint8(101)
	TypeControl2 = uint8(102)
	TypeAuth     = uint8(2)
	TypeAccess   = uint8(3)
	TypeERC20    = uint8(4)
	TypePoolP    = uint8(5)
	TypeRole     = uint8(6)
	TypeToken    = uint8(7)
	TypePledge   = uint8(8)
	TypeIssu     = uint8(9)
	Typefs       = uint8(10)
	TypeGroup    = uint8(11)
	TypePoolF    = uint8(12)
	TypeGetter   = uint8(150)

	UserType     = uint8(1)
	ProviderType = uint8(2)
	KeeperType   = uint8(3)
)

var (
	// dev
	ProxyAddr    = common.HexToAddress("0xef40939A4F1F4b70250B1A240c8336A038878973")
	AccessAddr   = common.HexToAddress("0x3438C5f3012CE65c29F06Ca1dc87e0f24ef9c5Da")
	ERC20Addr    = common.HexToAddress("0xA13d8dac10F8102F67CEa848Ec20E2972C96C4Fc")
	AuthAddr     = common.HexToAddress("0x88dc511d82B125dC7A4b762c77C69A4BA72D717c")
	InstanceAddr = common.HexToAddress("0xC00166859f9443A0Df655266915166180FA57690")
	Control1Addr = common.HexToAddress("0x156DDc3b4874dF2275c8f44D6DDF7211B10f35FC")
	Control2Addr = common.HexToAddress("0x5734f31df97f2b3Cc81D90a068EB02Ed918f287e")
	TokenAddr    = common.HexToAddress("0x477Ad9d34F69DEeef5660cDcfB8FE1a0227D55D1")
	FilesysAddr  = common.HexToAddress("0x5cf7D409cCA50C99f67bcd475caA9ce220BC7bDb")
	GroupAddr    = common.HexToAddress("0x4801b71717Ec638db33F2CE582bf8Dc58ace8202")
	IssuanceAddr = common.HexToAddress("0x3a5A68ad5F585d10072cE49f056aa2010A4c5c40")
	PoolPAddr    = common.HexToAddress("0x539787Bd188d6ce78F8A7E7EC6e6d500C10Db64A")
	PledgeAddr   = common.HexToAddress("0xA0d29C31749925c5ACA48e2b88E4D331E94458C8")
	PoolFAddr    = common.HexToAddress("0x4324582F67097546577d1D05f014385A2BAb884b")
	RoleAddr     = common.HexToAddress("0x153F4EF777c88401610cF2F2f5a9feF7f36E0ec0")
	GetterAddr   = common.HexToAddress("0x5b3707852430b19dCa0f671B821fa0c446695F30")

	// testnet
	TestnetInstanceAddr = common.HexToAddress("")

	// product
	ProductInstanceAddr = common.HexToAddress("")
)

var (
	ErrMemoAmount = xerrors.New("memoAmount's uint should be memo、nanomemo、attomemo; number and uint should be separated with a space; such as:1 memo")
)

// TODO: deploy contract on other chain
func GetInsEndPointByChain(chain string) (common.Address, string) {
	switch chain {
	case "dev":
		return InstanceAddr, DevChain
	case "test":
		return TestnetInstanceAddr, TestChain
	case "product":
		return ProductInstanceAddr, ProductChain
	}
	return InstanceAddr, DevChain
}

// ParseMemo parse an amount string with units
func ParseMemo(memoAmount string) (*big.Int, error) {
	strSli := strings.Split(memoAmount, " ")
	if len(strSli) != 2 {
		return nil, ErrMemoAmount
	}

	num, err := strconv.ParseInt(strSli[0], 10, 0)
	if err != nil {
		return nil, ErrMemoAmount
	}

	value := big.NewInt(num)

	switch strSli[1] {
	case "memo":
		value.Mul(value, memo)
	case "nanomemo":
		value.Mul(value, nano)
	case "attomemo":
	default:
		return nil, ErrMemoAmount
	}
	return value, nil
}

func FormatMemo(value *big.Int) string {
	var m, n, a int64
	var mStr, nStr, aStr string

	tmp := big.NewInt(0)

	if value.Cmp(memo) >= 0 {
		m = tmp.Div(value, memo).Int64()
		mStr = strconv.Itoa(int(m)) + " memo "
		value.Mod(value, memo)
	}
	if value.Cmp(nano) >= 0 {
		n = tmp.Div(value, nano).Int64()
		nStr = strconv.Itoa(int(n)) + " nanomemo "
		value.Mod(value, nano)
	}
	a = value.Int64()
	if a != 0 {
		aStr = strconv.Itoa(int(a)) + " attomemo"
	}

	res := mStr + nStr + aStr
	if res == "" {
		return "0 attomemo"
	}
	return res
}
