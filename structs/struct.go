package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type CloudWatchEvent struct {
	Version    string            `json:"version"`
	ID         string            `json:"id"`
	DetailType string            `json:"detail-type"`
	Source     string            `json:"source"`
	AccountID  string            `json:"account"`
	Time       time.Time         `json:"time"`
	Region     string            `json:"region"`
	Resources  []string          `json:"resources"`
	Detail     CloudWatchDetails `json:"detail"`
}

type CloudWatchDetails struct {
	StatusCode           string `json:"StatusCode"`
	AutoScalingGroupName string `json:"AutoScalingGroupName"`
	ActivityID           string `json:"ActivityId"`
	Details              struct {
		AvailabilityZone string `json:"Availability Zone"`
		SubnetID         string `json:"Subnet ID"`
	} `json:"Details"`
	RequestID     string    `json:"RequestId"`
	EndTime       time.Time `json:"EndTime"`
	EC2InstanceID string    `json:"EC2InstanceId"`
	StartTime     time.Time `json:"StartTime"`
	Cause         string    `json:"Cause"`
}

func main() {
	payload := []byte(`{
  			"version": "0",
  			"id": "3e3c153a-8339-4e30-8c35-687ebef853fe",
  			"detail-type": "EC2 Instance Launch Successful",
  			"source": "aws.autoscaling",
  			"account": "123456789012",
  			"time": "2015-11-11T21:31:47Z",
  			"region": "us-east-1",
  			"resources": [
  			  "arn:aws:autoscaling:us-east-1:123456789012:autoScalingGroup:eb56d16b-bbf0-401d-b893-d5978ed4a025:autoScalingGroupName/sampleLuanchSucASG",
  			  "arn:aws:ec2:us-east-1:123456789012:instance/i-b188560f"
  			],
  			"detail": {
  			  "StatusCode": "InProgress",
  			  "AutoScalingGroupName": "sampleLuanchSucASG",
  			  "ActivityId": "9cabb81f-42de-417d-8aa7-ce16bf026590",
  			  "Details": {
  			    "Availability Zone": "us-east-1b",
  			    "Subnet ID": "subnet-95bfcebe"
  			  },
  			  "RequestId": "9cabb81f-42de-417d-8aa7-ce16bf026590",
  			  "EndTime": "2015-11-11T21:31:47.208Z",
  			  "EC2InstanceId": "i-b188560f",
  			  "StartTime": "2015-11-11T21:31:13.671Z",
  			  "Cause": "At 2015-11-11T21:31:10Z a user request created an AutoScalingGroup changing the desired capacity from 0 to 1.  At 2015-11-11T21:31:11Z an instance was started in response to a difference between desired and actual capacity, increasing the capacity from 0 to 1."
  			}
			}`)	

	object := &CloudWatchEvent{}

	err := json.Unmarshal(payload, object)

	if err != nil {
		panic(err)
	}

	fmt.Println(object.Detail.Details.AvailabilityZone)

}