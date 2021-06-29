package block

import (
	"engineersbox/forgecli/logging/log"
	"engineersbox/forgecli/registration/common"
	template "engineersbox/forgecli/templates"
	"io/fs"
	"io/ioutil"
	"strings"
)

const FileModeOct fs.FileMode = 0644

func CreateBlockState(resDir string, modName string, blockName string) {
	blockState := template.ImportTemplate(template.SingleModelBlockState)
	formatters := map[string]string{
		"<0>": modName,
		"<1>": blockName,
	}
	contents := template.ReplaceInlineFormatters(blockState, formatters)
	ioutil.WriteFile(
		strings.TrimSuffix(resDir, "/")+"/assets/"+modName+"/blockstates/"+blockName+".json",
		[]byte(contents),
		FileModeOct,
	)
	log.Info("(block." + modName + "." + blockName + ") Created blockstate")
}

func CreateBlockModel(resDir string, modName string, blockName string) {
	blockModel := template.ImportTemplate(template.FullBlockSingleTextureModel)
	formatters := map[string]string{
		"<0>": modName,
		"<1>": blockName,
	}
	contents := template.ReplaceInlineFormatters(blockModel, formatters)
	ioutil.WriteFile(
		strings.TrimSuffix(resDir, "/")+"/assets/"+modName+"/models/block/"+blockName+".json",
		[]byte(contents),
		FileModeOct,
	)
	log.Info("(block." + modName + "." + blockName + ") Created block model")
}

func CreateBlockItemModel(resDir string, modName string, blockName string) {
	blockItemModel := template.ImportTemplate(template.BlockItemParent)
	formatters := map[string]string{
		"<0>": modName,
		"<1>": blockName,
	}
	contents := template.ReplaceInlineFormatters(blockItemModel, formatters)
	ioutil.WriteFile(
		strings.TrimSuffix(resDir, "/")+"/assets/"+modName+"/models/item/"+blockName+".json",
		[]byte(contents),
		FileModeOct,
	)
	log.Info("(block." + modName + "." + blockName + ") Created block item")
}

func CreateRegistryObject(modName string, blockName string, material string, registryDir string) {
	b, err := ioutil.ReadFile(registryDir)
	common.CheckError(err)
	contents := string(b)

	constBlockName := strings.ReplaceAll(strings.ToUpper(blockName), "-", "_")
	blockID := strings.ToLower(constBlockName)

	if strings.Contains(contents, constBlockName) {
		log.Warn("Registry already has entry for " + blockID + ", skipping")
		return
	}

	contents = strings.TrimSuffix(strings.TrimSpace(contents), "}")

	registryTemplate := template.ImportTemplate(template.RegistryBasicBlockWithMaterial)
	formatters := map[string]string{
		"<0>": constBlockName,
		"<1>": blockID,
		"<2>": strings.ToUpper(material),
	}
	registryEntry := template.ReplaceInlineFormatters(registryTemplate, formatters)

	contents += "\t" + registryEntry + "\n}\n"
	ioutil.WriteFile(
		registryDir,
		[]byte(contents),
		FileModeOct,
	)
	log.Info("(block." + modName + "." + blockID + ") Added block registry entry")
}
