// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package group

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GroupABI is the input ABI used to generate the binding from.
const GroupABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inst\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dc\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"pc\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pr\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_kc\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_pm\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"checkG\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_dc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_pc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getKManage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getSStra\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GroupFuncSigs maps the 4-byte function signature to its string representation.
var GroupFuncSigs = map[string]string{
	"ad578aeb": "activate(uint64,uint8)",
	"4bf1b134": "add(address,bool,bytes[5])",
	"4acc2468": "add(uint8,uint64,uint256)",
	"de9375f2": "auth()",
	"0f78c61a": "ban(uint64,bool)",
	"83889fe3": "checkG(uint64)",
	"eaf724ab": "create(uint8,uint8,uint8,uint8,uint256,uint256)",
	"059e783d": "getGCnt()",
	"38add6a1": "getGInfo(uint64)",
	"929629d2": "getKManage(uint64)",
	"770609a8": "getLevel(uint64)",
	"50aa2329": "getPInfo(uint64)",
	"dc698953": "getSStra(uint64)",
	"bd6b2222": "inst()",
	"022914a7": "owners(address)",
	"54fd4d50": "version()",
}

// GroupBin is the compiled bytecode used for deploying new contracts.
var GroupBin = "0x60806040526001805461ffff60a01b1916600160a11b17905534801561002457600080fd5b506040516126a33803806126a3833981016040819052610043916101e9565b600180546001600160a01b039384166001600160a01b03199182161782556002805493851693909116929092179091556040805161010080820183526000808352928201838152606083018481526080840185815260a0850186815260c0860187815260e08701888152602088018a8152600380549b8c018155998a905297517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9a909902998a0180549851965195519451935161ffff1990991699151561ff001916999099179515159096029490941763ffff000019166201000060ff9485160263ff00000019161763010000009284169290920291909117600160201b600160c81b0319166401000000009290911691909102600160281b600160c81b0319161765010000000000939096169290920294909417909155517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c82015590517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d9091015561021c565b80516001600160a01b03811681146101e457600080fd5b919050565b600080604083850312156101fc57600080fd5b610205836101cd565b9150610213602084016101cd565b90509250929050565b6124788061022b6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063770609a811610097578063bd6b222211610066578063bd6b222214610283578063dc69895314610296578063de9375f2146102c3578063eaf724ab146102d657600080fd5b8063770609a81461020d57806383889fe314610232578063929629d214610245578063ad578aeb1461027057600080fd5b80634acc2468116100d35780634acc2468146101975780634bf1b134146101aa57806350aa2329146101bd57806354fd4d50146101e557600080fd5b8063022914a714610105578063059e783d1461013d5780630f78c61a1461015857806338add6a11461016d575b600080fd5b610128610113366004610e04565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6003546040516001600160401b039091168152602001610134565b61016b610166366004610e54565b6102e9565b005b61018061017b366004610e87565b610368565b604080519215158352901515602083015201610134565b61016b6101a5366004610eb3565b6103e2565b61016b6101b8366004610f5d565b61063f565b6101d06101cb366004610e87565b6107b8565b60408051928352602083019190915201610134565b6001546101fa90600160a01b900461ffff1681565b60405161ffff9091168152602001610134565b61022061021b366004610e87565b610820565b60405160ff9091168152602001610134565b61016b610240366004610e87565b61085e565b610258610253366004610e87565b6108d8565b6040516001600160a01b039091168152602001610134565b61016b61027e36600461106d565b61091f565b600254610258906001600160a01b031681565b6102a96102a4366004610e87565b610a41565b6040805160ff938416815292909116602083015201610134565b600154610258906001600160a01b031681565b61016b6102e4366004611097565b610ac3565b3360009081526020819052604090205460ff166103215760405162461bcd60e51b8152600401610318906110fd565b60405180910390fd5b806003836001600160401b03168154811061033e5761033e611120565b6000918252602090912060039091020180549115156101000261ff00199092169190911790555050565b6000806003836001600160401b03168154811061038757610387611120565b6000918252602090912060039182020154815460ff90911691906001600160401b0386169081106103ba576103ba611120565b906000526020600020906003020160000160019054906101000a900460ff1691509150915091565b6003826001600160401b0316815481106103fe576103fe611120565b6000918252602090912060039091020154610100900460ff16156104535760405162461bcd60e51b815260206004820152600c60248201526b1859190e99c818985b9b995960a21b6044820152606401610318565b8260ff166001036104d3576003826001600160401b03168154811061047a5761047a611120565b600091825260209091206003909102015460ff166104ce5760405162461bcd60e51b8152602060048201526011602482015270616464753a67206e6f742061637469766560781b6044820152606401610318565b505050565b8260ff166002036105c1576003826001600160401b0316815481106104fa576104fa611120565b600091825260209091206003909102015460ff1661054e5760405162461bcd60e51b8152602060048201526011602482015270616464503a67206e6f742061637469766560781b6044820152606401610318565b6003826001600160401b03168154811061056a5761056a611120565b9060005260206000209060030201600201548110156104ce5760405162461bcd60e51b81526020600482015260136024820152720e040e0d8cac8ceca40dcdee840cadcdeeaced606b1b6044820152606401610318565b8260ff166003036104ce576003826001600160401b0316815481106105e8576105e8611120565b9060005260206000209060030201600101548110156104ce5760405162461bcd60e51b81526020600482015260136024820152720d640e0d8cac8ceca40dcdee840cadcdeeaced606b1b6044820152606401610318565b823b60008190036106875760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b6044820152606401610318565b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906107149084908790600401611136565b600060405180830381600087803b15801561072e57600080fd5b505af1158015610742573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b6000806003836001600160401b0316815481106107d7576107d7611120565b9060005260206000209060030201600101546003846001600160401b03168154811061080557610805611120565b90600052602060002090600302016002015491509150915091565b60006003826001600160401b03168154811061083e5761083e611120565b600091825260209091206003909102015462010000900460ff1692915050565b6003816001600160401b03168154811061087a5761087a611120565b600091825260209091206003909102015460ff1680156108cc57506003816001600160401b0316815481106108b1576108b1611120565b6000918252602090912060039091020154610100900460ff16155b6108d557600080fd5b50565b60006003826001600160401b0316815481106108f6576108f6611120565b60009182526020909120600390910201546501000000000090046001600160a01b031692915050565b3360009081526020819052604090205460ff1661094e5760405162461bcd60e51b8152600401610318906110fd565b6003826001600160401b03168154811061096a5761096a611120565b6000918252602090912060039091020154610100900460ff16156109bf5760405162461bcd60e51b815260206004820152600c60248201526b1858dd0e99c818985b9b995960a21b6044820152606401610318565b6003826001600160401b0316815481106109db576109db611120565b600091825260209091206003909102015460ff62010000909104811690821610610a3d5760016003836001600160401b031681548110610a1d57610a1d611120565b60009182526020909120600390910201805460ff19169115159190911790555b5050565b6000806003836001600160401b031681548110610a6057610a60611120565b906000526020600020906003020160000160039054906101000a900460ff166003846001600160401b031681548110610a9b57610a9b611120565b906000526020600020906003020160000160049054906101000a900460ff1691509150915091565b3360009081526020819052604090205460ff16610af25760405162461bcd60e51b8152600401610318906110fd565b60035460028054604051633ec7d5b960e01b815260048101929092526001600160a01b0316906000908290633ec7d5b990602401602060405180830381865afa158015610b43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6791906111c6565b88604051610b7490610de2565b6001600160a01b03909216825260ff166020820152604001604051809103906000f080158015610ba8573d6000803e3d6000fd5b506040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290915060ff808b16604080840191825260c0840189815260e085018981528c8516606087019081528c8616608088019081526001600160a01b03808a1660a08a01908152600160208b0181815260038054928301815560008190528c5192027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805492519b51975196519451909516650100000000000265010000000000600160c81b0319948d166401000000000294909416640100000000600160c81b0319968d1663010000000263ff0000001998909d1662010000029790971663ffff0000199b15156101000261ff00199415159490941661ffff199093169290921792909217999099169890981798909817919091169190911795909517909455517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c83015591517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d90910155517f2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d9990610dce9086908d908d908d908d908d908d906001600160401b0397909716875260ff958616602088015293851660408701529184166060860152909216608084015260a083019190915260c082015260e00190565b60405180910390a150505050505050505050565b61125f806111e483390190565b6001600160a01b03811681146108d557600080fd5b600060208284031215610e1657600080fd5b8135610e2181610def565b9392505050565b80356001600160401b0381168114610e3f57600080fd5b919050565b80358015158114610e3f57600080fd5b60008060408385031215610e6757600080fd5b610e7083610e28565b9150610e7e60208401610e44565b90509250929050565b600060208284031215610e9957600080fd5b610e2182610e28565b803560ff81168114610e3f57600080fd5b600080600060608486031215610ec857600080fd5b610ed184610ea2565b9250610edf60208501610e28565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715610f2757610f27610eef565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610f5557610f55610eef565b604052919050565b600080600060608486031215610f7257600080fd5b8335610f7d81610def565b92506020610f8c858201610e44565b925060408501356001600160401b0380821115610fa857600080fd5b8187019150601f8881840112610fbd57600080fd5b610fc5610f05565b8060a085018b811115610fd757600080fd5b855b8181101561105b57803586811115610ff15760008081fd5b87018581018e136110025760008081fd5b80358781111561101457611014610eef565b611025818801601f19168b01610f2d565b8181528f8b83850101111561103a5760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610fd9565b50508096505050505050509250925092565b6000806040838503121561108057600080fd5b61108983610e28565b9150610e7e60208401610ea2565b60008060008060008060c087890312156110b057600080fd5b6110b987610ea2565b95506110c760208801610ea2565b94506110d560408801610ea2565b93506110e360608801610ea2565b92506080870135915060a087013590509295509295509295565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052603260045260246000fd5b8281526040602080830182905260009160e08401919084018584805b60058110156111b857878603603f1901845282518051808852835b8181101561118857828101880151898201890152870161116d565b81811115611198578488838b0101525b50601f01601f191696909601850195509284019291840191600101611152565b509398975050505050505050565b6000602082840312156111d857600080fd5b8151610e2181610def56fe608060405260018054600160a01b600160f81b03191665049d4002000160a11b17905534801561002e57600080fd5b5060405161125f38038061125f83398101604081905261004d91610081565b6001805460ff909216600160b01b02600162ff000160a01b03199092166001600160a01b03909316929092171790556100cd565b6000806040838503121561009457600080fd5b82516001600160a01b03811681146100ab57600080fd5b602084015190925060ff811681146100c257600080fd5b809150509250929050565b611183806100dc6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806350cbb46f1161008c5780636d23f6c8116100665780636d23f6c81461020d578063dc0e7f4514610221578063de9375f21461024c578063fc3ba0ad1461027757600080fd5b806350cbb46f146101bf57806354fd4d50146101d257806355d3d7ef146101fa57600080fd5b8063022914a7146100d4578063024130e41461010c578063259c6d5e1461012157806339f2db96146101425780634bf1b134146101585780634f3c2eab1461016b575b600080fd5b6100f76100e2366004610c7f565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61011f61011a366004610cb1565b61029f565b005b61013461012f366004610cf5565b61041f565b604051908152602001610103565b6003545b60405160ff9091168152602001610103565b61011f610166366004610d9f565b61080b565b6101a0610179366004610eb4565b60ff16600090815260056020526040902080546001909101546001600160401b0390911691565b604080516001600160401b039093168352602083019190915201610103565b61011f6101cd366004610ecf565b610984565b6001546101e790600160a01b900461ffff1681565b60405161ffff9091168152602001610103565b61011f610208366004610eea565b610ab6565b60015461014690600160b01b900460ff1681565b61023461022f366004610ecf565b610b44565b6040516001600160401b039091168152602001610103565b60015461025f906001600160a01b031681565b6040516001600160a01b039091168152602001610103565b61028a610285366004610f14565b610b92565b60408051928352602083019190915201610103565b3360009081526020819052604090205460ff166102d75760405162461bcd60e51b81526004016102ce90610f3e565b60405180910390fd5b6001600160401b03808316600090815260046020526040902054166103335760405162461bcd60e51b815260206004820152601260248201527139a7b93232b91d34b636329031b0b63632b960711b60448201526064016102ce565b6001600160401b0380831660009081526004602052604081208054849391929161035f91859116610f77565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555080600260008282829054906101000a90046001600160401b03166103a79190610f77565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550816001600160401b03167f5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c48260405161041391906001600160401b0391909116815260200190565b60405180910390a25050565b3360009081526020819052604081205460ff1661044e5760405162461bcd60e51b81526004016102ce90610f3e565b6001600160401b03808516600090815260046020526040812054909116900361047957506000610804565b60015460ff841660009081526005602052604090205442916001600160401b03600160b81b9091048116916104af911683610fa2565b6001600160401b031611156107905760ff84166000908152600560205260408120805467ffffffffffffffff19166001600160401b03938416178155600254600190910154919283926105059290911690610fe0565b905060005b600354811015610771576000600460006003848154811061052d5761052d610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a900482168452830193909352604090910190205461057391168461100a565b90506002600460006003858154811061058e5761058e610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a90048216845283019390935260409091019020546105d4929116611029565b6105df906001610f77565b60046000600385815481106105f6576105f6610ff4565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b03166001600160401b03166001600160401b0316815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b03160217905550600460006003848154811061067a5761067a610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a90048216845283019390935260409091019020546106c0911685610f77565b93508060066000600385815481106106da576106da610ff4565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b03168352828101939093526040918201812060ff8c1682529092528120805490919061072d90849061104f565b909155505060ff871660009081526005602052604081206001018054839290610757908490611067565b9091555082915061076990508161107e565b91505061050a565b50506002805467ffffffffffffffff19166001600160401b0383161790555b6001600160401b038516600090815260066020908152604080832060ff88168452909152902054808411156107c3578093505b6001600160401b038616600090815260066020908152604080832060ff89168452909152812080548692906107f9908490611067565b909155508493505050505b9392505050565b823b60008190036108535760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064016102ce565b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906108e09084908790600401611097565b600060405180830381600087803b1580156108fa57600080fd5b505af115801561090e573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b3360009081526020819052604090205460ff166109b35760405162461bcd60e51b81526004016102ce90610f3e565b6001600160401b038082166000908152600460205260409020541615610a055760405162461bcd60e51b81526020600482015260076024820152661ac8195e1a5cdd60ca1b60448201526064016102ce565b600380546001818101835560048083047fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160401b03808816600896909716959095026101000a86810290860219909116179055600093845260205260408320805467ffffffffffffffff191690911790556002805490911691610a8e83611127565b91906101000a8154816001600160401b0302191690836001600160401b031602179055505050565b3360009081526020819052604090205460ff16610ae55760405162461bcd60e51b81526004016102ce90610f3e565b60ff821660009081526005602052604081206001018054839290610b0a90849061104f565b909155505060405181815260ff8316907fe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e890602001610413565b60006003826001600160401b031681548110610b6257610b62610ff4565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b03169050919050565b6001600160401b03808316600081815260066020908152604080832060ff871680855290835281842054948452600483528184205460025491855260059093529083206001015492948594938593849390831692610bf292911690610fe0565b610bfc919061100a565b60015460ff88166000908152600560205260409020549192506001600160401b03600160b81b909104811691610c33911642611067565b1115610c4a57610c43818461104f565b9250610c57565b610c54818361104f565b91505b50909590945092505050565b80356001600160a01b0381168114610c7a57600080fd5b919050565b600060208284031215610c9157600080fd5b61080482610c63565b80356001600160401b0381168114610c7a57600080fd5b60008060408385031215610cc457600080fd5b610ccd83610c9a565b9150610cdb60208401610c9a565b90509250929050565b803560ff81168114610c7a57600080fd5b600080600060608486031215610d0a57600080fd5b610d1384610c9a565b9250610d2160208501610ce4565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715610d6957610d69610d31565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610d9757610d97610d31565b604052919050565b600080600060608486031215610db457600080fd5b610dbd84610c63565b92506020808501358015158114610dd357600080fd5b925060408501356001600160401b0380821115610def57600080fd5b8187019150601f8881840112610e0457600080fd5b610e0c610d47565b8060a085018b811115610e1e57600080fd5b855b81811015610ea257803586811115610e385760008081fd5b87018581018e13610e495760008081fd5b803587811115610e5b57610e5b610d31565b610e6c818801601f19168b01610d6f565b8181528f8b838501011115610e815760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610e20565b50508096505050505050509250925092565b600060208284031215610ec657600080fd5b61080482610ce4565b600060208284031215610ee157600080fd5b61080482610c9a565b60008060408385031215610efd57600080fd5b610f0683610ce4565b946020939093013593505050565b60008060408385031215610f2757600080fd5b610f3083610c9a565b9150610cdb60208401610ce4565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808316818516808303821115610f9957610f99610f61565b01949350505050565b60006001600160401b0383811690831681811015610fc257610fc2610f61565b039392505050565b634e487b7160e01b600052601260045260246000fd5b600082610fef57610fef610fca565b500490565b634e487b7160e01b600052603260045260246000fd5b600081600019048311821515161561102457611024610f61565b500290565b60006001600160401b038084168061104357611043610fca565b92169190910492915050565b6000821982111561106257611062610f61565b500190565b60008282101561107957611079610f61565b500390565b60006001820161109057611090610f61565b5060010190565b8281526040602080830182905260009160e08401919084018584805b600581101561111957878603603f1901845282518051808852835b818110156110e95782810188015189820189015287016110ce565b818111156110f9578488838b0101525b50601f01601f1916969096018501955092840192918401916001016110b3565b509398975050505050505050565b60006001600160401b0380831681810361114357611143610f61565b600101939250505056fea2646970667358221220d61b4c43fbc2afc6384ec21a987e1f29503e04a5de72681e2e2579554deeec9b64736f6c634300080e0033a2646970667358221220ad8f513e70578280cf9572013fca372f4a55288ec029545069584faf2975c32864736f6c634300080e0033"

