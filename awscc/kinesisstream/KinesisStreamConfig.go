package kinesisstream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisStreamConfig struct {
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
	// The name of the Kinesis stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#name KinesisStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#retention_period_hours KinesisStream#retention_period_hours}
	RetentionPeriodHours *float64 `field:"optional" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
	// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#shard_count KinesisStream#shard_count}
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#stream_encryption KinesisStream#stream_encryption}
	StreamEncryption *KinesisStreamStreamEncryption `field:"optional" json:"streamEncryption" yaml:"streamEncryption"`
	// The mode in which the stream is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#stream_mode_details KinesisStream#stream_mode_details}
	StreamModeDetails *KinesisStreamStreamModeDetails `field:"optional" json:"streamModeDetails" yaml:"streamModeDetails"`
	// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesis_stream#tags KinesisStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

