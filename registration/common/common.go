package common

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Lang map[string]string

const FileModeOct fs.FileMode = 0644

var ElementNameSplitRegex = *regexp.MustCompile("[\\s-_]")

func convertToReadibleName(rawName string) string {
	splitName := ElementNameSplitRegex.Split(rawName, -1)
	processedNameElements := make([]string, len(splitName))
	for i := range splitName {
		processedNameElements[i] = strings.Title(strings.ToLower(splitName[i]))
	}
	return strings.Join(processedNameElements, " ")
}

func CreateLangEntry(kind ElementType, resDir string, modName string, name string) {
	langFile := strings.TrimSuffix(resDir, "\\") + "\\assets\\" + modName + "\\lang\\en_us.json"
	contents := Lang{}
	if _, err := os.Stat(langFile); err == nil {
		b, err := ioutil.ReadFile(langFile)
		CheckError(err)
		err = json.Unmarshal(b, &contents)
		CheckError(err)
	}

	entry := string(kind) + "." + modName + "." + name

	if _, ok := contents[entry]; ok {
		log.Println("[WARN] Lang already has entry for " + name + ", skipping")
		return
	}

	contents[entry] = convertToReadibleName(name)
	b, err := json.Marshal(contents)
	CheckError(err)
	ioutil.WriteFile(
		langFile,
		[]byte(b),
		FileModeOct,
	)
	log.Println("[INFO] (" + string(kind) + "." + modName + "." + name + ") Created lang entry")
}
