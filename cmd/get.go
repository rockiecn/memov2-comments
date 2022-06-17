package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	com "github.com/memoio/contractsv2/common"
	"github.com/memoio/contractsv2/go_contracts/fs"
	"github.com/memoio/contractsv2/go_contracts/getter"
	"github.com/memoio/contractsv2/go_contracts/group"
	inst "github.com/memoio/contractsv2/go_contracts/instance"
	"github.com/memoio/contractsv2/go_contracts/issuance"
	"github.com/memoio/contractsv2/go_contracts/kmanage"
	"github.com/memoio/contractsv2/go_contracts/pledge"
	"github.com/memoio/contractsv2/go_contracts/pool"
	"github.com/memoio/contractsv2/go_contracts/role"
	"github.com/memoio/contractsv2/go_contracts/token"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	ErrRevert = xerrors.New(": maybe the value is too large and exceeds the limit of the array")
)

// AdminCmd is about contracts functions called by anyone to get information about contract
var GetCmd = &cli.Command{
	Name:  "get",
	Usage: "Get information about contract",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "endPoint",
			Aliases: []string{"ep"},
			Value:   "dev", //默认值为devchain
			Usage:   "indicate which chain(dev, test, product), default dev",
		},
	},
	Subcommands: []*cli.Command{
		// getter
		getterVersionCmd,
		instAddrCmd,
		// instance
		instanceCmd,
		// token
		getTVersionCmd,
		getTACmd,
		getTICmd,
		getTCntCmd,
		// filesys
		fsVersionCmd,
		taxRateCmd,
		fsBalCmd,
		gsInfoCmd,
		fsInfoCmd,
		storeInfoCmd,
		settleInfoCmd,
		// group
		gVersionCmd,
		gInstCmd,
		gCntCmd,
		gInfoCmd,
		gLevelCmd,
		gPInfoCmd,
		kmanageCmd,
		gStoreStraCmd,
		// issuance
		issuVersionCmd,
		issuRewardCmd,
		issuStoreCmd,
		issuStageCmd,
		expireCmd,
		// kmanage
		kmVersionCmd,
		manageRateCmd,
		kCntCmd,
		getKICmd,
		getProfitInfoCmd,
		kmBalCmd,
		// pledge
		pleVersionCmd,
		pleInstCmd,
		pleTInfoCmd,
		pleRewardsCmd,
		pleBalCmd,
		pledgeCmd,
		totalPledgeCmd,
		// pool
		poolVersionCmd,
		// role
		roleVersionCmd,
		rinfoCmd,
		accIndexCmd,
		addrCmd,
		acntCmd,
	},
}

var getterVersionCmd = &cli.Command{
	Name:  "getterVersion",
	Usage: "<Getter> -- Get getter-contract version.",
	Description: `
A function of Getter contract.
Get getter-contract version to identify the contract.
	`,
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := getterIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("getter-contract version:", v)

		return nil
	},
}

var instAddrCmd = &cli.Command{
	Name:  "instAddr",
	Usage: "<Getter> -- Get instance-contract address from getter-contract.",
	Description: `
A function of Getter contract.
Get instance-contract address to verify the consistency of the contract addresses.
	`,
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		inst, err := getterIns.Inst(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("instance-contract address:", inst.Hex())

		return nil
	},
}

var instanceCmd = &cli.Command{
	Name:      "instance",
	Usage:     "<Getter> -- Get instance address by instance type.",
	ArgsUsage: "<instance type>",
	Description: `
A function of Getter contract.
Get instance address by instance type from instance-contract.
Type:
	Proxy
	Control1
	Control2
	Auth
	Access
	ERC20
	PoolP
	Role
	Token
	Pledge
	Issu
	Fs
	Group
	PoolF
	Getter

Arguments:
type - instance type; such as: proxy
	`,
	Action: func(cctx *cli.Context) error {
		ty := cctx.Args().Get(0)
		var insType uint8
		switch ty {
		case "proxy":
			insType = com.TypeProxy
		case "control1":
			insType = com.TypeControl1
		case "control2":
			insType = com.TypeControl2
		case "auth":
			insType = com.TypeAuth
		case "access":
			insType = com.TypeAccess
		case "erc20":
			insType = com.TypeERC20
		case "poolp":
			insType = com.TypePoolP
		case "role":
			insType = com.TypeRole
		case "token":
			insType = com.TypeToken
		case "pledge":
			insType = com.TypePledge
		case "issu":
			insType = com.TypeIssu
		case "fs":
			insType = com.Typefs
		case "group":
			insType = com.TypeGroup
		case "poolf":
			insType = com.TypePoolF
		case "getter":
			insType = com.TypeGetter
		default:
			log.Fatal("please input correct type!")
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		addr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, insType)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("type:", ty, "=>", insType, ", address is:", addr.Hex())

		return nil
	},
}

