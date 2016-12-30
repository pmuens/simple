package plugin

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Unknwon/cae/zip"

	"github.com/pmuens/simple/util"
)

func ZipService() {
	serviceDir := util.GetServiceDir()
	simpleDir := util.GetSimpleDir()
	zipName := util.GetZipName()

	zipFile := filepath.Join(simpleDir, zipName)

	if _, err := os.Stat(zipFile); err == nil {
		os.Remove(zipFile)
	}

	util.Log("Creating .zip file...")
	zip.PackToFunc(serviceDir, zipFile, reportZipProgress, false)
	os.Chmod(zipFile, os.FileMode(0755))
}

func reportZipProgress(path string, fileInfo os.FileInfo) error {
	if !fileInfo.IsDir() {
		message := fmt.Sprintf("Compressing %s", path)
		util.Log(message)
	}
	return nil
}
