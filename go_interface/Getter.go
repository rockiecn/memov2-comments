package inter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/memoio/contractsv2/go_contracts/getter"
)

type IGetter interface {
	Version() (v uint16, err error)
	Inst() (inst common.Address, err error)

	GetTA(index uint8) (ta common.Address, err error)
	GetTI(t common.Address) (ti uint8, isban bool, mint bool, err error)
	GetTCnt() (tcnt uint8, err error)

	FsBalanceOf(index uint64, ti uint8) (avail *big.Int, lock *big.Int)
	GetGSInfo(gi uint64, ti uint8) (ginfo getter.GroupOut, err error)
	GetFsInfo(ui uint64, pi uint64) (fsinfo getter.FsOut, err error)
	GetStoreInfo(ui uint64, pi uint64, ti uint8) (storeinfo getter.StoreOut, err error)
	SettleInfo(pi uint64, ti uint8) (settleinfo getter.SettleOut, err error)

	GetGCnt() (gcnt uint64, err error)
	GetGInfo(gi uint64) (isActive bool, isBan bool, err error)
	GetLevel(gi uint64) (level uint8, err error)
	GetPInfo(gi uint64) (kpr, ppr *big.Int, err error)
	GetKManage(gi uint64) (kmanage common.Address, err error)
	GetSStra(gi uint64) (dc, pc uint8, err error)

	ManageRate(gi uint64) (mr uint8, err error)
	GetKCnt(gi uint64) (kcnt uint8, err error)
	GetK(index, gi uint64) (ki uint64, err error)
	GetPf(ti uint8, gi uint64) (last uint64, accu *big.Int, err error)
	KMBalanceOf(ki, gi uint64, ti uint8) (avail, tmp *big.Int, err error)

	Instances(_type uint8) (instance common.Address, err error)

	PleBalanceOf(index uint64, ti uint8) (val *big.Int, err error)
	GetPledge(ti uint8) (pledge *big.Int, err error)
	GetToTalPledge() (allPledge *big.Int, err error)

	GetAcc(a common.Address) (ri uint64, err error)
	GetAddr(ri uint64) (a common.Address, err error)
	GetACnt() (acnt uint64, err error)
	GetRInfo(a common.Address) (roleinfo getter.RoleOut, err error)
}
