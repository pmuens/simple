package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Log(message string) {
	l := log.New(os.Stdout, "[SIMPLE]: ", 0)
	l.Printf(message)
}

func LogError(err error) {
	l := log.New(os.Stdout, "[SIMPLE - Error]: ", 1)
	l.Fatalf(err.Error())
}

func LogPanic(err error) {
	l := log.New(os.Stdout, "[SIMPLE - Panic]: ", 1)
	l.Panicf(err.Error())
}

func GetServiceDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		LogPanic(err)
	}
	return cwd
}

func GetSimpleDir() string {
	return filepath.Join(GetServiceDir(), ".simple")
}

func GetServiceName() string {
	splittedServiceDir := strings.Split(GetServiceDir(), string(os.PathSeparator))
	return splittedServiceDir[len(splittedServiceDir)-1]
}

func GetServiceBucketName() string {
	return fmt.Sprintf("%s-bucket", GetServiceName())
}

func GetZipName() string {
	return fmt.Sprintf("%s.zip", GetServiceName())
}

func TitleizeAndConcat(name string) string {
	var r = regexp.MustCompile("[0-9A-Za-z]+")

	byteSrc := []byte(name)
	chunks := r.FindAll(byteSrc, -1)

	for i, _ := range chunks {
		chunks[i] = bytes.ToLower(chunks[i])
		chunks[i] = bytes.Title(chunks[i])
	}

	return string(bytes.Join(chunks, nil))
}
