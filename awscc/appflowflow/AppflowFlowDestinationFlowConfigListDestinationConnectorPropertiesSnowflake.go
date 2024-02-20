package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflake struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#intermediate_bucket_name AppflowFlow#intermediate_bucket_name}.
	IntermediateBucketName *string `field:"required" json:"intermediateBucketName" yaml:"intermediateBucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#object AppflowFlow#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#bucket_prefix AppflowFlow#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#error_handling_config AppflowFlow#error_handling_config}.
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesSnowflakeErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

