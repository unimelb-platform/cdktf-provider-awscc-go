package s3bucket


type S3BucketWebsiteConfigurationRoutingRulesRoutingRuleCondition struct {
	// The HTTP error code when the redirect is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#http_error_code_returned_equals S3Bucket#http_error_code_returned_equals}
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// The object key name prefix when the redirect is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#key_prefix_equals S3Bucket#key_prefix_equals}
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}

