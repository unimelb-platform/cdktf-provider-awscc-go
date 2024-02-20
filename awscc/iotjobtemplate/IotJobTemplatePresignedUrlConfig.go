package iotjobtemplate


type IotJobTemplatePresignedUrlConfig struct {
	// The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored.
	//
	// The role must also grant permission for IoT to download the files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#role_arn IotJobTemplate#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// How number (in seconds) pre-signed URLs are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#expires_in_sec IotJobTemplate#expires_in_sec}
	ExpiresInSec *float64 `field:"optional" json:"expiresInSec" yaml:"expiresInSec"`
}