var getTVersionCmd = &cli.Command{
	Name:  "tVersion",
	Usage: "<Token> -- Get the version of the token-contract.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		taddr, err := getterIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeToken)
		if err != nil {
			log.Fatal(err, ErrRevert)
		}
		tIns, err := token.NewToken(taddr, client)
		if err != nil {
			log.Fatal(err)
		}

		v, err := tIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("token version:", v)

		return nil
	},
}

var getTACmd = &cli.Command{
	Name:      "getTA",
	Usage:     "<Token> -- Get erc20-token address by token index.",
	ArgsUsage: "<ti>",
	Description: `
A function of Getter contract.
Get erc20-token address by token index.

Arguments:
ti - token index
	`,
	Action: func(cctx *cli.Context) error {
		ti, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		taddr, err := getterIns.GetTA(&bind.CallOpts{From: com.AdminAddr}, uint8(ti))
		if err != nil {
			log.Fatal(err, ErrRevert)
		}
		fmt.Println("erc20-token address:", taddr.Hex())

		return nil
	},
}

var getTICmd = &cli.Command{
	Name:      "getTI",
	Usage:     "<Token> -- Get erc20-token index、isBan、mint by erc20-token address.",
	ArgsUsage: "<ta>",
	Description: `
A function of Getter contract.
Get erc20-token index、isBan、mint by erc20-token address.

Arguments:
ta - erc20-token address
	`,
	Action: func(cctx *cli.Context) error {
		taHex := cctx.Args().Get(0)
		if taHex == "" {
			log.Fatal("please input correct erc20-token address")
		}
		ta := common.HexToAddress(taHex)

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		ti, isBan, mint, err := getterIns.GetTI(&bind.CallOpts{From: com.AdminAddr}, ta)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("erc20-token index:", ti, " isBan:", isBan, " mint:", mint)

		return nil
	},
}

var getTCntCmd = &cli.Command{
	Name:  "getTCnt",
	Usage: "<Token> -- Get erc20-token count.",
	Description: `
A function of Getter contract.
Get erc20-token count.
	`,
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		tcnt, err := getterIns.GetTCnt(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("erc20-token count:", tcnt)

		return nil
	},
}

var fsVersionCmd = &cli.Command{
	Name:  "fsVersion",
	Usage: "<Filesys> -- Get fs-contract version.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		fsAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.Typefs)
		if err != nil {
			log.Fatal(err)
		}
		fsIns, err := fs.NewFileSys(fsAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := fsIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("fs-contract version:", v)

		return nil
	},
}

