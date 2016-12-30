package plugin

import (
	"io/ioutil"
	"regexp"

	"github.com/pmuens/simple/util"
)

func GetFunctionDirectoryNames(serviceDir string) []string {
	cont, err := ioutil.ReadDir(serviceDir)
	if err != nil {
		util.LogPanic(err)
	}

	var dirNames []string

	for _, c := range cont {
		dirNames = append(dirNames, c.Name())
	}

	return extractFunctionDirectories(dirNames)
}

func extractFunctionDirectories(directoryNames []string) []string {
	re := regexp.MustCompile(`(none|s3|api|timer|alexa-skill)+`) // TODO only when at the beginning

	var funcs []string

	for _, dir := range directoryNames {
		m := re.MatchString(dir)
		if m {
			funcs = append(funcs, dir)
		}
	}

	return funcs
}
