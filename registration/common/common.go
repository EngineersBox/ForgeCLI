package common

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Lang map[string]string

const FileModeOct fs.FileMode = 0644

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

	contents[entry] = name
	b, err := json.Marshal(contents)
	CheckError(err)
	ioutil.WriteFile(
		langFile,
		[]byte(b),
		FileModeOct,
	)
	log.Println("[INFO] (" + string(kind) + "." + modName + "." + name + ") Created lang entry")
}
