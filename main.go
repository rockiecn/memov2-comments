package main

import (
	"fmt"
	"os"

	"github.com/memoio/contractsv2/cmd"
	"github.com/urfave/cli/v2"
)

// for test
func main() {
	fmt.Println("\nadmin call contracts!")
	fmt.Println()

	commands := []*cli.Command{
		cmd.AdminCmd,
		cmd.GetCmd,
	}

	app := &cli.App{
		Name:                 "contractsv2",
		HelpName:             "used to call memo contract-v2 for admin",
		Usage:                "Tool for admin call memo contracts",
		Version:              "v2",
		EnableBashCompletion: true,
		Commands:             commands,
	}

	app.Setup()

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
