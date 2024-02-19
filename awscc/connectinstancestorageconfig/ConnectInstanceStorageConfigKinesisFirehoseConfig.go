package connectinstancestorageconfig


type ConnectInstanceStorageConfigKinesisFirehoseConfig struct {
	// An ARN is a unique AWS resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#firehose_arn ConnectInstanceStorageConfig#firehose_arn}
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
}

