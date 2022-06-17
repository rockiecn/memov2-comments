package inter

import "github.com/ethereum/go-ethereum/common"

type IAccess interface {
	Set(acc common.Address, isSet bool, signs [5][]byte) (common.Hash, error)
	Access(acc common.Address) (bool, error)
}