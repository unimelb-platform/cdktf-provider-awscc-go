package cloudfrontfunction


type CloudfrontFunctionFunctionConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#comment CloudfrontFunction#comment}.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#runtime CloudfrontFunction#runtime}.
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_function#key_value_store_associations CloudfrontFunction#key_value_store_associations}.
	KeyValueStoreAssociations interface{} `field:"optional" json:"keyValueStoreAssociations" yaml:"keyValueStoreAssociations"`
}

