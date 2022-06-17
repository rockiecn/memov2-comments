package inter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IERC20 interface {
	Transfer(recipient common.Address, amount *big.Int) (common.Hash, error)
	Approve(spender common.Address, amount *big.Int) (common.Hash, error)
	TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (common.Hash, error)
	Mint(target common.Address, amount *big.Int) (common.Hash, error)
	Burn(amount *big.Int) (common.Hash, error)

	Name() (string, error)
	Symbol() (string, error)
	TotalSupply() (*big.Int, error)
	BalanceOf(account common.Address) (*big.Int, error)
	Allowance(owner common.Address, spender common.Address) (*big.Int, error)

	Version() (uint16, error)
	Access() (common.Address, error)
	InitialSupply() (*big.Int, error)
	MaxSupply() (*big.Int, error)
}
