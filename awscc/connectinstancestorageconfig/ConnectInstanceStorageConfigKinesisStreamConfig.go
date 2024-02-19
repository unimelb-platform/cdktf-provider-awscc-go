package connectinstancestorageconfig


type ConnectInstanceStorageConfigKinesisStreamConfig struct {
	// An ARN is a unique AWS resource identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_instance_storage_config#stream_arn ConnectInstanceStorageConfig#stream_arn}
	StreamArn *string `field:"required" json:"streamArn" yaml:"streamArn"`
}