var taxRateCmd = &cli.Command{
	Name:  "taxRate",
	Usage: "<Filesys> -- Get fs-contract taxRate.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		fsAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.Typefs)
		if err != nil {
			log.Fatal(err)
		}
		fsIns, err := fs.NewFileSys(fsAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := fsIns.TaxRate(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("fs-contract taxRate:", v)

		return nil
	},
}

var fsBalCmd = &cli.Command{
	Name:      "fsBal",
	Usage:     "<Filesys> -- Get the balance of the account in fs.",
	ArgsUsage: "<ri ti>",
	Description: `
A function of Getter contract.

Arguments:
ri - account index
ti - token index
	`,
	Action: func(cctx *cli.Context) error {
		ri, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		bal, lock, err := getterIns.FsBalanceOf(&bind.CallOpts{From: com.AdminAddr}, uint64(ri), uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("bal:", com.FormatMemo(bal), " lock:", com.FormatMemo(lock))

		return nil
	},
}

var gsInfoCmd = &cli.Command{
	Name:      "gsInfo",
	Usage:     "<Filesys> -- Get the group information in fs.",
	ArgsUsage: "<gi ti>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
ti - token index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		ginfo, err := getterIns.GetGSInfo(&bind.CallOpts{From: com.AdminAddr}, uint64(gi), uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		gout, _ := json.Marshal(ginfo)
		var out bytes.Buffer
		json.Indent(&out, gout, "", "\t")
		fmt.Printf("ginfo=%v\n", out.String())

		return nil
	},
}

var fsInfoCmd = &cli.Command{
	Name:      "fsInfo",
	Usage:     "<Filesys> -- Get the fs information.",
	ArgsUsage: "<ui pi>",
	Description: `
A function of Getter contract.

Arguments:
ui - user index
pi - provider index
	`,
	Action: func(cctx *cli.Context) error {
		ui, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		pi, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		fsinfo, err := getterIns.GetFsInfo(&bind.CallOpts{From: com.AdminAddr}, uint64(ui), uint64(pi))
		if err != nil {
			log.Fatal(err)
		}
		fsout, _ := json.Marshal(fsinfo)
		var out bytes.Buffer
		json.Indent(&out, fsout, "", "\t")
		fmt.Printf("fsinfo=%v\n", out.String())

		return nil
	},
}

var storeInfoCmd = &cli.Command{
	Name:      "storeInfo",
	Usage:     "<Filesys> -- Get the store information of the user.",
	ArgsUsage: "<ui pi ti>",
	Description: `
A function of Getter contract.

Arguments:
ui - user index
pi - provider index
ti - erc20-token index
	`,
	Action: func(cctx *cli.Context) error {
		ui, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		pi, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(2))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		storeinfo, err := getterIns.GetStoreInfo(&bind.CallOpts{From: com.AdminAddr}, uint64(ui), uint64(pi), uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		storeout, _ := json.Marshal(storeinfo)
		var out bytes.Buffer
		json.Indent(&out, storeout, "", "\t")
		fmt.Printf("storeInfo=%v\n", out.String())

		return nil
	},
}

var settleInfoCmd = &cli.Command{
	Name:      "settleInfo",
	Usage:     "<Filesys> -- Get the settle information of the provider.",
	ArgsUsage: "<pi ti>",
	Description: `
A function of Getter contract.

Arguments:
pi - provider index
ti - erc20-token index
	`,
	Action: func(cctx *cli.Context) error {
		pi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		fsinfo, err := getterIns.GetSettleInfo(&bind.CallOpts{From: com.AdminAddr}, uint64(pi), uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		fsout, _ := json.Marshal(fsinfo)
		var out bytes.Buffer
		json.Indent(&out, fsout, "", "\t")
		fmt.Printf("settleInfo=%v\n", out.String())

		return nil
	},
}

var gVersionCmd = &cli.Command{
	Name:  "gVersion",
	Usage: "<Group> -- Get group-contract version.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGroup)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := group.NewGroup(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group-contract version:", v)

		return nil
	},
}

var gInstCmd = &cli.Command{
	Name:  "gInst",
	Usage: "<Group> -- Get group-contract instance address.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGroup)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := group.NewGroup(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.Inst(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group-contract instance address:", v.Hex())

		return nil
	},
}

var gCntCmd = &cli.Command{
	Name:  "gCnt",
	Usage: "<Group> -- Get the group count.",
	Description: `
A function of Getter contract.
	`,
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		gcnt, err := getterIns.GetGCnt(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("gCnt:", gcnt)
		return nil
	},
}

var gInfoCmd = &cli.Command{
	Name:      "gInfo",
	Usage:     "<Group> -- Get the information of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		isActive, isBan, err := getterIns.GetGInfo(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group isActive:", isActive, " isBan:", isBan)
		return nil
	},
}

var gLevelCmd = &cli.Command{
	Name:      "gLevel",
	Usage:     "<Group> -- Get the level of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		level, err := getterIns.GetLevel(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group level:", level)
		return nil
	},
}

var gPInfoCmd = &cli.Command{
	Name:      "gPInfo",
	Usage:     "<Group> -- Get the pledge information of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		kpledge, ppledge, err := getterIns.GetPInfo(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group keeper pledge:", com.FormatMemo(kpledge), "\t", "provider pledge:", com.FormatMemo(ppledge))
		return nil
	},
}

var kmanageCmd = &cli.Command{
	Name:      "kmanage",
	Usage:     "<Group> -- Get the kmanage-contract address of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		k, err := getterIns.GetKManage(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group kmanage:", k.Hex())
		return nil
	},
}

var gStoreStraCmd = &cli.Command{
	Name:      "gStore",
	Usage:     "<Group> -- Get the storage strategy of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		dc, pc, err := getterIns.GetSStra(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group dc:", dc, "\t", "pc:", pc)
		return nil
	},
}

var issuVersionCmd = &cli.Command{
	Name:  "issuVersion",
	Usage: "<Issuance> -- Get issuance-contract version.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeIssu)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := issuance.NewIssuance(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("issuance-contract version:", v)

		return nil
	},
}

