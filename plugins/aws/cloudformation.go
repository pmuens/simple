package aws

import (
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/pmuens/simple/util"
)

var region = "us-east-1"

func CreateChangeSet(templateBody string, stackName string, changeSetName string) *cloudformation.CreateChangeSetOutput {
	cf := newCloudFormation(region)

	var changeSetType = "CREATE"

	if IsStackReady(stackName) {
		changeSetType = "UPDATE"
	}

	params := &cloudformation.CreateChangeSetInput{
		StackName:     aws.String(stackName),
		ChangeSetName: aws.String(changeSetName),
		ChangeSetType: aws.String(changeSetType),
		Capabilities: []*string{
			aws.String("CAPABILITY_IAM"),
		},
		TemplateBody: aws.String(templateBody),
	}

	resp, err := cf.CreateChangeSet(params)
	if err != nil {
		util.LogPanic(err)
	}

	return resp
}

func DescribeChangeSet(stackName string, changeSetName string) *cloudformation.DescribeChangeSetOutput {
	cf := newCloudFormation(region)
	params := &cloudformation.DescribeChangeSetInput{
		StackName:     aws.String(stackName),
		ChangeSetName: aws.String(changeSetName),
	}

	resp, err := cf.DescribeChangeSet(params)
	if err != nil {
		util.LogPanic(err)
	}

	return resp
}

func ExecuteChangeSet(stackName string, changeSetName string) *cloudformation.ExecuteChangeSetOutput {
	cf := newCloudFormation(region)
	params := &cloudformation.ExecuteChangeSetInput{
		StackName:     aws.String(stackName),
		ChangeSetName: aws.String(changeSetName),
	}

	resp, err := cf.ExecuteChangeSet(params)
	if err != nil {
		re := regexp.MustCompile(`cannot be executed in its current status`)
		if re.MatchString(err.Error()) {
			util.Log("Service was not updated (Don't worry about the FAILED status)...")
			return resp
		}
		util.LogPanic(err)
	}

	return resp
}

func DescribeStacks(stackName string) *cloudformation.DescribeStacksOutput {
	cf := newCloudFormation(region)
	params := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}

	resp, err := cf.DescribeStacks(params)
	if err != nil {
		re := regexp.MustCompile(`does not exist`)
		if re.MatchString(err.Error()) {
			return resp
		}
		util.LogPanic(err)
	}

	return resp
}

func DeleteStack(stackName string) *cloudformation.DeleteStackOutput {
	cf := newCloudFormation(region)
	params := &cloudformation.DeleteStackInput{
		StackName: aws.String(stackName),
	}

	resp, err := cf.DeleteStack(params)
	if err != nil {
		util.LogPanic(err)
	}

	return resp
}

func IsStackReady(stackName string) bool {
	stack := DescribeStacks(stackName)
	if len(stack.Stacks) > 0 {
		stack := stack.Stacks[0]
		re := regexp.MustCompile(`_COMPLETE`)
		if *stack.StackName == stackName && (re.MatchString(*stack.StackStatus)) {
			return true
		}
	}
	return false
}

func newCloudFormation(region string) *cloudformation.CloudFormation {
	sess := CreateSession(region)
	return cloudformation.New(sess)
}
