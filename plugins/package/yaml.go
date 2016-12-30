package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pmuens/simple/util"
)

func CreateYAMLFile(name string, cont []byte) {
	logMessage := fmt.Sprintf("Creating %s file...", name)
	util.Log(logMessage)

	simpleDir := util.GetSimpleDir()

	os.MkdirAll(simpleDir, os.ModePerm)

	simpleFilePath := filepath.Join(simpleDir, name)
	ioutil.WriteFile(simpleFilePath, cont, 0755)
}
