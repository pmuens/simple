package plugin

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/pmuens/simple/plugins/aws"
	"github.com/pmuens/simple/util"
)

var simpleDir = util.GetSimpleDir()
var stackName = util.GetServiceName()
var bucketName = util.GetServiceBucketName()
var zipName = util.GetZipName()
var changeSetName = fmt.Sprintf("%s-%s", stackName, strconv.FormatInt(time.Now().UTC().UnixNano(), 10))

func CreateStack() {
	if aws.IsStackReady(stackName) {
		return
	}

	yamlContent, err := ioutil.ReadFile(filepath.Join(simpleDir, "create-stack.yml"))
	if err != nil {
		util.LogPanic(err)
	}

	util.Log("Creating Changeset...")
	aws.CreateChangeSet(string(yamlContent), stackName, changeSetName)

	var duration = time.Duration(5) * time.Second
	var currentStatus = ""
	re := regexp.MustCompile(`_COMPLETE|FAILED`)

	for range time.Tick(duration) {
		resp := aws.DescribeChangeSet(stackName, changeSetName)
		currentStatus = string(*resp.Status)
		if currentStatus != "FAILED" { // stack was not updated / new changeset was not created
			util.Log(currentStatus)
		}
		if re.MatchString(currentStatus) {
			break
		}
	}

	util.Log("Creating Stack...")
	aws.ExecuteChangeSet(stackName, changeSetName)

	for range time.Tick(duration) {
		resp := aws.DescribeStacks(stackName)
		stack := resp.Stacks[0]
		currentStatus = string(*stack.StackStatus)
		util.Log(currentStatus)
		if re.MatchString(currentStatus) {
			break
		}
	}

	util.Log("Stack successfully created...")
}

func UpdateStack() {
	util.Log("Uploading artifacts...")

	zip, err := ioutil.ReadFile(filepath.Join(simpleDir, zipName))
	if err != nil {
		util.LogPanic(err)
	}

	aws.PutObject(bucketName, zipName, zip)

	util.Log("Creating Changeset...")
	yamlFileName := "update-stack.yml"
	yamlContent, err := ioutil.ReadFile(filepath.Join(simpleDir, yamlFileName))
	if err != nil {
		util.LogPanic(err)
	}
	aws.CreateChangeSet(string(yamlContent), stackName, changeSetName)

	var duration = time.Duration(5) * time.Second
	var currentStatus = ""
	re := regexp.MustCompile(`_COMPLETE|FAILED`)

	for range time.Tick(duration) {
		resp := aws.DescribeChangeSet(stackName, changeSetName)
		currentStatus = string(*resp.Status)
		util.Log(currentStatus)
		if re.MatchString(currentStatus) {
			break
		}
	}

	util.Log("Updating Stack...")
	aws.ExecuteChangeSet(stackName, changeSetName)

	for range time.Tick(duration) {
		resp := aws.DescribeStacks(stackName)
		stack := resp.Stacks[0]
		currentStatus = string(*stack.StackStatus)
		util.Log(currentStatus)
		if re.MatchString(currentStatus) {
			break
		}
	}
}
