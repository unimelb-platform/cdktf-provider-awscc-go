package qbusinessplugin


type QbusinessPluginCustomPluginConfigurationApiSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_plugin#payload QbusinessPlugin#payload}.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_plugin#s3 QbusinessPlugin#s3}.
	S3 *QbusinessPluginCustomPluginConfigurationApiSchemaS3 `field:"optional" json:"s3" yaml:"s3"`
}

