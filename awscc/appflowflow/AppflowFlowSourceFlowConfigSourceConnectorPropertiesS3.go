package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#bucket_name AppflowFlow#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#bucket_prefix AppflowFlow#bucket_prefix}.
	BucketPrefix *string `field:"required" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#s3_input_format_config AppflowFlow#s3_input_format_config}.
	S3InputFormatConfig *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3S3InputFormatConfig `field:"optional" json:"s3InputFormatConfig" yaml:"s3InputFormatConfig"`
}