var issuRewardCmd = &cli.Command{
	Name:  "issuReward",
	Usage: "<Issuance> -- Get issuance-contract reward information.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeIssu)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := issuance.NewIssuance(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.RewardResid(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		tr, err := gIns.TargetReward(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("issuance-contract target reward:", com.FormatMemo(tr), "\t", "residual reward:", com.FormatMemo(v))

		return nil
	},
}

var issuStageCmd = &cli.Command{
	Name:  "issuStage",
	Usage: "<Issuance> -- Get issuance-contract stage information.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeIssu)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := issuance.NewIssuance(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		sr, err := gIns.StageRatio(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		ss, err := gIns.StageSize(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		sre, err := gIns.StageReward(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		srr, err := gIns.StageRewardResid(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("issuance-contract stageRatio:", sr, "\t", "stageSize:", ss, "\t", "stageReward:", com.FormatMemo(sre), "\t", "stageRewardResid:", com.FormatMemo(srr))

		return nil
	},
}

var issuStoreCmd = &cli.Command{
	Name:  "issuStore",
	Usage: "<Issuance> -- Get issuance-contract storage information.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeIssu)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := issuance.NewIssuance(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		la, err := gIns.Last(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		si, err := gIns.Size(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		sp, err := gIns.Sprice(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		st, err := gIns.SpaceTime(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("issuance-contract lastIssuTime:", time.Unix(int64(la), 0).Format("2006-01-02 15:04:05"), "\t", "size:", si, "\t", "totalSizePrice:", sp, "\t", "totalSpaceTime:", st)

		return nil
	},
}

var expireCmd = &cli.Command{
	Name:      "expire",
	Usage:     "<Issuance> -- Get issuance-contract expire information.",
	ArgsUsage: "<end time>",
	Description: `
Arguments:
end - end time(align to day), such as 2006-01-02
	`,
	Action: func(cctx *cli.Context) error {
		end := cctx.Args().Get(0)
		t, err := time.Parse("2006-01-02", end)
		fmt.Println("end time duration:", t.Unix(), " end time format:", t.Format("2006-01-02 15:04:05"))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeIssu)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := issuance.NewIssuance(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.Expire(&bind.CallOpts{From: com.AdminAddr}, uint64(t.Unix()))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("issuance-contract end=>sizePrice:", v)

		return nil
	},
}

var kmVersionCmd = &cli.Command{
	Name:      "kmVersion",
	Usage:     "<Kmanage> -- Get kmanage-contract version in the specified group.",
	ArgsUsage: "<gi>",
	Description: `
Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		giStr := cctx.Args().Get(0)
		gi, err := strconv.ParseUint(giStr, 10, 0)
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGroup)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := group.NewGroup(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		km, err := gIns.GetKManage(&bind.CallOpts{From: com.AdminAddr}, gi)
		if err != nil {
			log.Fatal(err)
		}
		kmIns, err := kmanage.NewKmanage(km, client)
		if err != nil {
			log.Fatal(err)
		}
		v, err := kmIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Kmanage-contract version:", v)

		return nil
	},
}

var manageRateCmd = &cli.Command{
	Name:      "mr",
	Usage:     "<Kmanage> -- Get the manageRate of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		mr, err := getterIns.ManageRate(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group manageRate:", mr)
		return nil
	},
}

var kCntCmd = &cli.Command{
	Name:      "kCnt",
	Usage:     "<Kmanage> -- Get the keeper count of the group.",
	ArgsUsage: "<gi>",
	Description: `
A function of Getter contract.

Arguments:
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		gi, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		kcnt, err := getterIns.GetKCnt(&bind.CallOpts{From: com.AdminAddr}, uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group keeper count:", kcnt)
		return nil
	},
}

var getKICmd = &cli.Command{
	Name:      "kI",
	Usage:     "<Kmanage> -- Get the keeper index of the group by array index.",
	ArgsUsage: "<i gi>",
	Description: `
A function of Getter contract.

Arguments:
i - array index
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		i, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		gi, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		ki, err := getterIns.GetK(&bind.CallOpts{From: com.AdminAddr}, uint64(i), uint64(gi))
		if err != nil {
			log.Fatal(err, ErrRevert)
		}
		fmt.Println("group keeper index(ki):", ki)
		return nil
	},
}

var getProfitInfoCmd = &cli.Command{
	Name:      "profit",
	Usage:     "<Kmanage> -- Get the profit information of the group.",
	ArgsUsage: "<ti gi>",
	Description: `
A function of Getter contract.

Arguments:
ti - erc20-token index
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		ti, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		gi, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		last, accu, err := getterIns.GetPf(&bind.CallOpts{From: com.AdminAddr}, uint8(ti), uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("group last share profit:", last, "\t", "profit accumulation:", com.FormatMemo(accu))
		return nil
	},
}

