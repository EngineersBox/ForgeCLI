package main

import (
	"engineersbox/forgecli/config"
	"engineersbox/forgecli/logging"
	"engineersbox/forgecli/registration/block"
	"engineersbox/forgecli/registration/common"
	"flag"
	"fmt"
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
	"help": {
		ErrorHandler: flag.ExitOnError,
	},
}

func check(e error) {
	if e != nil {
		logging.Fatal(e)
	}
}

type HandlerFunc func(*cli.Command, config.ForgeCLIConfig)

var handlerMapper = map[string]HandlerFunc{
	"block": func(blockCmd *cli.Command, cfg config.ForgeCLIConfig) {
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
	},
	"item": func(itemCmd *cli.Command, cfg config.ForgeCLIConfig) {
		return
	},
	"help": func(itemCmd *cli.Command, cfg config.ForgeCLIConfig) {
		fmt.Printf(`
$> forgecli <COMMAND> [...<OPTIONS>]

Commands:
  block		Create a new block registration
	--name	The name of the block\n
	--mat	Block material type, this must be one of the
		net.minecraft.block.material.Material enum values

  item		Create a new item registration
	--name	The name of the item

  fluid		Create a new fluid registration
	--name	The name of the fluid

  tilentity	Create a new tile entity registration
	--name	The name of the tilentity

  help		Displays this message
		`)
	},
}

func main() {
	config := config.ReadConfig()
	forgeCLI, err := cli.CreateCLI(commands)
	common.CheckError(err)

	err = forgeCLI.Parse()
	common.CheckError(err)

	cmd := os.Args[1]
	handler, ok := handlerMapper[cmd]
	if !ok {
		logging.Fatal(" Unknown command: %s", cmd)
	}

	handler(forgeCLI.Commands[cmd], config)
}