// DeployGroup deploys a new Ethereum contract, binding an instance of Group to it.
func DeployGroup(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, _inst common.Address) (common.Address, *types.Transaction, *Group, error) {
	parsed, err := abi.JSON(strings.NewReader(GroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GroupBin), backend, _a, _inst)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Group{GroupCaller: GroupCaller{contract: contract}, GroupTransactor: GroupTransactor{contract: contract}, GroupFilterer: GroupFilterer{contract: contract}}, nil
}

// Group is an auto generated Go binding around an Ethereum contract.
type Group struct {
	GroupCaller     // Read-only binding to the contract
	GroupTransactor // Write-only binding to the contract
	GroupFilterer   // Log filterer for contract events
}

// GroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type GroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GroupSession struct {
	Contract     *Group            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GroupCallerSession struct {
	Contract *GroupCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GroupTransactorSession struct {
	Contract     *GroupTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type GroupRaw struct {
	Contract *Group // Generic contract binding to access the raw methods on
}

// GroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GroupCallerRaw struct {
	Contract *GroupCaller // Generic read-only contract binding to access the raw methods on
}

// GroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GroupTransactorRaw struct {
	Contract *GroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGroup creates a new instance of Group, bound to a specific deployed contract.
func NewGroup(address common.Address, backend bind.ContractBackend) (*Group, error) {
	contract, err := bindGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Group{GroupCaller: GroupCaller{contract: contract}, GroupTransactor: GroupTransactor{contract: contract}, GroupFilterer: GroupFilterer{contract: contract}}, nil
}

// NewGroupCaller creates a new read-only instance of Group, bound to a specific deployed contract.
func NewGroupCaller(address common.Address, caller bind.ContractCaller) (*GroupCaller, error) {
	contract, err := bindGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GroupCaller{contract: contract}, nil
}

// NewGroupTransactor creates a new write-only instance of Group, bound to a specific deployed contract.
func NewGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*GroupTransactor, error) {
	contract, err := bindGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GroupTransactor{contract: contract}, nil
}

// NewGroupFilterer creates a new log filterer instance of Group, bound to a specific deployed contract.
func NewGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*GroupFilterer, error) {
	contract, err := bindGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GroupFilterer{contract: contract}, nil
}

