package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformInfrastructure(t *testing.T) {
	// Load the Terraform configuration
	terraformOptions := &terraform.Options{
		TerraformDir: "path/to/terraform/code",
	}

	// Clean up the infrastructure after the test
	defer terraform.Destroy(t, terraformOptions)

	// Apply the Terraform configuration
	terraform.InitAndApply(t, terraformOptions)

	// Get the ID of the EC2 instance
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Get the name of the S3 bucket
	bucketName := terraform.Output(t, terraformOptions, "bucket_name")

	// Verify that the EC2 instance is running
	awsRegion := aws.String("us-west-2")
	ec2Client := ec2.New(nil, awsRegion)
	instance, err := ec2Client.DescribeInstances(&ec2.DescribeInstancesInput{
		InstanceIds: []*string{&instanceID},
	})
	if err != nil {
		t.Fatalf("Failed to describe EC2 instance: %s", err)
	}
	if *instance.Reservations[0].Instances[0].State.Name != "running" {
		t.Fatalf("Expected EC2 instance to be running but got %s", *instance.Reservations[0].Instances[0].State.Name)
	}

	// Verify that the S3 bucket exists
	s3Client := s3.New(nil, awsRegion)
	_, err = s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		t.Fatalf("Failed to head S3 bucket: %s", err)
	}
}
