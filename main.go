package main

import (
	"engineersbox/forgecli/config"
	"engineersbox/forgecli/registration/block"
	"engineersbox/forgecli/registration/common"
	"flag"
	"log"
	"os"

	"github.com/EngineersBox/ModularCLI/cli"
)

var commands = map[string]cli.SubCommand{
	"block": {
		ErrorHandler: flag.ExitOnError,
		Flags: []*cli.Flag{
			{
				Type:         cli.TypeString,
				Name:         "name",
				DefaultValue: "block",
				HelpMsg:      "Name of the block",
				Required:     true,
			},
			{
				Type:         cli.TypeString,
				Name:         "mat",
				DefaultValue: "./",
				HelpMsg:      "Block material type, this should be one of the net.minecraft.block.material.Material enum values",
				Required:     true,
			},
		},
	},
	"item": {
		ErrorHandler: flag.ExitOnError,
		Flags: []*cli.Flag{
			{
				Type:         cli.TypeString,
				Name:         "name",
				DefaultValue: "block",
				HelpMsg:      "Name of the item",
				Required:     true,
			},
		},
	},
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func blockCommandHandler(blockCmd *cli.Command, cfg config.ForgeCLIConfig) {
	block.CreateBlockState(
		cfg.ResourcesDir,
		cfg.ModName,
		*blockCmd.Flags["name"].GetString(),
	)
	block.CreateBlockModel(
		cfg.ResourcesDir,
		cfg.ModName,
		*blockCmd.Flags["name"].GetString(),
	)
	block.CreateBlockItemModel(
		cfg.ResourcesDir,
		cfg.ModName,
		*blockCmd.Flags["name"].GetString(),
	)
	common.CreateLangEntry(
		common.Block,
		cfg.ResourcesDir,
		cfg.ModName,
		*blockCmd.Flags["name"].GetString(),
	)
	block.CreateRegistryObject(
		cfg.ModName,
		*blockCmd.Flags["name"].GetString(),
		*blockCmd.Flags["mat"].GetString(),
		cfg.RegFiles.Block,
	)
}

func itemCommandHandler(itemCmd *cli.Command, cfg config.ForgeCLIConfig) {
	return
}

func main() {
	config := config.ReadConfig()
	forgeCLI, err := cli.CreateCLI(commands)
	common.CheckError(err)

	err = forgeCLI.Parse()
	common.CheckError(err)

	cmd := os.Args[1]
	if cmd == "block" {
		blockCommandHandler(forgeCLI.Commands["block"], config)
	} else if cmd == "item" {
		itemCommandHandler(forgeCLI.Commands["item"], config)
	}
}