// bindGroup binds a generic wrapper to an already deployed contract.
func bindGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Group *GroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Group.Contract.GroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Group *GroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Group.Contract.GroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Group *GroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Group.Contract.GroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Group *GroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Group.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Group *GroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Group.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Group *GroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Group.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_Group *GroupCaller) Add(opts *bind.CallOpts, _rType uint8, _gi uint64, _pm *big.Int) error {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "add", _rType, _gi, _pm)

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_Group *GroupSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _Group.Contract.Add(&_Group.CallOpts, _rType, _gi, _pm)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_Group *GroupCallerSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _Group.Contract.Add(&_Group.CallOpts, _rType, _gi, _pm)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Group *GroupCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Group *GroupSession) Auth() (common.Address, error) {
	return _Group.Contract.Auth(&_Group.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Group *GroupCallerSession) Auth() (common.Address, error) {
	return _Group.Contract.Auth(&_Group.CallOpts)
}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_Group *GroupCaller) CheckG(opts *bind.CallOpts, i uint64) error {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "checkG", i)

	if err != nil {
		return err
	}

	return err

}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_Group *GroupSession) CheckG(i uint64) error {
	return _Group.Contract.CheckG(&_Group.CallOpts, i)
}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_Group *GroupCallerSession) CheckG(i uint64) error {
	return _Group.Contract.CheckG(&_Group.CallOpts, i)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_Group *GroupCaller) GetGCnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getGCnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_Group *GroupSession) GetGCnt() (uint64, error) {
	return _Group.Contract.GetGCnt(&_Group.CallOpts)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_Group *GroupCallerSession) GetGCnt() (uint64, error) {
	return _Group.Contract.GetGCnt(&_Group.CallOpts)
}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_Group *GroupCaller) GetGInfo(opts *bind.CallOpts, i uint64) (bool, bool, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getGInfo", i)

	if err != nil {
		return *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_Group *GroupSession) GetGInfo(i uint64) (bool, bool, error) {
	return _Group.Contract.GetGInfo(&_Group.CallOpts, i)
}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_Group *GroupCallerSession) GetGInfo(i uint64) (bool, bool, error) {
	return _Group.Contract.GetGInfo(&_Group.CallOpts, i)
}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 i) view returns(address)
func (_Group *GroupCaller) GetKManage(opts *bind.CallOpts, i uint64) (common.Address, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getKManage", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 i) view returns(address)
func (_Group *GroupSession) GetKManage(i uint64) (common.Address, error) {
	return _Group.Contract.GetKManage(&_Group.CallOpts, i)
}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 i) view returns(address)
func (_Group *GroupCallerSession) GetKManage(i uint64) (common.Address, error) {
	return _Group.Contract.GetKManage(&_Group.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_Group *GroupCaller) GetLevel(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getLevel", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_Group *GroupSession) GetLevel(i uint64) (uint8, error) {
	return _Group.Contract.GetLevel(&_Group.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_Group *GroupCallerSession) GetLevel(i uint64) (uint8, error) {
	return _Group.Contract.GetLevel(&_Group.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 i) view returns(uint256, uint256)
func (_Group *GroupCaller) GetPInfo(opts *bind.CallOpts, i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getPInfo", i)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 i) view returns(uint256, uint256)
func (_Group *GroupSession) GetPInfo(i uint64) (*big.Int, *big.Int, error) {
	return _Group.Contract.GetPInfo(&_Group.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 i) view returns(uint256, uint256)
func (_Group *GroupCallerSession) GetPInfo(i uint64) (*big.Int, *big.Int, error) {
	return _Group.Contract.GetPInfo(&_Group.CallOpts, i)
}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_Group *GroupCaller) GetSStra(opts *bind.CallOpts, i uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "getSStra", i)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_Group *GroupSession) GetSStra(i uint64) (uint8, uint8, error) {
	return _Group.Contract.GetSStra(&_Group.CallOpts, i)
}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_Group *GroupCallerSession) GetSStra(i uint64) (uint8, uint8, error) {
	return _Group.Contract.GetSStra(&_Group.CallOpts, i)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Group *GroupCaller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Group *GroupSession) Inst() (common.Address, error) {
	return _Group.Contract.Inst(&_Group.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Group *GroupCallerSession) Inst() (common.Address, error) {
	return _Group.Contract.Inst(&_Group.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Group *GroupCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Group *GroupSession) Owners(arg0 common.Address) (bool, error) {
	return _Group.Contract.Owners(&_Group.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Group *GroupCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Group.Contract.Owners(&_Group.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Group *GroupCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Group.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Group *GroupSession) Version() (uint16, error) {
	return _Group.Contract.Version(&_Group.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Group *GroupCallerSession) Version() (uint16, error) {
	return _Group.Contract.Version(&_Group.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_Group *GroupTransactor) Activate(opts *bind.TransactOpts, _gi uint64, _kc uint8) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "activate", _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_Group *GroupSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _Group.Contract.Activate(&_Group.TransactOpts, _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_Group *GroupTransactorSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _Group.Contract.Activate(&_Group.TransactOpts, _gi, _kc)
}

// Add0 is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Group *GroupTransactor) Add0(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "add0", _a, isSet, signs)
}

// Add0 is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Group *GroupSession) Add0(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Group.Contract.Add0(&_Group.TransactOpts, _a, isSet, signs)
}

// Add0 is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Group *GroupTransactorSession) Add0(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Group.Contract.Add0(&_Group.TransactOpts, _a, isSet, signs)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_Group *GroupTransactor) Ban(opts *bind.TransactOpts, _gi uint64, _isBan bool) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "ban", _gi, _isBan)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_Group *GroupSession) Ban(_gi uint64, _isBan bool) (*types.Transaction, error) {
	return _Group.Contract.Ban(&_Group.TransactOpts, _gi, _isBan)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_Group *GroupTransactorSession) Ban(_gi uint64, _isBan bool) (*types.Transaction, error) {
	return _Group.Contract.Ban(&_Group.TransactOpts, _gi, _isBan)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_Group *GroupTransactor) Create(opts *bind.TransactOpts, _level uint8, mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _Group.contract.Transact(opts, "create", _level, mr, _dc, _pc, _kr, _pr)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_Group *GroupSession) Create(_level uint8, mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _Group.Contract.Create(&_Group.TransactOpts, _level, mr, _dc, _pc, _kr, _pr)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_Group *GroupTransactorSession) Create(_level uint8, mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _Group.Contract.Create(&_Group.TransactOpts, _level, mr, _dc, _pc, _kr, _pr)
}

// GroupAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Group contract.
type GroupAddOwnerIterator struct {
	Event *GroupAddOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GroupAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupAddOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GroupAddOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GroupAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupAddOwner represents a AddOwner event raised by the Group contract.
type GroupAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Group *GroupFilterer) FilterAddOwner(opts *bind.FilterOpts) (*GroupAddOwnerIterator, error) {

	logs, sub, err := _Group.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &GroupAddOwnerIterator{contract: _Group.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Group *GroupFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *GroupAddOwner) (event.Subscription, error) {

	logs, sub, err := _Group.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupAddOwner)
				if err := _Group.contract.UnpackLog(event, "AddOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Group *GroupFilterer) ParseAddOwner(log types.Log) (*GroupAddOwner, error) {
	event := new(GroupAddOwner)
	if err := _Group.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GroupCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the Group contract.
type GroupCreateIterator struct {
	Event *GroupCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GroupCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GroupCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GroupCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GroupCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GroupCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GroupCreate represents a Create event raised by the Group contract.
type GroupCreate struct {
	GIndex uint64
	Level  uint8
	Mr     uint8
	Dc     uint8
	Pc     uint8
	Kr     *big.Int
	Pr     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_Group *GroupFilterer) FilterCreate(opts *bind.FilterOpts) (*GroupCreateIterator, error) {

	logs, sub, err := _Group.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &GroupCreateIterator{contract: _Group.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_Group *GroupFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *GroupCreate) (event.Subscription, error) {

	logs, sub, err := _Group.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GroupCreate)
				if err := _Group.contract.UnpackLog(event, "Create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreate is a log parse operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_Group *GroupFilterer) ParseCreate(log types.Log) (*GroupCreate, error) {
	event := new(GroupCreate)
	if err := _Group.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuthABI is the input ABI used to generate the binding from.
const IAuthABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"perm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAuthFuncSigs maps the 4-byte function signature to its string representation.
var IAuthFuncSigs = map[string]string{
	"a96bba9d": "perm(bytes32,bytes[5])",
}

// IAuth is an auto generated Go binding around an Ethereum contract.
type IAuth struct {
	IAuthCaller     // Read-only binding to the contract
	IAuthTransactor // Write-only binding to the contract
	IAuthFilterer   // Log filterer for contract events
}

// IAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthSession struct {
	Contract     *IAuth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthCallerSession struct {
	Contract *IAuthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthTransactorSession struct {
	Contract     *IAuthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthRaw struct {
	Contract *IAuth // Generic contract binding to access the raw methods on
}

// IAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthCallerRaw struct {
	Contract *IAuthCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthTransactorRaw struct {
	Contract *IAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuth creates a new instance of IAuth, bound to a specific deployed contract.
func NewIAuth(address common.Address, backend bind.ContractBackend) (*IAuth, error) {
	contract, err := bindIAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuth{IAuthCaller: IAuthCaller{contract: contract}, IAuthTransactor: IAuthTransactor{contract: contract}, IAuthFilterer: IAuthFilterer{contract: contract}}, nil
}

// NewIAuthCaller creates a new read-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthCaller(address common.Address, caller bind.ContractCaller) (*IAuthCaller, error) {
	contract, err := bindIAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthCaller{contract: contract}, nil
}

// NewIAuthTransactor creates a new write-only instance of IAuth, bound to a specific deployed contract.
func NewIAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthTransactor, error) {
	contract, err := bindIAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthTransactor{contract: contract}, nil
}

// NewIAuthFilterer creates a new log filterer instance of IAuth, bound to a specific deployed contract.
func NewIAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthFilterer, error) {
	contract, err := bindIAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthFilterer{contract: contract}, nil
}

// bindIAuth binds a generic wrapper to an already deployed contract.
func bindIAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.IAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.IAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuth *IAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuth *IAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuth *IAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuth.Contract.contract.Transact(opts, method, params...)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactor) Perm(opts *bind.TransactOpts, _h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.contract.Transact(opts, "perm", _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// Perm is a paid mutator transaction binding the contract method 0xa96bba9d.
//
// Solidity: function perm(bytes32 _h, bytes[5] signs) returns()
func (_IAuth *IAuthTransactorSession) Perm(_h [32]byte, signs [5][]byte) (*types.Transaction, error) {
	return _IAuth.Contract.Perm(&_IAuth.TransactOpts, _h, signs)
}

// IGroupABI is the input ABI used to generate the binding from.
const IGroupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dc\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"pc\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pr\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_kc\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_pm\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"checkG\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_dc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_pc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getKManage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getSStra\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IGroupFuncSigs maps the 4-byte function signature to its string representation.
var IGroupFuncSigs = map[string]string{
	"ad578aeb": "activate(uint64,uint8)",
	"4acc2468": "add(uint8,uint64,uint256)",
	"0f78c61a": "ban(uint64,bool)",
	"83889fe3": "checkG(uint64)",
	"eaf724ab": "create(uint8,uint8,uint8,uint8,uint256,uint256)",
	"059e783d": "getGCnt()",
	"38add6a1": "getGInfo(uint64)",
	"929629d2": "getKManage(uint64)",
	"770609a8": "getLevel(uint64)",
	"50aa2329": "getPInfo(uint64)",
	"dc698953": "getSStra(uint64)",
}

// IGroup is an auto generated Go binding around an Ethereum contract.
type IGroup struct {
	IGroupCaller     // Read-only binding to the contract
	IGroupTransactor // Write-only binding to the contract
	IGroupFilterer   // Log filterer for contract events
}

// IGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGroupSession struct {
	Contract     *IGroup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGroupCallerSession struct {
	Contract *IGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGroupTransactorSession struct {
	Contract     *IGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGroupRaw struct {
	Contract *IGroup // Generic contract binding to access the raw methods on
}

// IGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGroupCallerRaw struct {
	Contract *IGroupCaller // Generic read-only contract binding to access the raw methods on
}

// IGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGroupTransactorRaw struct {
	Contract *IGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGroup creates a new instance of IGroup, bound to a specific deployed contract.
func NewIGroup(address common.Address, backend bind.ContractBackend) (*IGroup, error) {
	contract, err := bindIGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGroup{IGroupCaller: IGroupCaller{contract: contract}, IGroupTransactor: IGroupTransactor{contract: contract}, IGroupFilterer: IGroupFilterer{contract: contract}}, nil
}

// NewIGroupCaller creates a new read-only instance of IGroup, bound to a specific deployed contract.
func NewIGroupCaller(address common.Address, caller bind.ContractCaller) (*IGroupCaller, error) {
	contract, err := bindIGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupCaller{contract: contract}, nil
}

// NewIGroupTransactor creates a new write-only instance of IGroup, bound to a specific deployed contract.
func NewIGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*IGroupTransactor, error) {
	contract, err := bindIGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupTransactor{contract: contract}, nil
}

// NewIGroupFilterer creates a new log filterer instance of IGroup, bound to a specific deployed contract.
func NewIGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*IGroupFilterer, error) {
	contract, err := bindIGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGroupFilterer{contract: contract}, nil
}

// bindIGroup binds a generic wrapper to an already deployed contract.
func bindIGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroup *IGroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroup.Contract.IGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroup *IGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroup.Contract.IGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroup *IGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroup.Contract.IGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroup *IGroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroup *IGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroup *IGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroup.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroup *IGroupCaller) Add(opts *bind.CallOpts, _rType uint8, _gi uint64, _pm *big.Int) error {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "add", _rType, _gi, _pm)

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroup *IGroupSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroup.Contract.Add(&_IGroup.CallOpts, _rType, _gi, _pm)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroup *IGroupCallerSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroup.Contract.Add(&_IGroup.CallOpts, _rType, _gi, _pm)
}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_IGroup *IGroupCaller) CheckG(opts *bind.CallOpts, i uint64) error {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "checkG", i)

	if err != nil {
		return err
	}

	return err

}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_IGroup *IGroupSession) CheckG(i uint64) error {
	return _IGroup.Contract.CheckG(&_IGroup.CallOpts, i)
}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_IGroup *IGroupCallerSession) CheckG(i uint64) error {
	return _IGroup.Contract.CheckG(&_IGroup.CallOpts, i)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroup *IGroupCaller) GetGCnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getGCnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroup *IGroupSession) GetGCnt() (uint64, error) {
	return _IGroup.Contract.GetGCnt(&_IGroup.CallOpts)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroup *IGroupCallerSession) GetGCnt() (uint64, error) {
	return _IGroup.Contract.GetGCnt(&_IGroup.CallOpts)
}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_IGroup *IGroupCaller) GetGInfo(opts *bind.CallOpts, i uint64) (bool, bool, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getGInfo", i)

	if err != nil {
		return *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_IGroup *IGroupSession) GetGInfo(i uint64) (bool, bool, error) {
	return _IGroup.Contract.GetGInfo(&_IGroup.CallOpts, i)
}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_IGroup *IGroupCallerSession) GetGInfo(i uint64) (bool, bool, error) {
	return _IGroup.Contract.GetGInfo(&_IGroup.CallOpts, i)
}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 _i) view returns(address)
func (_IGroup *IGroupCaller) GetKManage(opts *bind.CallOpts, _i uint64) (common.Address, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getKManage", _i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 _i) view returns(address)
func (_IGroup *IGroupSession) GetKManage(_i uint64) (common.Address, error) {
	return _IGroup.Contract.GetKManage(&_IGroup.CallOpts, _i)
}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 _i) view returns(address)
func (_IGroup *IGroupCallerSession) GetKManage(_i uint64) (common.Address, error) {
	return _IGroup.Contract.GetKManage(&_IGroup.CallOpts, _i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroup *IGroupCaller) GetLevel(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getLevel", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroup *IGroupSession) GetLevel(i uint64) (uint8, error) {
	return _IGroup.Contract.GetLevel(&_IGroup.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroup *IGroupCallerSession) GetLevel(i uint64) (uint8, error) {
	return _IGroup.Contract.GetLevel(&_IGroup.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroup *IGroupCaller) GetPInfo(opts *bind.CallOpts, _i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getPInfo", _i)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroup *IGroupSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroup.Contract.GetPInfo(&_IGroup.CallOpts, _i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroup *IGroupCallerSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroup.Contract.GetPInfo(&_IGroup.CallOpts, _i)
}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_IGroup *IGroupCaller) GetSStra(opts *bind.CallOpts, i uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _IGroup.contract.Call(opts, &out, "getSStra", i)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_IGroup *IGroupSession) GetSStra(i uint64) (uint8, uint8, error) {
	return _IGroup.Contract.GetSStra(&_IGroup.CallOpts, i)
}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_IGroup *IGroupCallerSession) GetSStra(i uint64) (uint8, uint8, error) {
	return _IGroup.Contract.GetSStra(&_IGroup.CallOpts, i)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroup *IGroupTransactor) Activate(opts *bind.TransactOpts, _gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "activate", _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroup *IGroupSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroup.Contract.Activate(&_IGroup.TransactOpts, _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroup *IGroupTransactorSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroup.Contract.Activate(&_IGroup.TransactOpts, _gi, _kc)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_IGroup *IGroupTransactor) Ban(opts *bind.TransactOpts, _gi uint64, _isBan bool) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "ban", _gi, _isBan)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_IGroup *IGroupSession) Ban(_gi uint64, _isBan bool) (*types.Transaction, error) {
	return _IGroup.Contract.Ban(&_IGroup.TransactOpts, _gi, _isBan)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_IGroup *IGroupTransactorSession) Ban(_gi uint64, _isBan bool) (*types.Transaction, error) {
	return _IGroup.Contract.Ban(&_IGroup.TransactOpts, _gi, _isBan)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IGroup *IGroupTransactor) Create(opts *bind.TransactOpts, _level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IGroup.contract.Transact(opts, "create", _level, _mr, _dc, _pc, _kr, _pr)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IGroup *IGroupSession) Create(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IGroup.Contract.Create(&_IGroup.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IGroup *IGroupTransactorSession) Create(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IGroup.Contract.Create(&_IGroup.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// IGroupCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the IGroup contract.
type IGroupCreateIterator struct {
	Event *IGroupCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGroupCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGroupCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGroupCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGroupCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGroupCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGroupCreate represents a Create event raised by the IGroup contract.
type IGroupCreate struct {
	GIndex uint64
	Level  uint8
	Mr     uint8
	Dc     uint8
	Pc     uint8
	Kr     *big.Int
	Pr     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_IGroup *IGroupFilterer) FilterCreate(opts *bind.FilterOpts) (*IGroupCreateIterator, error) {

	logs, sub, err := _IGroup.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &IGroupCreateIterator{contract: _IGroup.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_IGroup *IGroupFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *IGroupCreate) (event.Subscription, error) {

	logs, sub, err := _IGroup.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGroupCreate)
				if err := _IGroup.contract.UnpackLog(event, "Create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreate is a log parse operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_IGroup *IGroupFilterer) ParseCreate(log types.Log) (*IGroupCreate, error) {
	event := new(IGroupCreate)
	if err := _IGroup.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGroupGetterABI is the input ABI used to generate the binding from.
const IGroupGetterABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_rType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_pm\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"checkG\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGCnt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getGInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getKManage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getPInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"i\",\"type\":\"uint64\"}],\"name\":\"getSStra\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IGroupGetterFuncSigs maps the 4-byte function signature to its string representation.
var IGroupGetterFuncSigs = map[string]string{
	"4acc2468": "add(uint8,uint64,uint256)",
	"83889fe3": "checkG(uint64)",
	"059e783d": "getGCnt()",
	"38add6a1": "getGInfo(uint64)",
	"929629d2": "getKManage(uint64)",
	"770609a8": "getLevel(uint64)",
	"50aa2329": "getPInfo(uint64)",
	"dc698953": "getSStra(uint64)",
}

// IGroupGetter is an auto generated Go binding around an Ethereum contract.
type IGroupGetter struct {
	IGroupGetterCaller     // Read-only binding to the contract
	IGroupGetterTransactor // Write-only binding to the contract
	IGroupGetterFilterer   // Log filterer for contract events
}

// IGroupGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGroupGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGroupGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGroupGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGroupGetterSession struct {
	Contract     *IGroupGetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGroupGetterCallerSession struct {
	Contract *IGroupGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGroupGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGroupGetterTransactorSession struct {
	Contract     *IGroupGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGroupGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGroupGetterRaw struct {
	Contract *IGroupGetter // Generic contract binding to access the raw methods on
}

// IGroupGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGroupGetterCallerRaw struct {
	Contract *IGroupGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IGroupGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGroupGetterTransactorRaw struct {
	Contract *IGroupGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGroupGetter creates a new instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetter(address common.Address, backend bind.ContractBackend) (*IGroupGetter, error) {
	contract, err := bindIGroupGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGroupGetter{IGroupGetterCaller: IGroupGetterCaller{contract: contract}, IGroupGetterTransactor: IGroupGetterTransactor{contract: contract}, IGroupGetterFilterer: IGroupGetterFilterer{contract: contract}}, nil
}

// NewIGroupGetterCaller creates a new read-only instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetterCaller(address common.Address, caller bind.ContractCaller) (*IGroupGetterCaller, error) {
	contract, err := bindIGroupGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupGetterCaller{contract: contract}, nil
}

// NewIGroupGetterTransactor creates a new write-only instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IGroupGetterTransactor, error) {
	contract, err := bindIGroupGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupGetterTransactor{contract: contract}, nil
}

// NewIGroupGetterFilterer creates a new log filterer instance of IGroupGetter, bound to a specific deployed contract.
func NewIGroupGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IGroupGetterFilterer, error) {
	contract, err := bindIGroupGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGroupGetterFilterer{contract: contract}, nil
}

// bindIGroupGetter binds a generic wrapper to an already deployed contract.
func bindIGroupGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGroupGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupGetter *IGroupGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupGetter.Contract.IGroupGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupGetter *IGroupGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupGetter.Contract.IGroupGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupGetter *IGroupGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupGetter.Contract.IGroupGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupGetter *IGroupGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupGetter *IGroupGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupGetter *IGroupGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupGetter.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroupGetter *IGroupGetterCaller) Add(opts *bind.CallOpts, _rType uint8, _gi uint64, _pm *big.Int) error {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "add", _rType, _gi, _pm)

	if err != nil {
		return err
	}

	return err

}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroupGetter *IGroupGetterSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroupGetter.Contract.Add(&_IGroupGetter.CallOpts, _rType, _gi, _pm)
}

// Add is a free data retrieval call binding the contract method 0x4acc2468.
//
// Solidity: function add(uint8 _rType, uint64 _gi, uint256 _pm) view returns()
func (_IGroupGetter *IGroupGetterCallerSession) Add(_rType uint8, _gi uint64, _pm *big.Int) error {
	return _IGroupGetter.Contract.Add(&_IGroupGetter.CallOpts, _rType, _gi, _pm)
}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_IGroupGetter *IGroupGetterCaller) CheckG(opts *bind.CallOpts, i uint64) error {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "checkG", i)

	if err != nil {
		return err
	}

	return err

}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_IGroupGetter *IGroupGetterSession) CheckG(i uint64) error {
	return _IGroupGetter.Contract.CheckG(&_IGroupGetter.CallOpts, i)
}

// CheckG is a free data retrieval call binding the contract method 0x83889fe3.
//
// Solidity: function checkG(uint64 i) view returns()
func (_IGroupGetter *IGroupGetterCallerSession) CheckG(i uint64) error {
	return _IGroupGetter.Contract.CheckG(&_IGroupGetter.CallOpts, i)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroupGetter *IGroupGetterCaller) GetGCnt(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getGCnt")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroupGetter *IGroupGetterSession) GetGCnt() (uint64, error) {
	return _IGroupGetter.Contract.GetGCnt(&_IGroupGetter.CallOpts)
}

// GetGCnt is a free data retrieval call binding the contract method 0x059e783d.
//
// Solidity: function getGCnt() view returns(uint64)
func (_IGroupGetter *IGroupGetterCallerSession) GetGCnt() (uint64, error) {
	return _IGroupGetter.Contract.GetGCnt(&_IGroupGetter.CallOpts)
}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_IGroupGetter *IGroupGetterCaller) GetGInfo(opts *bind.CallOpts, i uint64) (bool, bool, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getGInfo", i)

	if err != nil {
		return *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_IGroupGetter *IGroupGetterSession) GetGInfo(i uint64) (bool, bool, error) {
	return _IGroupGetter.Contract.GetGInfo(&_IGroupGetter.CallOpts, i)
}

// GetGInfo is a free data retrieval call binding the contract method 0x38add6a1.
//
// Solidity: function getGInfo(uint64 i) view returns(bool, bool)
func (_IGroupGetter *IGroupGetterCallerSession) GetGInfo(i uint64) (bool, bool, error) {
	return _IGroupGetter.Contract.GetGInfo(&_IGroupGetter.CallOpts, i)
}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 _i) view returns(address)
func (_IGroupGetter *IGroupGetterCaller) GetKManage(opts *bind.CallOpts, _i uint64) (common.Address, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getKManage", _i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 _i) view returns(address)
func (_IGroupGetter *IGroupGetterSession) GetKManage(_i uint64) (common.Address, error) {
	return _IGroupGetter.Contract.GetKManage(&_IGroupGetter.CallOpts, _i)
}

// GetKManage is a free data retrieval call binding the contract method 0x929629d2.
//
// Solidity: function getKManage(uint64 _i) view returns(address)
func (_IGroupGetter *IGroupGetterCallerSession) GetKManage(_i uint64) (common.Address, error) {
	return _IGroupGetter.Contract.GetKManage(&_IGroupGetter.CallOpts, _i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterCaller) GetLevel(opts *bind.CallOpts, i uint64) (uint8, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getLevel", i)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterSession) GetLevel(i uint64) (uint8, error) {
	return _IGroupGetter.Contract.GetLevel(&_IGroupGetter.CallOpts, i)
}

// GetLevel is a free data retrieval call binding the contract method 0x770609a8.
//
// Solidity: function getLevel(uint64 i) view returns(uint8)
func (_IGroupGetter *IGroupGetterCallerSession) GetLevel(i uint64) (uint8, error) {
	return _IGroupGetter.Contract.GetLevel(&_IGroupGetter.CallOpts, i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterCaller) GetPInfo(opts *bind.CallOpts, _i uint64) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getPInfo", _i)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroupGetter.Contract.GetPInfo(&_IGroupGetter.CallOpts, _i)
}

// GetPInfo is a free data retrieval call binding the contract method 0x50aa2329.
//
// Solidity: function getPInfo(uint64 _i) view returns(uint256, uint256)
func (_IGroupGetter *IGroupGetterCallerSession) GetPInfo(_i uint64) (*big.Int, *big.Int, error) {
	return _IGroupGetter.Contract.GetPInfo(&_IGroupGetter.CallOpts, _i)
}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_IGroupGetter *IGroupGetterCaller) GetSStra(opts *bind.CallOpts, i uint64) (uint8, uint8, error) {
	var out []interface{}
	err := _IGroupGetter.contract.Call(opts, &out, "getSStra", i)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_IGroupGetter *IGroupGetterSession) GetSStra(i uint64) (uint8, uint8, error) {
	return _IGroupGetter.Contract.GetSStra(&_IGroupGetter.CallOpts, i)
}

// GetSStra is a free data retrieval call binding the contract method 0xdc698953.
//
// Solidity: function getSStra(uint64 i) view returns(uint8, uint8)
func (_IGroupGetter *IGroupGetterCallerSession) GetSStra(i uint64) (uint8, uint8, error) {
	return _IGroupGetter.Contract.GetSStra(&_IGroupGetter.CallOpts, i)
}

// IGroupSetterABI is the input ABI used to generate the binding from.
const IGroupSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"gIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"level\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dc\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"pc\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pr\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_kc\",\"type\":\"uint8\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gi\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isBan\",\"type\":\"bool\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_level\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_mr\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_dc\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_pc\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_kr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pr\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGroupSetterFuncSigs maps the 4-byte function signature to its string representation.
var IGroupSetterFuncSigs = map[string]string{
	"ad578aeb": "activate(uint64,uint8)",
	"0f78c61a": "ban(uint64,bool)",
	"eaf724ab": "create(uint8,uint8,uint8,uint8,uint256,uint256)",
}

// IGroupSetter is an auto generated Go binding around an Ethereum contract.
type IGroupSetter struct {
	IGroupSetterCaller     // Read-only binding to the contract
	IGroupSetterTransactor // Write-only binding to the contract
	IGroupSetterFilterer   // Log filterer for contract events
}

// IGroupSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGroupSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGroupSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGroupSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGroupSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGroupSetterSession struct {
	Contract     *IGroupSetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGroupSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGroupSetterCallerSession struct {
	Contract *IGroupSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGroupSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGroupSetterTransactorSession struct {
	Contract     *IGroupSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGroupSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGroupSetterRaw struct {
	Contract *IGroupSetter // Generic contract binding to access the raw methods on
}

// IGroupSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGroupSetterCallerRaw struct {
	Contract *IGroupSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IGroupSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGroupSetterTransactorRaw struct {
	Contract *IGroupSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGroupSetter creates a new instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetter(address common.Address, backend bind.ContractBackend) (*IGroupSetter, error) {
	contract, err := bindIGroupSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGroupSetter{IGroupSetterCaller: IGroupSetterCaller{contract: contract}, IGroupSetterTransactor: IGroupSetterTransactor{contract: contract}, IGroupSetterFilterer: IGroupSetterFilterer{contract: contract}}, nil
}

// NewIGroupSetterCaller creates a new read-only instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetterCaller(address common.Address, caller bind.ContractCaller) (*IGroupSetterCaller, error) {
	contract, err := bindIGroupSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupSetterCaller{contract: contract}, nil
}

// NewIGroupSetterTransactor creates a new write-only instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IGroupSetterTransactor, error) {
	contract, err := bindIGroupSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGroupSetterTransactor{contract: contract}, nil
}

// NewIGroupSetterFilterer creates a new log filterer instance of IGroupSetter, bound to a specific deployed contract.
func NewIGroupSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IGroupSetterFilterer, error) {
	contract, err := bindIGroupSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGroupSetterFilterer{contract: contract}, nil
}

// bindIGroupSetter binds a generic wrapper to an already deployed contract.
func bindIGroupSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGroupSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupSetter *IGroupSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupSetter.Contract.IGroupSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupSetter *IGroupSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupSetter.Contract.IGroupSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupSetter *IGroupSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupSetter.Contract.IGroupSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGroupSetter *IGroupSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGroupSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGroupSetter *IGroupSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGroupSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGroupSetter *IGroupSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGroupSetter.Contract.contract.Transact(opts, method, params...)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroupSetter *IGroupSetterTransactor) Activate(opts *bind.TransactOpts, _gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "activate", _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroupSetter *IGroupSetterSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Activate(&_IGroupSetter.TransactOpts, _gi, _kc)
}

