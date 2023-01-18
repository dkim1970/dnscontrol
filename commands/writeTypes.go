package commands

import (
	_ "embed"
	"os"

	versionInfo "github.com/StackExchange/dnscontrol/v3/pkg/version"
	"github.com/urfave/cli/v2"
)

var _ = cmd(catUtils, func() *cli.Command {
	var args TypesArgs
	return &cli.Command{
		Name:  "write-types",
		Usage: "[BETA] Write a TypeScript declaration file in the current directory",
		Action: func(c *cli.Context) error {
			return exit(WriteTypes(args))
		},
		Flags: args.flags(),
	}
}())

// TypesArgs stores arguments related to the types subcommand.
type TypesArgs struct {
	DTSFile string
}

func (args *TypesArgs) flags() []cli.Flag {
	var flags []cli.Flag
	flags = append(flags, &cli.StringFlag{
		Name:        "dts-file",
		Aliases:     []string{"o"},
		Value:       "types-dnscontrol.d.ts",
		Usage:       "Path to the .d.ts file to create",
		Destination: &args.DTSFile,
	})
	return flags
}

//go:embed types/dnscontrol.d.ts
var dtsContent string

func WriteTypes(args TypesArgs) error {
	file, err := os.Create(args.DTSFile)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("// This file was automatically generated by DNSControl. Do not edit it directly.\n")
	file.WriteString("// To update it, run `dnscontrol write-types`.\n\n")
	file.WriteString("// DNSControl version: " + versionInfo.Banner() + "\n")
	file.WriteString(dtsContent)
	if err != nil {
		return err
	}

	print("Successfully wrote " + args.DTSFile + "\n")
	return nil
}