var kmBalCmd = &cli.Command{
	Name:      "kmBal",
	Usage:     "<Kmanage> -- Get the keeper's balance in the group.",
	ArgsUsage: "<ki ti gi>",
	Description: `
A function of Getter contract.

Arguments:
ki - keeper index
ti - erc20-token index
gi - group index
	`,
	Action: func(cctx *cli.Context) error {
		ki, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}
		gi, err := strconv.Atoi(cctx.Args().Get(2))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		avail, tmp, err := getterIns.KmBalanceOf(&bind.CallOpts{From: com.AdminAddr}, uint64(ki), uint8(ti), uint64(gi))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("keeper avail-profit in group:", com.FormatMemo(avail), "\t", "tmp-profit:", com.FormatMemo(tmp))
		return nil
	},
}

var pleVersionCmd = &cli.Command{
	Name:  "pleVersion",
	Usage: "<Pledge> -- Get pledge-contract version.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePledge)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := pledge.NewPledge(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledge-contract version:", v)

		return nil
	},
}

var pleInstCmd = &cli.Command{
	Name:  "pleInst",
	Usage: "<Pledge> -- Get pledge-contract instance address.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePledge)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := pledge.NewPledge(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		v, err := gIns.Inst(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledge-contract instance address:", v.Hex())

		return nil
	},
}