// Activate is a paid mutator transaction binding the contract method 0xad578aeb.
//
// Solidity: function activate(uint64 _gi, uint8 _kc) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) Activate(_gi uint64, _kc uint8) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Activate(&_IGroupSetter.TransactOpts, _gi, _kc)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_IGroupSetter *IGroupSetterTransactor) Ban(opts *bind.TransactOpts, _gi uint64, _isBan bool) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "ban", _gi, _isBan)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_IGroupSetter *IGroupSetterSession) Ban(_gi uint64, _isBan bool) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Ban(&_IGroupSetter.TransactOpts, _gi, _isBan)
}

// Ban is a paid mutator transaction binding the contract method 0x0f78c61a.
//
// Solidity: function ban(uint64 _gi, bool _isBan) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) Ban(_gi uint64, _isBan bool) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Ban(&_IGroupSetter.TransactOpts, _gi, _isBan)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IGroupSetter *IGroupSetterTransactor) Create(opts *bind.TransactOpts, _level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.contract.Transact(opts, "create", _level, _mr, _dc, _pc, _kr, _pr)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IGroupSetter *IGroupSetterSession) Create(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Create(&_IGroupSetter.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// Create is a paid mutator transaction binding the contract method 0xeaf724ab.
//
// Solidity: function create(uint8 _level, uint8 _mr, uint8 _dc, uint8 _pc, uint256 _kr, uint256 _pr) returns()
func (_IGroupSetter *IGroupSetterTransactorSession) Create(_level uint8, _mr uint8, _dc uint8, _pc uint8, _kr *big.Int, _pr *big.Int) (*types.Transaction, error) {
	return _IGroupSetter.Contract.Create(&_IGroupSetter.TransactOpts, _level, _mr, _dc, _pc, _kr, _pr)
}

// IGroupSetterCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the IGroupSetter contract.
type IGroupSetterCreateIterator struct {
	Event *IGroupSetterCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGroupSetterCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGroupSetterCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGroupSetterCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGroupSetterCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGroupSetterCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGroupSetterCreate represents a Create event raised by the IGroupSetter contract.
type IGroupSetterCreate struct {
	GIndex uint64
	Level  uint8
	Mr     uint8
	Dc     uint8
	Pc     uint8
	Kr     *big.Int
	Pr     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_IGroupSetter *IGroupSetterFilterer) FilterCreate(opts *bind.FilterOpts) (*IGroupSetterCreateIterator, error) {

	logs, sub, err := _IGroupSetter.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &IGroupSetterCreateIterator{contract: _IGroupSetter.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_IGroupSetter *IGroupSetterFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *IGroupSetterCreate) (event.Subscription, error) {

	logs, sub, err := _IGroupSetter.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGroupSetterCreate)
				if err := _IGroupSetter.contract.UnpackLog(event, "Create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreate is a log parse operation binding the contract event 0x2dec96e260e12f1bbe324e5f5471344d4fb705cf6995ff4a6e136b9da6460d99.
//
// Solidity: event Create(uint64 gIndex, uint8 level, uint8 mr, uint8 dc, uint8 pc, uint256 kr, uint256 pr)
func (_IGroupSetter *IGroupSetterFilterer) ParseCreate(log types.Log) (*IGroupSetterCreate, error) {
	event := new(IGroupSetterCreate)
	if err := _IGroupSetter.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInstanceABI is the input ABI used to generate the binding from.
const IInstanceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Alter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"instances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IInstanceFuncSigs maps the 4-byte function signature to its string representation.
var IInstanceFuncSigs = map[string]string{
	"3ec7d5b9": "instances(uint8)",
}

// IInstance is an auto generated Go binding around an Ethereum contract.
type IInstance struct {
	IInstanceCaller     // Read-only binding to the contract
	IInstanceTransactor // Write-only binding to the contract
	IInstanceFilterer   // Log filterer for contract events
}

// IInstanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInstanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInstanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInstanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInstanceSession struct {
	Contract     *IInstance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInstanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInstanceCallerSession struct {
	Contract *IInstanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IInstanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInstanceTransactorSession struct {
	Contract     *IInstanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInstanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInstanceRaw struct {
	Contract *IInstance // Generic contract binding to access the raw methods on
}

// IInstanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInstanceCallerRaw struct {
	Contract *IInstanceCaller // Generic read-only contract binding to access the raw methods on
}

// IInstanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInstanceTransactorRaw struct {
	Contract *IInstanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInstance creates a new instance of IInstance, bound to a specific deployed contract.
func NewIInstance(address common.Address, backend bind.ContractBackend) (*IInstance, error) {
	contract, err := bindIInstance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInstance{IInstanceCaller: IInstanceCaller{contract: contract}, IInstanceTransactor: IInstanceTransactor{contract: contract}, IInstanceFilterer: IInstanceFilterer{contract: contract}}, nil
}

// NewIInstanceCaller creates a new read-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceCaller(address common.Address, caller bind.ContractCaller) (*IInstanceCaller, error) {
	contract, err := bindIInstance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceCaller{contract: contract}, nil
}

// NewIInstanceTransactor creates a new write-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IInstanceTransactor, error) {
	contract, err := bindIInstance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceTransactor{contract: contract}, nil
}

// NewIInstanceFilterer creates a new log filterer instance of IInstance, bound to a specific deployed contract.
func NewIInstanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IInstanceFilterer, error) {
	contract, err := bindIInstance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInstanceFilterer{contract: contract}, nil
}

// bindIInstance binds a generic wrapper to an already deployed contract.
func bindIInstance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInstanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.IInstanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transact(opts, method, params...)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCaller) Instances(opts *bind.CallOpts, _type uint8) (common.Address, error) {
	var out []interface{}
	err := _IInstance.contract.Call(opts, &out, "instances", _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCallerSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// IInstanceAlterIterator is returned from FilterAlter and is used to iterate over the raw logs and unpacked data for Alter events raised by the IInstance contract.
type IInstanceAlterIterator struct {
	Event *IInstanceAlter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IInstanceAlterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInstanceAlter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IInstanceAlter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IInstanceAlterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInstanceAlterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInstanceAlter represents a Alter event raised by the IInstance contract.
type IInstanceAlter struct {
	Type uint8
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlter is a free log retrieval operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) FilterAlter(opts *bind.FilterOpts) (*IInstanceAlterIterator, error) {

	logs, sub, err := _IInstance.contract.FilterLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return &IInstanceAlterIterator{contract: _IInstance.contract, event: "Alter", logs: logs, sub: sub}, nil
}

// WatchAlter is a free log subscription operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) WatchAlter(opts *bind.WatchOpts, sink chan<- *IInstanceAlter) (event.Subscription, error) {

	logs, sub, err := _IInstance.contract.WatchLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInstanceAlter)
				if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAlter is a log parse operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) ParseAlter(log types.Log) (*IInstanceAlter, error) {
	event := new(IInstanceAlter)
	if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageABI is the input ABI used to generate the binding from.
const IKmanageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manageRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKmanageFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageFuncSigs = map[string]string{
	"024130e4": "addCnt(uint64,uint64)",
	"50cbb46f": "addKeeper(uint64)",
	"55d3d7ef": "addProfit(uint8,uint256)",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"dc0e7f45": "getK(uint64)",
	"39f2db96": "getKCnt()",
	"4f3c2eab": "getPf(uint8)",
	"6d23f6c8": "manageRate()",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IKmanage is an auto generated Go binding around an Ethereum contract.
type IKmanage struct {
	IKmanageCaller     // Read-only binding to the contract
	IKmanageTransactor // Write-only binding to the contract
	IKmanageFilterer   // Log filterer for contract events
}

// IKmanageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageSession struct {
	Contract     *IKmanage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageCallerSession struct {
	Contract *IKmanageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IKmanageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageTransactorSession struct {
	Contract     *IKmanageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IKmanageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageRaw struct {
	Contract *IKmanage // Generic contract binding to access the raw methods on
}

// IKmanageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageCallerRaw struct {
	Contract *IKmanageCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageTransactorRaw struct {
	Contract *IKmanageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanage creates a new instance of IKmanage, bound to a specific deployed contract.
func NewIKmanage(address common.Address, backend bind.ContractBackend) (*IKmanage, error) {
	contract, err := bindIKmanage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanage{IKmanageCaller: IKmanageCaller{contract: contract}, IKmanageTransactor: IKmanageTransactor{contract: contract}, IKmanageFilterer: IKmanageFilterer{contract: contract}}, nil
}

// NewIKmanageCaller creates a new read-only instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageCaller(address common.Address, caller bind.ContractCaller) (*IKmanageCaller, error) {
	contract, err := bindIKmanage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageCaller{contract: contract}, nil
}

// NewIKmanageTransactor creates a new write-only instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageTransactor, error) {
	contract, err := bindIKmanage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageTransactor{contract: contract}, nil
}

// NewIKmanageFilterer creates a new log filterer instance of IKmanage, bound to a specific deployed contract.
func NewIKmanageFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageFilterer, error) {
	contract, err := bindIKmanage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageFilterer{contract: contract}, nil
}

// bindIKmanage binds a generic wrapper to an already deployed contract.
func bindIKmanage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanage *IKmanageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanage.Contract.IKmanageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanage *IKmanageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanage.Contract.IKmanageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanage *IKmanageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanage.Contract.IKmanageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanage *IKmanageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanage *IKmanageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanage *IKmanageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanage.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageCaller) BalanceOf(opts *bind.CallOpts, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "balanceOf", _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanage.Contract.BalanceOf(&_IKmanage.CallOpts, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanage *IKmanageCallerSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanage.Contract.BalanceOf(&_IKmanage.CallOpts, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageCaller) GetK(opts *bind.CallOpts, _i uint64) (uint64, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getK", _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageSession) GetK(_i uint64) (uint64, error) {
	return _IKmanage.Contract.GetK(&_IKmanage.CallOpts, _i)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanage *IKmanageCallerSession) GetK(_i uint64) (uint64, error) {
	return _IKmanage.Contract.GetK(&_IKmanage.CallOpts, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanage *IKmanageCaller) GetKCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getKCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanage *IKmanageSession) GetKCnt() (uint8, error) {
	return _IKmanage.Contract.GetKCnt(&_IKmanage.CallOpts)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanage *IKmanageCallerSession) GetKCnt() (uint8, error) {
	return _IKmanage.Contract.GetKCnt(&_IKmanage.CallOpts)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanage *IKmanageCaller) GetPf(opts *bind.CallOpts, _ti uint8) (uint64, *big.Int, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "getPf", _ti)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanage *IKmanageSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanage.Contract.GetPf(&_IKmanage.CallOpts, _ti)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanage *IKmanageCallerSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanage.Contract.GetPf(&_IKmanage.CallOpts, _ti)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanage *IKmanageCaller) ManageRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanage.contract.Call(opts, &out, "manageRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanage *IKmanageSession) ManageRate() (uint8, error) {
	return _IKmanage.Contract.ManageRate(&_IKmanage.CallOpts)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanage *IKmanageCallerSession) ManageRate() (uint8, error) {
	return _IKmanage.Contract.ManageRate(&_IKmanage.CallOpts)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageTransactor) AddCnt(opts *bind.TransactOpts, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addCnt", _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddCnt(&_IKmanage.TransactOpts, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanage *IKmanageTransactorSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddCnt(&_IKmanage.TransactOpts, _ki, _cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanage *IKmanageTransactor) AddKeeper(opts *bind.TransactOpts, _ki uint64) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addKeeper", _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanage *IKmanageSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddKeeper(&_IKmanage.TransactOpts, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanage *IKmanageTransactorSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanage.Contract.AddKeeper(&_IKmanage.TransactOpts, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanage *IKmanageTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "addProfit", _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanage *IKmanageSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.AddProfit(&_IKmanage.TransactOpts, _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanage *IKmanageTransactorSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.AddProfit(&_IKmanage.TransactOpts, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanage *IKmanageTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.contract.Transact(opts, "withdraw", _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanage *IKmanageSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.Withdraw(&_IKmanage.TransactOpts, _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanage *IKmanageTransactorSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanage.Contract.Withdraw(&_IKmanage.TransactOpts, _ki, _ti, money)
}

// IKmanageAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the IKmanage contract.
type IKmanageAddCntIterator struct {
	Event *IKmanageAddCnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IKmanageAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageAddCnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IKmanageAddCnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IKmanageAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageAddCnt represents a AddCnt event raised by the IKmanage contract.
type IKmanageAddCnt struct {
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) FilterAddCnt(opts *bind.FilterOpts, ki []uint64) (*IKmanageAddCntIterator, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanage.contract.FilterLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageAddCntIterator{contract: _IKmanage.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *IKmanageAddCnt, ki []uint64) (event.Subscription, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanage.contract.WatchLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageAddCnt)
				if err := _IKmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddCnt is a log parse operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanage *IKmanageFilterer) ParseAddCnt(log types.Log) (*IKmanageAddCnt, error) {
	event := new(IKmanageAddCnt)
	if err := _IKmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the IKmanage contract.
type IKmanageAddProfitIterator struct {
	Event *IKmanageAddProfit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IKmanageAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageAddProfit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IKmanageAddProfit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IKmanageAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageAddProfit represents a AddProfit event raised by the IKmanage contract.
type IKmanageAddProfit struct {
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) FilterAddProfit(opts *bind.FilterOpts, ti []uint8) (*IKmanageAddProfitIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanage.contract.FilterLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageAddProfitIterator{contract: _IKmanage.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *IKmanageAddProfit, ti []uint8) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanage.contract.WatchLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageAddProfit)
				if err := _IKmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddProfit is a log parse operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanage *IKmanageFilterer) ParseAddProfit(log types.Log) (*IKmanageAddProfit, error) {
	event := new(IKmanageAddProfit)
	if err := _IKmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageGetterABI is the input ABI used to generate the binding from.
const IKmanageGetterABI = "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manageRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IKmanageGetterFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageGetterFuncSigs = map[string]string{
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"dc0e7f45": "getK(uint64)",
	"39f2db96": "getKCnt()",
	"4f3c2eab": "getPf(uint8)",
	"6d23f6c8": "manageRate()",
}

// IKmanageGetter is an auto generated Go binding around an Ethereum contract.
type IKmanageGetter struct {
	IKmanageGetterCaller     // Read-only binding to the contract
	IKmanageGetterTransactor // Write-only binding to the contract
	IKmanageGetterFilterer   // Log filterer for contract events
}

// IKmanageGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageGetterSession struct {
	Contract     *IKmanageGetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageGetterCallerSession struct {
	Contract *IKmanageGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKmanageGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageGetterTransactorSession struct {
	Contract     *IKmanageGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKmanageGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageGetterRaw struct {
	Contract *IKmanageGetter // Generic contract binding to access the raw methods on
}

// IKmanageGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageGetterCallerRaw struct {
	Contract *IKmanageGetterCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageGetterTransactorRaw struct {
	Contract *IKmanageGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanageGetter creates a new instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetter(address common.Address, backend bind.ContractBackend) (*IKmanageGetter, error) {
	contract, err := bindIKmanageGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetter{IKmanageGetterCaller: IKmanageGetterCaller{contract: contract}, IKmanageGetterTransactor: IKmanageGetterTransactor{contract: contract}, IKmanageGetterFilterer: IKmanageGetterFilterer{contract: contract}}, nil
}

// NewIKmanageGetterCaller creates a new read-only instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterCaller(address common.Address, caller bind.ContractCaller) (*IKmanageGetterCaller, error) {
	contract, err := bindIKmanageGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterCaller{contract: contract}, nil
}

// NewIKmanageGetterTransactor creates a new write-only instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageGetterTransactor, error) {
	contract, err := bindIKmanageGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterTransactor{contract: contract}, nil
}

// NewIKmanageGetterFilterer creates a new log filterer instance of IKmanageGetter, bound to a specific deployed contract.
func NewIKmanageGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageGetterFilterer, error) {
	contract, err := bindIKmanageGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageGetterFilterer{contract: contract}, nil
}

// bindIKmanageGetter binds a generic wrapper to an already deployed contract.
func bindIKmanageGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageGetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageGetter *IKmanageGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageGetter.Contract.IKmanageGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageGetter *IKmanageGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.IKmanageGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageGetter *IKmanageGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.IKmanageGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageGetter *IKmanageGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageGetter *IKmanageGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageGetter *IKmanageGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageGetter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterCaller) BalanceOf(opts *bind.CallOpts, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "balanceOf", _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanageGetter.Contract.BalanceOf(&_IKmanageGetter.CallOpts, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_IKmanageGetter *IKmanageGetterCallerSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _IKmanageGetter.Contract.BalanceOf(&_IKmanageGetter.CallOpts, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterCaller) GetK(opts *bind.CallOpts, _i uint64) (uint64, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getK", _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterSession) GetK(_i uint64) (uint64, error) {
	return _IKmanageGetter.Contract.GetK(&_IKmanageGetter.CallOpts, _i)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetK(_i uint64) (uint64, error) {
	return _IKmanageGetter.Contract.GetK(&_IKmanageGetter.CallOpts, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCaller) GetKCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getKCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterSession) GetKCnt() (uint8, error) {
	return _IKmanageGetter.Contract.GetKCnt(&_IKmanageGetter.CallOpts)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetKCnt() (uint8, error) {
	return _IKmanageGetter.Contract.GetKCnt(&_IKmanageGetter.CallOpts)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterCaller) GetPf(opts *bind.CallOpts, _ti uint8) (uint64, *big.Int, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "getPf", _ti)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanageGetter.Contract.GetPf(&_IKmanageGetter.CallOpts, _ti)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_IKmanageGetter *IKmanageGetterCallerSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _IKmanageGetter.Contract.GetPf(&_IKmanageGetter.CallOpts, _ti)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCaller) ManageRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IKmanageGetter.contract.Call(opts, &out, "manageRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterSession) ManageRate() (uint8, error) {
	return _IKmanageGetter.Contract.ManageRate(&_IKmanageGetter.CallOpts)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_IKmanageGetter *IKmanageGetterCallerSession) ManageRate() (uint8, error) {
	return _IKmanageGetter.Contract.ManageRate(&_IKmanageGetter.CallOpts)
}

// IKmanageSetterABI is the input ABI used to generate the binding from.
const IKmanageSetterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IKmanageSetterFuncSigs maps the 4-byte function signature to its string representation.
var IKmanageSetterFuncSigs = map[string]string{
	"024130e4": "addCnt(uint64,uint64)",
	"50cbb46f": "addKeeper(uint64)",
	"55d3d7ef": "addProfit(uint8,uint256)",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// IKmanageSetter is an auto generated Go binding around an Ethereum contract.
type IKmanageSetter struct {
	IKmanageSetterCaller     // Read-only binding to the contract
	IKmanageSetterTransactor // Write-only binding to the contract
	IKmanageSetterFilterer   // Log filterer for contract events
}

// IKmanageSetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKmanageSetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKmanageSetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKmanageSetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKmanageSetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKmanageSetterSession struct {
	Contract     *IKmanageSetter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKmanageSetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKmanageSetterCallerSession struct {
	Contract *IKmanageSetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IKmanageSetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKmanageSetterTransactorSession struct {
	Contract     *IKmanageSetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IKmanageSetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKmanageSetterRaw struct {
	Contract *IKmanageSetter // Generic contract binding to access the raw methods on
}

// IKmanageSetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKmanageSetterCallerRaw struct {
	Contract *IKmanageSetterCaller // Generic read-only contract binding to access the raw methods on
}

// IKmanageSetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKmanageSetterTransactorRaw struct {
	Contract *IKmanageSetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKmanageSetter creates a new instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetter(address common.Address, backend bind.ContractBackend) (*IKmanageSetter, error) {
	contract, err := bindIKmanageSetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetter{IKmanageSetterCaller: IKmanageSetterCaller{contract: contract}, IKmanageSetterTransactor: IKmanageSetterTransactor{contract: contract}, IKmanageSetterFilterer: IKmanageSetterFilterer{contract: contract}}, nil
}

// NewIKmanageSetterCaller creates a new read-only instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterCaller(address common.Address, caller bind.ContractCaller) (*IKmanageSetterCaller, error) {
	contract, err := bindIKmanageSetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterCaller{contract: contract}, nil
}

// NewIKmanageSetterTransactor creates a new write-only instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterTransactor(address common.Address, transactor bind.ContractTransactor) (*IKmanageSetterTransactor, error) {
	contract, err := bindIKmanageSetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterTransactor{contract: contract}, nil
}

// NewIKmanageSetterFilterer creates a new log filterer instance of IKmanageSetter, bound to a specific deployed contract.
func NewIKmanageSetterFilterer(address common.Address, filterer bind.ContractFilterer) (*IKmanageSetterFilterer, error) {
	contract, err := bindIKmanageSetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterFilterer{contract: contract}, nil
}

// bindIKmanageSetter binds a generic wrapper to an already deployed contract.
func bindIKmanageSetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKmanageSetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageSetter *IKmanageSetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageSetter.Contract.IKmanageSetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageSetter *IKmanageSetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.IKmanageSetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageSetter *IKmanageSetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.IKmanageSetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKmanageSetter *IKmanageSetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKmanageSetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKmanageSetter *IKmanageSetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKmanageSetter *IKmanageSetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.contract.Transact(opts, method, params...)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddCnt(opts *bind.TransactOpts, _ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addCnt", _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddCnt(&_IKmanageSetter.TransactOpts, _ki, _cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 _cnt) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddCnt(_ki uint64, _cnt uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddCnt(&_IKmanageSetter.TransactOpts, _ki, _cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddKeeper(opts *bind.TransactOpts, _ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addKeeper", _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddKeeper(&_IKmanageSetter.TransactOpts, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddKeeper(&_IKmanageSetter.TransactOpts, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "addProfit", _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddProfit(&_IKmanageSetter.TransactOpts, _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_IKmanageSetter *IKmanageSetterTransactorSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.AddProfit(&_IKmanageSetter.TransactOpts, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.contract.Transact(opts, "withdraw", _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.Withdraw(&_IKmanageSetter.TransactOpts, _ki, _ti, money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 money) returns(uint256)
func (_IKmanageSetter *IKmanageSetterTransactorSession) Withdraw(_ki uint64, _ti uint8, money *big.Int) (*types.Transaction, error) {
	return _IKmanageSetter.Contract.Withdraw(&_IKmanageSetter.TransactOpts, _ki, _ti, money)
}

// IKmanageSetterAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the IKmanageSetter contract.
type IKmanageSetterAddCntIterator struct {
	Event *IKmanageSetterAddCnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IKmanageSetterAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageSetterAddCnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IKmanageSetterAddCnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IKmanageSetterAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageSetterAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageSetterAddCnt represents a AddCnt event raised by the IKmanageSetter contract.
type IKmanageSetterAddCnt struct {
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) FilterAddCnt(opts *bind.FilterOpts, ki []uint64) (*IKmanageSetterAddCntIterator, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.FilterLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterAddCntIterator{contract: _IKmanageSetter.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *IKmanageSetterAddCnt, ki []uint64) (event.Subscription, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.WatchLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageSetterAddCnt)
				if err := _IKmanageSetter.contract.UnpackLog(event, "AddCnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddCnt is a log parse operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_IKmanageSetter *IKmanageSetterFilterer) ParseAddCnt(log types.Log) (*IKmanageSetterAddCnt, error) {
	event := new(IKmanageSetterAddCnt)
	if err := _IKmanageSetter.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKmanageSetterAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the IKmanageSetter contract.
type IKmanageSetterAddProfitIterator struct {
	Event *IKmanageSetterAddProfit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IKmanageSetterAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKmanageSetterAddProfit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IKmanageSetterAddProfit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IKmanageSetterAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKmanageSetterAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKmanageSetterAddProfit represents a AddProfit event raised by the IKmanageSetter contract.
type IKmanageSetterAddProfit struct {
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) FilterAddProfit(opts *bind.FilterOpts, ti []uint8) (*IKmanageSetterAddProfitIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.FilterLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return &IKmanageSetterAddProfitIterator{contract: _IKmanageSetter.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *IKmanageSetterAddProfit, ti []uint8) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _IKmanageSetter.contract.WatchLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKmanageSetterAddProfit)
				if err := _IKmanageSetter.contract.UnpackLog(event, "AddProfit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddProfit is a log parse operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_IKmanageSetter *IKmanageSetterFilterer) ParseAddProfit(log types.Log) (*IKmanageSetterAddProfit, error) {
	event := new(IKmanageSetterAddProfit)
	if err := _IKmanageSetter.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOwnerABI is the input ABI used to generate the binding from.
const IOwnerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"}]"

// IOwner is an auto generated Go binding around an Ethereum contract.
type IOwner struct {
	IOwnerCaller     // Read-only binding to the contract
	IOwnerTransactor // Write-only binding to the contract
	IOwnerFilterer   // Log filterer for contract events
}

// IOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOwnerSession struct {
	Contract     *IOwner           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOwnerCallerSession struct {
	Contract *IOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOwnerTransactorSession struct {
	Contract     *IOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOwnerRaw struct {
	Contract *IOwner // Generic contract binding to access the raw methods on
}

// IOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOwnerCallerRaw struct {
	Contract *IOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// IOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOwnerTransactorRaw struct {
	Contract *IOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOwner creates a new instance of IOwner, bound to a specific deployed contract.
func NewIOwner(address common.Address, backend bind.ContractBackend) (*IOwner, error) {
	contract, err := bindIOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOwner{IOwnerCaller: IOwnerCaller{contract: contract}, IOwnerTransactor: IOwnerTransactor{contract: contract}, IOwnerFilterer: IOwnerFilterer{contract: contract}}, nil
}

// NewIOwnerCaller creates a new read-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerCaller(address common.Address, caller bind.ContractCaller) (*IOwnerCaller, error) {
	contract, err := bindIOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerCaller{contract: contract}, nil
}

// NewIOwnerTransactor creates a new write-only instance of IOwner, bound to a specific deployed contract.
func NewIOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*IOwnerTransactor, error) {
	contract, err := bindIOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOwnerTransactor{contract: contract}, nil
}

// NewIOwnerFilterer creates a new log filterer instance of IOwner, bound to a specific deployed contract.
func NewIOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*IOwnerFilterer, error) {
	contract, err := bindIOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOwnerFilterer{contract: contract}, nil
}

// bindIOwner binds a generic wrapper to an already deployed contract.
func bindIOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.IOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.IOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOwner *IOwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOwner *IOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOwner *IOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOwner.Contract.contract.Transact(opts, method, params...)
}

// IOwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the IOwner contract.
type IOwnerAddOwnerIterator struct {
	Event *IOwnerAddOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IOwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOwnerAddOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IOwnerAddOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IOwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOwnerAddOwner represents a AddOwner event raised by the IOwner contract.
type IOwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*IOwnerAddOwnerIterator, error) {

	logs, sub, err := _IOwner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &IOwnerAddOwnerIterator{contract: _IOwner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *IOwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _IOwner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOwnerAddOwner)
				if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_IOwner *IOwnerFilterer) ParseAddOwner(log types.Log) (*IOwnerAddOwner, error) {
	event := new(IOwnerAddOwner)
	if err := _IOwner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolABI is the input ABI used to generate the binding from.
const IPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Inflow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"Outflow\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"inflow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"outflow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IPoolFuncSigs maps the 4-byte function signature to its string representation.
var IPoolFuncSigs = map[string]string{
	"368007fe": "inflow(address,address,uint256)",
	"530345bb": "outflow(address,address,uint256)",
}

// IPool is an auto generated Go binding around an Ethereum contract.
type IPool struct {
	IPoolCaller     // Read-only binding to the contract
	IPoolTransactor // Write-only binding to the contract
	IPoolFilterer   // Log filterer for contract events
}

// IPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPoolSession struct {
	Contract     *IPool            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPoolCallerSession struct {
	Contract *IPoolCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPoolTransactorSession struct {
	Contract     *IPoolTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPoolRaw struct {
	Contract *IPool // Generic contract binding to access the raw methods on
}

// IPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPoolCallerRaw struct {
	Contract *IPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPoolTransactorRaw struct {
	Contract *IPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPool creates a new instance of IPool, bound to a specific deployed contract.
func NewIPool(address common.Address, backend bind.ContractBackend) (*IPool, error) {
	contract, err := bindIPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPool{IPoolCaller: IPoolCaller{contract: contract}, IPoolTransactor: IPoolTransactor{contract: contract}, IPoolFilterer: IPoolFilterer{contract: contract}}, nil
}

// NewIPoolCaller creates a new read-only instance of IPool, bound to a specific deployed contract.
func NewIPoolCaller(address common.Address, caller bind.ContractCaller) (*IPoolCaller, error) {
	contract, err := bindIPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolCaller{contract: contract}, nil
}

// NewIPoolTransactor creates a new write-only instance of IPool, bound to a specific deployed contract.
func NewIPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IPoolTransactor, error) {
	contract, err := bindIPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPoolTransactor{contract: contract}, nil
}

// NewIPoolFilterer creates a new log filterer instance of IPool, bound to a specific deployed contract.
func NewIPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IPoolFilterer, error) {
	contract, err := bindIPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPoolFilterer{contract: contract}, nil
}

// bindIPool binds a generic wrapper to an already deployed contract.
func bindIPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.IPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.IPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPool *IPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPool *IPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPool *IPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPool.Contract.contract.Transact(opts, method, params...)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolTransactor) Inflow(opts *bind.TransactOpts, tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "inflow", tAddr, from, money)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolSession) Inflow(tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Inflow(&_IPool.TransactOpts, tAddr, from, money)
}

// Inflow is a paid mutator transaction binding the contract method 0x368007fe.
//
// Solidity: function inflow(address tAddr, address from, uint256 money) payable returns()
func (_IPool *IPoolTransactorSession) Inflow(tAddr common.Address, from common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Inflow(&_IPool.TransactOpts, tAddr, from, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolTransactor) Outflow(opts *bind.TransactOpts, tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.contract.Transact(opts, "outflow", tAddr, to, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolSession) Outflow(tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Outflow(&_IPool.TransactOpts, tAddr, to, money)
}

// Outflow is a paid mutator transaction binding the contract method 0x530345bb.
//
// Solidity: function outflow(address tAddr, address to, uint256 money) payable returns()
func (_IPool *IPoolTransactorSession) Outflow(tAddr common.Address, to common.Address, money *big.Int) (*types.Transaction, error) {
	return _IPool.Contract.Outflow(&_IPool.TransactOpts, tAddr, to, money)
}

// IPoolInflowIterator is returned from FilterInflow and is used to iterate over the raw logs and unpacked data for Inflow events raised by the IPool contract.
type IPoolInflowIterator struct {
	Event *IPoolInflow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPoolInflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolInflow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPoolInflow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPoolInflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolInflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolInflow represents a Inflow event raised by the IPool contract.
type IPoolInflow struct {
	From  common.Address
	TAddr common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInflow is a free log retrieval operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) FilterInflow(opts *bind.FilterOpts, from []common.Address) (*IPoolInflowIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Inflow", fromRule)
	if err != nil {
		return nil, err
	}
	return &IPoolInflowIterator{contract: _IPool.contract, event: "Inflow", logs: logs, sub: sub}, nil
}

// WatchInflow is a free log subscription operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) WatchInflow(opts *bind.WatchOpts, sink chan<- *IPoolInflow, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Inflow", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolInflow)
				if err := _IPool.contract.UnpackLog(event, "Inflow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInflow is a log parse operation binding the contract event 0xde52ea03a3979fafeaf8ea9d7fe6b3ddc6a95e9e8c0922562db3a047c0d72578.
//
// Solidity: event Inflow(address indexed from, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) ParseInflow(log types.Log) (*IPoolInflow, error) {
	event := new(IPoolInflow)
	if err := _IPool.contract.UnpackLog(event, "Inflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPoolOutflowIterator is returned from FilterOutflow and is used to iterate over the raw logs and unpacked data for Outflow events raised by the IPool contract.
type IPoolOutflowIterator struct {
	Event *IPoolOutflow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IPoolOutflowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPoolOutflow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IPoolOutflow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IPoolOutflowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPoolOutflowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPoolOutflow represents a Outflow event raised by the IPool contract.
type IPoolOutflow struct {
	To    common.Address
	TAddr common.Address
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOutflow is a free log retrieval operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) FilterOutflow(opts *bind.FilterOpts, to []common.Address) (*IPoolOutflowIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.FilterLogs(opts, "Outflow", toRule)
	if err != nil {
		return nil, err
	}
	return &IPoolOutflowIterator{contract: _IPool.contract, event: "Outflow", logs: logs, sub: sub}, nil
}

// WatchOutflow is a free log subscription operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) WatchOutflow(opts *bind.WatchOpts, sink chan<- *IPoolOutflow, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IPool.contract.WatchLogs(opts, "Outflow", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPoolOutflow)
				if err := _IPool.contract.UnpackLog(event, "Outflow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutflow is a log parse operation binding the contract event 0x4eb2adb2eba0bb5546cc2e8d62ae0c32e7b6fb40567d53be433eb4b17f5f996e.
//
// Solidity: event Outflow(address indexed to, address tAddr, uint256 money)
func (_IPool *IPoolFilterer) ParseOutflow(log types.Log) (*IPoolOutflow, error) {
	event := new(IPoolOutflow)
	if err := _IPool.contract.UnpackLog(event, "Outflow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KmanageABI is the input ABI used to generate the binding from.
const KmanageABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"mr\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"ki\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"AddCnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"ti\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"AddProfit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cnt\",\"type\":\"uint64\"}],\"name\":\"addCnt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_money\",\"type\":\"uint256\"}],\"name\":\"addProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_i\",\"type\":\"uint64\"}],\"name\":\"getK\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKCnt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"}],\"name\":\"getPf\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manageRate\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_ki\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"_ti\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KmanageFuncSigs maps the 4-byte function signature to its string representation.
var KmanageFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"024130e4": "addCnt(uint64,uint64)",
	"50cbb46f": "addKeeper(uint64)",
	"55d3d7ef": "addProfit(uint8,uint256)",
	"de9375f2": "auth()",
	"fc3ba0ad": "balanceOf(uint64,uint8)",
	"dc0e7f45": "getK(uint64)",
	"39f2db96": "getKCnt()",
	"4f3c2eab": "getPf(uint8)",
	"6d23f6c8": "manageRate()",
	"022914a7": "owners(address)",
	"54fd4d50": "version()",
	"259c6d5e": "withdraw(uint64,uint8,uint256)",
}

// KmanageBin is the compiled bytecode used for deploying new contracts.
var KmanageBin = "0x608060405260018054600160a01b600160f81b03191665049d4002000160a11b17905534801561002e57600080fd5b5060405161125f38038061125f83398101604081905261004d91610081565b6001805460ff909216600160b01b02600162ff000160a01b03199092166001600160a01b03909316929092171790556100cd565b6000806040838503121561009457600080fd5b82516001600160a01b03811681146100ab57600080fd5b602084015190925060ff811681146100c257600080fd5b809150509250929050565b611183806100dc6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806350cbb46f1161008c5780636d23f6c8116100665780636d23f6c81461020d578063dc0e7f4514610221578063de9375f21461024c578063fc3ba0ad1461027757600080fd5b806350cbb46f146101bf57806354fd4d50146101d257806355d3d7ef146101fa57600080fd5b8063022914a7146100d4578063024130e41461010c578063259c6d5e1461012157806339f2db96146101425780634bf1b134146101585780634f3c2eab1461016b575b600080fd5b6100f76100e2366004610c7f565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61011f61011a366004610cb1565b61029f565b005b61013461012f366004610cf5565b61041f565b604051908152602001610103565b6003545b60405160ff9091168152602001610103565b61011f610166366004610d9f565b61080b565b6101a0610179366004610eb4565b60ff16600090815260056020526040902080546001909101546001600160401b0390911691565b604080516001600160401b039093168352602083019190915201610103565b61011f6101cd366004610ecf565b610984565b6001546101e790600160a01b900461ffff1681565b60405161ffff9091168152602001610103565b61011f610208366004610eea565b610ab6565b60015461014690600160b01b900460ff1681565b61023461022f366004610ecf565b610b44565b6040516001600160401b039091168152602001610103565b60015461025f906001600160a01b031681565b6040516001600160a01b039091168152602001610103565b61028a610285366004610f14565b610b92565b60408051928352602083019190915201610103565b3360009081526020819052604090205460ff166102d75760405162461bcd60e51b81526004016102ce90610f3e565b60405180910390fd5b6001600160401b03808316600090815260046020526040902054166103335760405162461bcd60e51b815260206004820152601260248201527139a7b93232b91d34b636329031b0b63632b960711b60448201526064016102ce565b6001600160401b0380831660009081526004602052604081208054849391929161035f91859116610f77565b92506101000a8154816001600160401b0302191690836001600160401b0316021790555080600260008282829054906101000a90046001600160401b03166103a79190610f77565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550816001600160401b03167f5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c48260405161041391906001600160401b0391909116815260200190565b60405180910390a25050565b3360009081526020819052604081205460ff1661044e5760405162461bcd60e51b81526004016102ce90610f3e565b6001600160401b03808516600090815260046020526040812054909116900361047957506000610804565b60015460ff841660009081526005602052604090205442916001600160401b03600160b81b9091048116916104af911683610fa2565b6001600160401b031611156107905760ff84166000908152600560205260408120805467ffffffffffffffff19166001600160401b03938416178155600254600190910154919283926105059290911690610fe0565b905060005b600354811015610771576000600460006003848154811061052d5761052d610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a900482168452830193909352604090910190205461057391168461100a565b90506002600460006003858154811061058e5761058e610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a90048216845283019390935260409091019020546105d4929116611029565b6105df906001610f77565b60046000600385815481106105f6576105f6610ff4565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b03166001600160401b03166001600160401b0316815260200190815260200160002060006101000a8154816001600160401b0302191690836001600160401b03160217905550600460006003848154811061067a5761067a610ff4565b600091825260208083206004830401546001600160401b0360086003909416939093026101000a90048216845283019390935260409091019020546106c0911685610f77565b93508060066000600385815481106106da576106da610ff4565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b03168352828101939093526040918201812060ff8c1682529092528120805490919061072d90849061104f565b909155505060ff871660009081526005602052604081206001018054839290610757908490611067565b9091555082915061076990508161107e565b91505061050a565b50506002805467ffffffffffffffff19166001600160401b0383161790555b6001600160401b038516600090815260066020908152604080832060ff88168452909152902054808411156107c3578093505b6001600160401b038616600090815260066020908152604080832060ff89168452909152812080548692906107f9908490611067565b909155508493505050505b9392505050565b823b60008190036108535760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b60448201526064016102ce565b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d906108e09084908790600401611097565b600060405180830381600087803b1580156108fa57600080fd5b505af115801561090e573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b3360009081526020819052604090205460ff166109b35760405162461bcd60e51b81526004016102ce90610f3e565b6001600160401b038082166000908152600460205260409020541615610a055760405162461bcd60e51b81526020600482015260076024820152661ac8195e1a5cdd60ca1b60448201526064016102ce565b600380546001818101835560048083047fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160401b03808816600896909716959095026101000a86810290860219909116179055600093845260205260408320805467ffffffffffffffff191690911790556002805490911691610a8e83611127565b91906101000a8154816001600160401b0302191690836001600160401b031602179055505050565b3360009081526020819052604090205460ff16610ae55760405162461bcd60e51b81526004016102ce90610f3e565b60ff821660009081526005602052604081206001018054839290610b0a90849061104f565b909155505060405181815260ff8316907fe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e890602001610413565b60006003826001600160401b031681548110610b6257610b62610ff4565b90600052602060002090600491828204019190066008029054906101000a90046001600160401b03169050919050565b6001600160401b03808316600081815260066020908152604080832060ff871680855290835281842054948452600483528184205460025491855260059093529083206001015492948594938593849390831692610bf292911690610fe0565b610bfc919061100a565b60015460ff88166000908152600560205260409020549192506001600160401b03600160b81b909104811691610c33911642611067565b1115610c4a57610c43818461104f565b9250610c57565b610c54818361104f565b91505b50909590945092505050565b80356001600160a01b0381168114610c7a57600080fd5b919050565b600060208284031215610c9157600080fd5b61080482610c63565b80356001600160401b0381168114610c7a57600080fd5b60008060408385031215610cc457600080fd5b610ccd83610c9a565b9150610cdb60208401610c9a565b90509250929050565b803560ff81168114610c7a57600080fd5b600080600060608486031215610d0a57600080fd5b610d1384610c9a565b9250610d2160208501610ce4565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b60405160a081016001600160401b0381118282101715610d6957610d69610d31565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610d9757610d97610d31565b604052919050565b600080600060608486031215610db457600080fd5b610dbd84610c63565b92506020808501358015158114610dd357600080fd5b925060408501356001600160401b0380821115610def57600080fd5b8187019150601f8881840112610e0457600080fd5b610e0c610d47565b8060a085018b811115610e1e57600080fd5b855b81811015610ea257803586811115610e385760008081fd5b87018581018e13610e495760008081fd5b803587811115610e5b57610e5b610d31565b610e6c818801601f19168b01610d6f565b8181528f8b838501011115610e815760008081fd5b818b84018c83013760009181018b0191909152855250928701928701610e20565b50508096505050505050509250925092565b600060208284031215610ec657600080fd5b61080482610ce4565b600060208284031215610ee157600080fd5b61080482610c9a565b60008060408385031215610efd57600080fd5b610f0683610ce4565b946020939093013593505050565b60008060408385031215610f2757600080fd5b610f3083610c9a565b9150610cdb60208401610ce4565b6020808252600990820152683737ba1037bbb732b960b91b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b03808316818516808303821115610f9957610f99610f61565b01949350505050565b60006001600160401b0383811690831681811015610fc257610fc2610f61565b039392505050565b634e487b7160e01b600052601260045260246000fd5b600082610fef57610fef610fca565b500490565b634e487b7160e01b600052603260045260246000fd5b600081600019048311821515161561102457611024610f61565b500290565b60006001600160401b038084168061104357611043610fca565b92169190910492915050565b6000821982111561106257611062610f61565b500190565b60008282101561107957611079610f61565b500390565b60006001820161109057611090610f61565b5060010190565b8281526040602080830182905260009160e08401919084018584805b600581101561111957878603603f1901845282518051808852835b818110156110e95782810188015189820189015287016110ce565b818111156110f9578488838b0101525b50601f01601f1916969096018501955092840192918401916001016110b3565b509398975050505050505050565b60006001600160401b0380831681810361114357611143610f61565b600101939250505056fea2646970667358221220d61b4c43fbc2afc6384ec21a987e1f29503e04a5de72681e2e2579554deeec9b64736f6c634300080e0033"

// DeployKmanage deploys a new Ethereum contract, binding an instance of Kmanage to it.
func DeployKmanage(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address, mr uint8) (common.Address, *types.Transaction, *Kmanage, error) {
	parsed, err := abi.JSON(strings.NewReader(KmanageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KmanageBin), backend, _a, mr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kmanage{KmanageCaller: KmanageCaller{contract: contract}, KmanageTransactor: KmanageTransactor{contract: contract}, KmanageFilterer: KmanageFilterer{contract: contract}}, nil
}

// Kmanage is an auto generated Go binding around an Ethereum contract.
type Kmanage struct {
	KmanageCaller     // Read-only binding to the contract
	KmanageTransactor // Write-only binding to the contract
	KmanageFilterer   // Log filterer for contract events
}

// KmanageCaller is an auto generated read-only Go binding around an Ethereum contract.
type KmanageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KmanageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KmanageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KmanageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KmanageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KmanageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KmanageSession struct {
	Contract     *Kmanage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KmanageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KmanageCallerSession struct {
	Contract *KmanageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KmanageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KmanageTransactorSession struct {
	Contract     *KmanageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KmanageRaw is an auto generated low-level Go binding around an Ethereum contract.
type KmanageRaw struct {
	Contract *Kmanage // Generic contract binding to access the raw methods on
}

// KmanageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KmanageCallerRaw struct {
	Contract *KmanageCaller // Generic read-only contract binding to access the raw methods on
}

// KmanageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KmanageTransactorRaw struct {
	Contract *KmanageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKmanage creates a new instance of Kmanage, bound to a specific deployed contract.
func NewKmanage(address common.Address, backend bind.ContractBackend) (*Kmanage, error) {
	contract, err := bindKmanage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kmanage{KmanageCaller: KmanageCaller{contract: contract}, KmanageTransactor: KmanageTransactor{contract: contract}, KmanageFilterer: KmanageFilterer{contract: contract}}, nil
}

// NewKmanageCaller creates a new read-only instance of Kmanage, bound to a specific deployed contract.
func NewKmanageCaller(address common.Address, caller bind.ContractCaller) (*KmanageCaller, error) {
	contract, err := bindKmanage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KmanageCaller{contract: contract}, nil
}

// NewKmanageTransactor creates a new write-only instance of Kmanage, bound to a specific deployed contract.
func NewKmanageTransactor(address common.Address, transactor bind.ContractTransactor) (*KmanageTransactor, error) {
	contract, err := bindKmanage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KmanageTransactor{contract: contract}, nil
}

// NewKmanageFilterer creates a new log filterer instance of Kmanage, bound to a specific deployed contract.
func NewKmanageFilterer(address common.Address, filterer bind.ContractFilterer) (*KmanageFilterer, error) {
	contract, err := bindKmanage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KmanageFilterer{contract: contract}, nil
}

// bindKmanage binds a generic wrapper to an already deployed contract.
func bindKmanage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KmanageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kmanage *KmanageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kmanage.Contract.KmanageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kmanage *KmanageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kmanage.Contract.KmanageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kmanage *KmanageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kmanage.Contract.KmanageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kmanage *KmanageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kmanage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kmanage *KmanageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kmanage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kmanage *KmanageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kmanage.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Kmanage *KmanageCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Kmanage *KmanageSession) Auth() (common.Address, error) {
	return _Kmanage.Contract.Auth(&_Kmanage.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Kmanage *KmanageCallerSession) Auth() (common.Address, error) {
	return _Kmanage.Contract.Auth(&_Kmanage.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_Kmanage *KmanageCaller) BalanceOf(opts *bind.CallOpts, _ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "balanceOf", _ki, _ti)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_Kmanage *KmanageSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _Kmanage.Contract.BalanceOf(&_Kmanage.CallOpts, _ki, _ti)
}

// BalanceOf is a free data retrieval call binding the contract method 0xfc3ba0ad.
//
// Solidity: function balanceOf(uint64 _ki, uint8 _ti) view returns(uint256, uint256)
func (_Kmanage *KmanageCallerSession) BalanceOf(_ki uint64, _ti uint8) (*big.Int, *big.Int, error) {
	return _Kmanage.Contract.BalanceOf(&_Kmanage.CallOpts, _ki, _ti)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_Kmanage *KmanageCaller) GetK(opts *bind.CallOpts, _i uint64) (uint64, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "getK", _i)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_Kmanage *KmanageSession) GetK(_i uint64) (uint64, error) {
	return _Kmanage.Contract.GetK(&_Kmanage.CallOpts, _i)
}

// GetK is a free data retrieval call binding the contract method 0xdc0e7f45.
//
// Solidity: function getK(uint64 _i) view returns(uint64)
func (_Kmanage *KmanageCallerSession) GetK(_i uint64) (uint64, error) {
	return _Kmanage.Contract.GetK(&_Kmanage.CallOpts, _i)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_Kmanage *KmanageCaller) GetKCnt(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "getKCnt")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_Kmanage *KmanageSession) GetKCnt() (uint8, error) {
	return _Kmanage.Contract.GetKCnt(&_Kmanage.CallOpts)
}

// GetKCnt is a free data retrieval call binding the contract method 0x39f2db96.
//
// Solidity: function getKCnt() view returns(uint8)
func (_Kmanage *KmanageCallerSession) GetKCnt() (uint8, error) {
	return _Kmanage.Contract.GetKCnt(&_Kmanage.CallOpts)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_Kmanage *KmanageCaller) GetPf(opts *bind.CallOpts, _ti uint8) (uint64, *big.Int, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "getPf", _ti)

	if err != nil {
		return *new(uint64), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_Kmanage *KmanageSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _Kmanage.Contract.GetPf(&_Kmanage.CallOpts, _ti)
}

// GetPf is a free data retrieval call binding the contract method 0x4f3c2eab.
//
// Solidity: function getPf(uint8 _ti) view returns(uint64, uint256)
func (_Kmanage *KmanageCallerSession) GetPf(_ti uint8) (uint64, *big.Int, error) {
	return _Kmanage.Contract.GetPf(&_Kmanage.CallOpts, _ti)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_Kmanage *KmanageCaller) ManageRate(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "manageRate")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_Kmanage *KmanageSession) ManageRate() (uint8, error) {
	return _Kmanage.Contract.ManageRate(&_Kmanage.CallOpts)
}

// ManageRate is a free data retrieval call binding the contract method 0x6d23f6c8.
//
// Solidity: function manageRate() view returns(uint8)
func (_Kmanage *KmanageCallerSession) ManageRate() (uint8, error) {
	return _Kmanage.Contract.ManageRate(&_Kmanage.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Kmanage *KmanageCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Kmanage *KmanageSession) Owners(arg0 common.Address) (bool, error) {
	return _Kmanage.Contract.Owners(&_Kmanage.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Kmanage *KmanageCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Kmanage.Contract.Owners(&_Kmanage.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Kmanage *KmanageCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Kmanage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Kmanage *KmanageSession) Version() (uint16, error) {
	return _Kmanage.Contract.Version(&_Kmanage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Kmanage *KmanageCallerSession) Version() (uint16, error) {
	return _Kmanage.Contract.Version(&_Kmanage.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Kmanage *KmanageTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Kmanage *KmanageSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Kmanage.Contract.Add(&_Kmanage.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Kmanage *KmanageTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Kmanage.Contract.Add(&_Kmanage.TransactOpts, _a, isSet, signs)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 cnt) returns()
func (_Kmanage *KmanageTransactor) AddCnt(opts *bind.TransactOpts, _ki uint64, cnt uint64) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "addCnt", _ki, cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 cnt) returns()
func (_Kmanage *KmanageSession) AddCnt(_ki uint64, cnt uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddCnt(&_Kmanage.TransactOpts, _ki, cnt)
}

// AddCnt is a paid mutator transaction binding the contract method 0x024130e4.
//
// Solidity: function addCnt(uint64 _ki, uint64 cnt) returns()
func (_Kmanage *KmanageTransactorSession) AddCnt(_ki uint64, cnt uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddCnt(&_Kmanage.TransactOpts, _ki, cnt)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_Kmanage *KmanageTransactor) AddKeeper(opts *bind.TransactOpts, _ki uint64) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "addKeeper", _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_Kmanage *KmanageSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddKeeper(&_Kmanage.TransactOpts, _ki)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x50cbb46f.
//
// Solidity: function addKeeper(uint64 _ki) returns()
func (_Kmanage *KmanageTransactorSession) AddKeeper(_ki uint64) (*types.Transaction, error) {
	return _Kmanage.Contract.AddKeeper(&_Kmanage.TransactOpts, _ki)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_Kmanage *KmanageTransactor) AddProfit(opts *bind.TransactOpts, _ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "addProfit", _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_Kmanage *KmanageSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.AddProfit(&_Kmanage.TransactOpts, _ti, _money)
}

// AddProfit is a paid mutator transaction binding the contract method 0x55d3d7ef.
//
// Solidity: function addProfit(uint8 _ti, uint256 _money) returns()
func (_Kmanage *KmanageTransactorSession) AddProfit(_ti uint8, _money *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.AddProfit(&_Kmanage.TransactOpts, _ti, _money)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 amount) returns(uint256)
func (_Kmanage *KmanageTransactor) Withdraw(opts *bind.TransactOpts, _ki uint64, _ti uint8, amount *big.Int) (*types.Transaction, error) {
	return _Kmanage.contract.Transact(opts, "withdraw", _ki, _ti, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 amount) returns(uint256)
func (_Kmanage *KmanageSession) Withdraw(_ki uint64, _ti uint8, amount *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.Withdraw(&_Kmanage.TransactOpts, _ki, _ti, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x259c6d5e.
//
// Solidity: function withdraw(uint64 _ki, uint8 _ti, uint256 amount) returns(uint256)
func (_Kmanage *KmanageTransactorSession) Withdraw(_ki uint64, _ti uint8, amount *big.Int) (*types.Transaction, error) {
	return _Kmanage.Contract.Withdraw(&_Kmanage.TransactOpts, _ki, _ti, amount)
}

// KmanageAddCntIterator is returned from FilterAddCnt and is used to iterate over the raw logs and unpacked data for AddCnt events raised by the Kmanage contract.
type KmanageAddCntIterator struct {
	Event *KmanageAddCnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KmanageAddCntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KmanageAddCnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KmanageAddCnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KmanageAddCntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KmanageAddCntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KmanageAddCnt represents a AddCnt event raised by the Kmanage contract.
type KmanageAddCnt struct {
	Ki  uint64
	Cnt uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddCnt is a free log retrieval operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_Kmanage *KmanageFilterer) FilterAddCnt(opts *bind.FilterOpts, ki []uint64) (*KmanageAddCntIterator, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _Kmanage.contract.FilterLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return &KmanageAddCntIterator{contract: _Kmanage.contract, event: "AddCnt", logs: logs, sub: sub}, nil
}

// WatchAddCnt is a free log subscription operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_Kmanage *KmanageFilterer) WatchAddCnt(opts *bind.WatchOpts, sink chan<- *KmanageAddCnt, ki []uint64) (event.Subscription, error) {

	var kiRule []interface{}
	for _, kiItem := range ki {
		kiRule = append(kiRule, kiItem)
	}

	logs, sub, err := _Kmanage.contract.WatchLogs(opts, "AddCnt", kiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KmanageAddCnt)
				if err := _Kmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddCnt is a log parse operation binding the contract event 0x5372d6aad551334f508508499c71755ebc6cde46b83fc1f944f1f0ae33cbb4c4.
//
// Solidity: event AddCnt(uint64 indexed ki, uint64 cnt)
func (_Kmanage *KmanageFilterer) ParseAddCnt(log types.Log) (*KmanageAddCnt, error) {
	event := new(KmanageAddCnt)
	if err := _Kmanage.contract.UnpackLog(event, "AddCnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KmanageAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Kmanage contract.
type KmanageAddOwnerIterator struct {
	Event *KmanageAddOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KmanageAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KmanageAddOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KmanageAddOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KmanageAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KmanageAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KmanageAddOwner represents a AddOwner event raised by the Kmanage contract.
type KmanageAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Kmanage *KmanageFilterer) FilterAddOwner(opts *bind.FilterOpts) (*KmanageAddOwnerIterator, error) {

	logs, sub, err := _Kmanage.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &KmanageAddOwnerIterator{contract: _Kmanage.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Kmanage *KmanageFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *KmanageAddOwner) (event.Subscription, error) {

	logs, sub, err := _Kmanage.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KmanageAddOwner)
				if err := _Kmanage.contract.UnpackLog(event, "AddOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Kmanage *KmanageFilterer) ParseAddOwner(log types.Log) (*KmanageAddOwner, error) {
	event := new(KmanageAddOwner)
	if err := _Kmanage.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KmanageAddProfitIterator is returned from FilterAddProfit and is used to iterate over the raw logs and unpacked data for AddProfit events raised by the Kmanage contract.
type KmanageAddProfitIterator struct {
	Event *KmanageAddProfit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KmanageAddProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KmanageAddProfit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KmanageAddProfit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KmanageAddProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KmanageAddProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KmanageAddProfit represents a AddProfit event raised by the Kmanage contract.
type KmanageAddProfit struct {
	Ti    uint8
	Money *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddProfit is a free log retrieval operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_Kmanage *KmanageFilterer) FilterAddProfit(opts *bind.FilterOpts, ti []uint8) (*KmanageAddProfitIterator, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _Kmanage.contract.FilterLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return &KmanageAddProfitIterator{contract: _Kmanage.contract, event: "AddProfit", logs: logs, sub: sub}, nil
}

// WatchAddProfit is a free log subscription operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_Kmanage *KmanageFilterer) WatchAddProfit(opts *bind.WatchOpts, sink chan<- *KmanageAddProfit, ti []uint8) (event.Subscription, error) {

	var tiRule []interface{}
	for _, tiItem := range ti {
		tiRule = append(tiRule, tiItem)
	}

	logs, sub, err := _Kmanage.contract.WatchLogs(opts, "AddProfit", tiRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KmanageAddProfit)
				if err := _Kmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddProfit is a log parse operation binding the contract event 0xe5fd8ec20bdfeb1fadd32ec3786545a1db18d74f952effda5c6cd50af7c2e9e8.
//
// Solidity: event AddProfit(uint8 indexed ti, uint256 money)
func (_Kmanage *KmanageFilterer) ParseAddProfit(log types.Log) (*KmanageAddProfit, error) {
	event := new(KmanageAddProfit)
	if err := _Kmanage.contract.UnpackLog(event, "AddProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerABI is the input ABI used to generate the binding from.
const OwnerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"AddOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"},{\"internalType\":\"bytes[5]\",\"name\":\"signs\",\"type\":\"bytes[5]\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnerFuncSigs maps the 4-byte function signature to its string representation.
var OwnerFuncSigs = map[string]string{
	"4bf1b134": "add(address,bool,bytes[5])",
	"de9375f2": "auth()",
	"022914a7": "owners(address)",
}

// OwnerBin is the compiled bytecode used for deploying new contracts.
var OwnerBin = "0x608060405234801561001057600080fd5b5060405161055838038061055883398101604081905261002f91610054565b600180546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b6104c5806100936000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063022914a7146100465780634bf1b1341461007e578063de9375f214610093575b600080fd5b610069610054366004610257565b60006020819052908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61009161008c3660046102e9565b6100be565b005b6001546100a6906001600160a01b031681565b6040516001600160a01b039091168152602001610075565b823b600081900361010a5760405162461bcd60e51b81526020600482015260126024820152713732b2b21031b7b73a3930b1ba1030b2323960711b604482015260640160405180910390fd5b6040516bffffffffffffffffffffffff1930606090811b821660208401526218591960ea1b603484015286901b16603782015283151560f81b604b820152600090604c0160408051601f1981840301815290829052805160209091012060015463a96bba9d60e01b83529092506001600160a01b03169063a96bba9d9061019790849087906004016103ff565b600060405180830381600087803b1580156101b157600080fd5b505af11580156101c5573d6000803e3d6000fd5b5050604080516001600160a01b038916815287151560208201527f938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807935001905060405180910390a15050506001600160a01b03919091166000908152602081905260409020805460ff1916911515919091179055565b80356001600160a01b038116811461025257600080fd5b919050565b60006020828403121561026957600080fd5b6102728261023b565b9392505050565b634e487b7160e01b600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156102b2576102b2610279565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156102e1576102e1610279565b604052919050565b6000806000606084860312156102fe57600080fd5b6103078461023b565b9250602080850135801515811461031d57600080fd5b9250604085013567ffffffffffffffff8082111561033a57600080fd5b8187019150601f888184011261034f57600080fd5b61035761028f565b8060a085018b81111561036957600080fd5b855b818110156103ed578035868111156103835760008081fd5b87018581018e136103945760008081fd5b8035878111156103a6576103a6610279565b6103b7818801601f19168b016102b8565b8181528f8b8385010111156103cc5760008081fd5b818b84018c83013760009181018b019190915285525092870192870161036b565b50508096505050505050509250925092565b8281526040602080830182905260009160e08401919084018584805b600581101561048157878603603f1901845282518051808852835b81811015610451578281018801518982018901528701610436565b81811115610461578488838b0101525b50601f01601f19169690960185019550928401929184019160010161041b565b50939897505050505050505056fea2646970667358221220e4109531dbabeb33ccbfa77f28019b7953ba57d32bea8b64099ee0dd97784c0464736f6c634300080e0033"

// DeployOwner deploys a new Ethereum contract, binding an instance of Owner to it.
func DeployOwner(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Owner, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnerBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// Owner is an auto generated Go binding around an Ethereum contract.
type Owner struct {
	OwnerCaller     // Read-only binding to the contract
	OwnerTransactor // Write-only binding to the contract
	OwnerFilterer   // Log filterer for contract events
}

// OwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerSession struct {
	Contract     *Owner            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerCallerSession struct {
	Contract *OwnerCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerTransactorSession struct {
	Contract     *OwnerTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerRaw struct {
	Contract *Owner // Generic contract binding to access the raw methods on
}

// OwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerCallerRaw struct {
	Contract *OwnerCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerTransactorRaw struct {
	Contract *OwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwner creates a new instance of Owner, bound to a specific deployed contract.
func NewOwner(address common.Address, backend bind.ContractBackend) (*Owner, error) {
	contract, err := bindOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owner{OwnerCaller: OwnerCaller{contract: contract}, OwnerTransactor: OwnerTransactor{contract: contract}, OwnerFilterer: OwnerFilterer{contract: contract}}, nil
}

// NewOwnerCaller creates a new read-only instance of Owner, bound to a specific deployed contract.
func NewOwnerCaller(address common.Address, caller bind.ContractCaller) (*OwnerCaller, error) {
	contract, err := bindOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerCaller{contract: contract}, nil
}

// NewOwnerTransactor creates a new write-only instance of Owner, bound to a specific deployed contract.
func NewOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerTransactor, error) {
	contract, err := bindOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerTransactor{contract: contract}, nil
}

// NewOwnerFilterer creates a new log filterer instance of Owner, bound to a specific deployed contract.
func NewOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerFilterer, error) {
	contract, err := bindOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerFilterer{contract: contract}, nil
}

// bindOwner binds a generic wrapper to an already deployed contract.
func bindOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.OwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.OwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owner *OwnerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owner *OwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owner *OwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owner.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Owner *OwnerCallerSession) Auth() (common.Address, error) {
	return _Owner.Contract.Auth(&_Owner.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Owner.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) view returns(bool)
func (_Owner *OwnerCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Owner.Contract.Owners(&_Owner.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactor) Add(opts *bind.TransactOpts, _a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.contract.Transact(opts, "add", _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// Add is a paid mutator transaction binding the contract method 0x4bf1b134.
//
// Solidity: function add(address _a, bool isSet, bytes[5] signs) returns()
func (_Owner *OwnerTransactorSession) Add(_a common.Address, isSet bool, signs [5][]byte) (*types.Transaction, error) {
	return _Owner.Contract.Add(&_Owner.TransactOpts, _a, isSet, signs)
}

// OwnerAddOwnerIterator is returned from FilterAddOwner and is used to iterate over the raw logs and unpacked data for AddOwner events raised by the Owner contract.
type OwnerAddOwnerIterator struct {
	Event *OwnerAddOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnerAddOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAddOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnerAddOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnerAddOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAddOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAddOwner represents a AddOwner event raised by the Owner contract.
type OwnerAddOwner struct {
	From  common.Address
	IsSet bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddOwner is a free log retrieval operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) FilterAddOwner(opts *bind.FilterOpts) (*OwnerAddOwnerIterator, error) {

	logs, sub, err := _Owner.contract.FilterLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return &OwnerAddOwnerIterator{contract: _Owner.contract, event: "AddOwner", logs: logs, sub: sub}, nil
}

// WatchAddOwner is a free log subscription operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) WatchAddOwner(opts *bind.WatchOpts, sink chan<- *OwnerAddOwner) (event.Subscription, error) {

	logs, sub, err := _Owner.contract.WatchLogs(opts, "AddOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAddOwner)
				if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddOwner is a log parse operation binding the contract event 0x938b2a24c98e4e542127ffa74a91e48942c2bddccae3b6d75f82bfda7bbc0807.
//
// Solidity: event AddOwner(address from, bool isSet)
func (_Owner *OwnerFilterer) ParseAddOwner(log types.Log) (*OwnerAddOwner, error) {
	event := new(OwnerAddOwner)
	if err := _Owner.contract.UnpackLog(event, "AddOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
