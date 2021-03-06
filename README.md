# ForgeCLI

## DISCLAIMER

This tool is in no way assosciated with Forge or any content made by the developers of Forge. This is purely a third party tool by EngineersBox used for Forge mod development.

## Build

To build the CLI simply clone down down the repo and run

```shell
go build
```

## Using the tool

Once you have built it, you'll need to create your configuration file. An example one has been provided called `forgeCliConfig.json` alonside an example filesystem `testFS`.

If you are creating your own, the config file must be named `forgeCliConfig.json` and be present in the same directory as the CLI tool.

To run commands with the tool simply use the following (assuming you are in the same directory as the buld arefact):

Mac OS X and Linux
```shell
./forgecli <COMMAND> [...<OPTIONS>]
```

Windows
```shell
.\forgecli.exe <COMMAND> [...<OPTIONS>]
```

### Config

Below is the schema for the configuration file. Ensure that the config is named `forgeCliConfig.json` and is present in the same directory as the build executable.

```json
{
    "mod_name": "<string>",
    "resources_dir": "<directory>",
    "registry_files": {
        "block": "<.java file location>",
        "item": "<.java file location>",
        "tile_entity": "<.java file location>",
        "fluid": "<.java file location>"
    }
}
```

## Commands

```shell
$> forgecli <COMMAND> [...<OPTIONS>]

Commands:
  block         Create a new block registration
        --name  The name of the block\n
        --mat   Block material type, this must be one of the
                net.minecraft.block.material.Material enum values

  item          Create a new item registration
        --name  The name of the item

  fluid         Create a new fluid registration
        --name  The name of the fluid

  tilentity     Create a new tile entity registration
        --name  The name of the tilentity

  help          Displays this message
```