var pleTInfoCmd = &cli.Command{
	Name:      "pleTInfo",
	Usage:     "<Pledge> -- Get pledge-contract token information about rewards.",
	ArgsUsage: "<ti>",
	Description: `
Arguments:
ti - erc20-token index
	`,
	Action: func(cctx *cli.Context) error {
		ti, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePledge)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := pledge.NewPledge(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		v, err := gIns.TInfo(&bind.CallOpts{From: com.AdminAddr}, uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		var out bytes.Buffer
		vout, _ := json.Marshal(v)
		json.Indent(&out, vout, "", "\t")
		fmt.Printf("tInfo=%v\n", out.String())

		return nil
	},
}

var pleRewardsCmd = &cli.Command{
	Name:      "pleRewards",
	Usage:     "<Pledge> -- Get pledge-contract reward information about the account.",
	ArgsUsage: "<ri ti>",
	Description: `
Arguments:
ri - the account index
ti - erc20-token index
	`,
	Action: func(cctx *cli.Context) error {
		ri, err := strconv.ParseUint(cctx.Args().Get(0), 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePledge)
		if err != nil {
			log.Fatal(err)
		}
		gIns, err := pledge.NewPledge(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		v, err := gIns.Rewards(&bind.CallOpts{From: com.AdminAddr}, ri, uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		var out bytes.Buffer
		vout, _ := json.Marshal(v)
		json.Indent(&out, vout, "", "\t")
		fmt.Printf("rewardInfo=%v\n", out.String())

		return nil
	},
}

var pleBalCmd = &cli.Command{
	Name:      "pleBal",
	Usage:     "<Pledge> -- Get the pledge pool balance of the specified account.",
	ArgsUsage: "<ri ti>",
	Description: `
A function of Getter contract.
Include pledge amount and pledge profit.

Arguments:
ri - account index
ti - erc20-token index
	`,
	Action: func(cctx *cli.Context) error {
		ri, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}
		ti, err := strconv.Atoi(cctx.Args().Get(1))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		bal, err := getterIns.PleBalanceOf(&bind.CallOpts{From: com.AdminAddr}, uint64(ri), uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledge balance:", com.FormatMemo(bal))
		return nil
	},
}

var pledgeCmd = &cli.Command{
	Name:      "pledgeAmount",
	Usage:     "<Pledge> -- Get the pledge amount of the erc20-token.",
	ArgsUsage: "<ti>",
	Description: `
A function of Getter contract.

Arguments:
ti - erc20-token index
	`,
	Action: func(cctx *cli.Context) error {
		ti, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		bal, err := getterIns.GetPledge(&bind.CallOpts{From: com.AdminAddr}, uint8(ti))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("pledge amount:", com.FormatMemo(bal))
		return nil
	},
}

var totalPledgeCmd = &cli.Command{
	Name:  "totalPledge",
	Usage: "<Pledge> -- Get the total pledge amount of the primary erc20-token.",
	Description: `
A function of Getter contract.
	`,
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		bal, err := getterIns.GetTotalPledge(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("total pledge amount:", com.FormatMemo(bal))
		return nil
	},
}

var poolVersionCmd = &cli.Command{
	Name:  "poolVersion",
	Usage: "<Pool> -- Get pool-contract version.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		gAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePoolF)
		if err != nil {
			log.Fatal(err)
		}
		ppAddr, _ := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypePoolP)
		gIns, err := pool.NewPool(gAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := gIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		gIns, _ = pool.NewPool(ppAddr, client)
		vp, _ := gIns.Version(&bind.CallOpts{From: com.AdminAddr})
		fmt.Println("poolF-contract version:", v, "\t", "poolP-contract version:", vp)

		return nil
	},
}

var roleVersionCmd = &cli.Command{
	Name:  "roleVersion",
	Usage: "<Role> -- Get role-contract version.",
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		rAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeRole)
		if err != nil {
			log.Fatal(err)
		}
		rIns, err := role.NewRole(rAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		// get version
		v, err := rIns.Version(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("role-contract version:", v)

		return nil
	},
}

var rinfoCmd = &cli.Command{
	Name:      "rinfo",
	Usage:     "<Role> -- Anyone get role info about designated account.",
	ArgsUsage: "<account address>",
	Description: `
A function of Getter contract.
Anyone get role info about designated account.

Arguments:
acc - the designated account
	`,
	Action: func(cctx *cli.Context) error {
		acc := common.HexToAddress(cctx.Args().Get(0))
		fmt.Println("acc:", acc.Hex())

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new roleIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		roleAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeRole)
		if err != nil {
			log.Fatal(err)
		}
		roleIns, _ := role.NewRole(roleAddr, client)

		// get roleInfo
		rinfo, err := roleIns.GetRInfo(&bind.CallOpts{From: com.AdminAddr}, acc)
		if err != nil {
			log.Fatal("get rinfo err:", err)
		}
		roleout, _ := json.Marshal(rinfo)
		var out bytes.Buffer
		json.Indent(&out, roleout, "", "\t")
		fmt.Printf("rinfo=%v\n", out.String())
		return nil
	},
}

var accIndexCmd = &cli.Command{
	Name:      "accIndex",
	Usage:     "<Role> -- Get the index of the account.",
	ArgsUsage: "<addr>",
	Description: `
A function of Getter contract.

Arguments:
addr - the account address
	`,
	Action: func(cctx *cli.Context) error {
		addr := common.HexToAddress(cctx.Args().Get(0))

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		ri, err := getterIns.GetAcc(&bind.CallOpts{From: com.AdminAddr}, addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("account index(ri):", ri)
		return nil
	},
}

var addrCmd = &cli.Command{
	Name:      "addr",
	Usage:     "<Role> -- Get the address of the ri.",
	ArgsUsage: "<ri>",
	Description: `
A function of Getter contract.

Arguments:
ri - the account index
	`,
	Action: func(cctx *cli.Context) error {
		ri, err := strconv.Atoi(cctx.Args().Get(0))
		if err != nil {
			log.Fatal(err)
		}

		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		addr, err := getterIns.GetAddr(&bind.CallOpts{From: com.AdminAddr}, uint64(ri))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("account address:", addr.Hex())
		return nil
	},
}

var acntCmd = &cli.Command{
	Name:  "aCnt",
	Usage: "<Role> -- Get the count of the registered accounts.",
	Description: `
A function of Getter contract.
	`,
	Action: func(cctx *cli.Context) error {
		chain := cctx.String("endPoint")
		instanceAddr, endPoint := com.GetInsEndPointByChain(chain)
		fmt.Println("endPoint:", endPoint, "\t", "instanceAddr:", instanceAddr.Hex())
		fmt.Println()

		// get client
		client, err := ethclient.DialContext(cctx.Context, endPoint)
		if err != nil {
			log.Fatal(err)
		}

		// new getterIns
		instanceIns, _ := inst.NewInstance(instanceAddr, client)
		getterAddr, err := instanceIns.Instances(&bind.CallOpts{From: com.AdminAddr}, com.TypeGetter)
		if err != nil {
			log.Fatal(err)
		}
		getterIns, err := getter.NewGetter(getterAddr, client)
		if err != nil {
			log.Fatal(err)
		}

		acnt, err := getterIns.GetACnt(&bind.CallOpts{From: com.AdminAddr})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("accounts count:", acnt)
		return nil
	},
}
