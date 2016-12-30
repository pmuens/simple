package plugin

import (
	"regexp"
	"time"

	"github.com/pmuens/simple/plugins/aws"
	"github.com/pmuens/simple/util"
)

var stackName = util.GetServiceName()
var bucketName = util.GetServiceBucketName()

func EmptyBucket() {
	util.Log("Removing artifacts...")

	aws.EmptyBucket(bucketName)

	util.Log("Successfully removed artifacts...")
}

func DeleteStack() {
	util.Log("Removing Stack...")

	aws.DeleteStack(stackName)

	duration := time.Duration(5) * time.Second
	for range time.Tick(duration) {
		resp := aws.DescribeStacks(stackName)
		if len(resp.Stacks) > 0 {
			stack := resp.Stacks[0]
			currentStatus := string(*stack.StackStatus)
			util.Log(currentStatus)
			re := regexp.MustCompile(`_COMPLETE|FAILED`)
			if re.MatchString(currentStatus) {
				break
			}
		} else {
			break
		}
	}

	util.Log("Done...")
}
