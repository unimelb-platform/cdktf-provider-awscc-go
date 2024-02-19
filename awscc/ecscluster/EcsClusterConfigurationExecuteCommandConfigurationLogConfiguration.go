package ecscluster


type EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#cloudwatch_encryption_enabled EcsCluster#cloudwatch_encryption_enabled}.
	CloudwatchEncryptionEnabled interface{} `field:"optional" json:"cloudwatchEncryptionEnabled" yaml:"cloudwatchEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#cloudwatch_log_group_name EcsCluster#cloudwatch_log_group_name}.
	CloudwatchLogGroupName *string `field:"optional" json:"cloudwatchLogGroupName" yaml:"cloudwatchLogGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#s3_bucket_name EcsCluster#s3_bucket_name}.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#s3_encryption_enabled EcsCluster#s3_encryption_enabled}.
	S3EncryptionEnabled interface{} `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#s3_key_prefix EcsCluster#s3_key_prefix}.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

