package inter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/memoio/contractsv2/go_contracts/proxy"
)

type IProxy interface {
	Activate(i uint64, signs [5][]byte) (common.Hash, error)
	Ban(i uint64, ban bool, signs [5][]byte) (common.Hash, error)
	AddT(t common.Address, mint bool, signs [5][]byte) (common.Hash, error)
	BanT(t common.Address, ban bool, signs [5][]byte) (common.Hash, error)
	BanG(gi uint64, ban bool, signs [5][]byte) (common.Hash, error)

	CreateGroup(level uint8, mr uint8, dc uint8, pc uint8, kr *big.Int, pr *big.Int) (common.Hash, error)

	// register self to get index
	ReAcc() (common.Hash, error)
	ReRole(rtype uint8, extra []byte) (common.Hash, error)

	// add a user/keeper/provider to group
	AddToGroup(gi uint64) (common.Hash, error)

	Pledge(i uint64, money *big.Int) (common.Hash, error)
	Unpledge(i uint64, ti uint8, money *big.Int) (common.Hash, error)

	AlterPayee(p common.Address) (common.Hash, error)
	ConfirmPayee(i uint64) (common.Hash, error)
	AddOrder(oi proxy.OrderIn, uSign []byte, pSign []byte) (common.Hash, error)
	SubOrder(oi proxy.OrderIn, uSign []byte, pSign []byte) (common.Hash, error)

	AddRepair(oi proxy.OrderIn, kis []uint64, ksigns []byte) (common.Hash, error)

	Recharge(i uint64, ti uint8, isLock bool, money *big.Int) (common.Hash, error)
	Withdraw(i uint64, ti uint8, money *big.Int) (common.Hash, error)
	ProWithdraw(ps proxy.PWIn, kis []uint64, ksigns []byte) (common.Hash, error)
}
