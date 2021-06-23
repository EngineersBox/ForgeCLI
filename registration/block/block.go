package block

import (
	"engineersbox/forgecli/registration/common"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

const FileModeOct fs.FileMode = 0644

func CreateBlockState(resDir string, modName string, blockName string) {
	contents := strings.Replace(BlockState, "<0>", modName, 1)
	contents = strings.Replace(contents, "<1>", blockName, 1)
	ioutil.WriteFile(
		strings.TrimSuffix(resDir, "\\")+"\\assets\\"+modName+"\\blockstates\\"+blockName+".json",
		[]byte(contents),
		FileModeOct,
	)
	log.Println("[INFO] (block." + modName + "." + blockName + ") Created blockstate")
}

func CreateBlockModel(resDir string, modName string, blockName string) {
	contents := strings.Replace(BlockModel, "<0>", modName, 1)
	contents = strings.Replace(contents, "<1>", blockName, 1)
	ioutil.WriteFile(
		strings.TrimSuffix(resDir, "\\")+"\\assets\\"+modName+"\\models\\block\\"+blockName+".json",
		[]byte(contents),
		FileModeOct,
	)
	log.Println("[INFO] (block." + modName + "." + blockName + ") Created block model")
}

func CreateBlockItemModel(resDir string, modName string, blockName string) {
	contents := strings.Replace(BlockItemModel, "<0>", modName, 1)
	contents = strings.Replace(contents, "<1>", blockName, 1)
	ioutil.WriteFile(
		strings.TrimSuffix(resDir, "\\")+"\\assets\\"+modName+"\\models\\item\\"+blockName+".json",
		[]byte(contents),
		FileModeOct,
	)
	log.Println("[INFO] (block." + modName + "." + blockName + ") Created block item")
}

func CreateRegistryObject(modName string, blockName string, material string, registryDir string) {
	b, err := ioutil.ReadFile(registryDir)
	common.CheckError(err)
	contents := string(b)

	constBlockName := strings.ReplaceAll(strings.ToUpper(blockName), "-", "_")
	blockID := strings.ToLower(constBlockName)

	if strings.Contains(contents, constBlockName) {
		log.Println("[WARN] Registry already has entry for " + blockID + ", skipping")
		return
	}

	contents = strings.TrimSuffix(strings.TrimSpace(contents), "}")
	registryEntry := strings.Replace(RegistryObject, "<0>", constBlockName, 1)
	registryEntry = strings.Replace(registryEntry, "<1>", blockID, 1)
	registryEntry = strings.Replace(registryEntry, "<2>", strings.ToUpper(material), 1)
	contents += "\t" + registryEntry + "\n}\n"
	ioutil.WriteFile(
		registryDir,
		[]byte(contents),
		FileModeOct,
	)
	log.Println("[INFO] (block." + modName + "." + blockID + ") Added block registry entry")
}
