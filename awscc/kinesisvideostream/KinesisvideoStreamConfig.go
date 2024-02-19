package kinesisvideostream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisvideoStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The number of hours till which Kinesis Video will retain the data in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_stream#data_retention_in_hours KinesisvideoStream#data_retention_in_hours}
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
	// The name of the device that is writing to the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_stream#device_name KinesisvideoStream#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_stream#kms_key_id KinesisvideoStream#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_stream#media_type KinesisvideoStream#media_type}
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// The name of the Kinesis Video stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_stream#name KinesisvideoStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs associated with the Kinesis Video Stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_stream#tags KinesisvideoStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

