package s3bucket


type S3BucketWebsiteConfigurationRoutingRulesRedirectRule struct {
	// The host name to use in the redirect request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#host_name S3Bucket#host_name}
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The HTTP redirect code to use on the response. Not required if one of the siblings is present.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#http_redirect_code S3Bucket#http_redirect_code}
	HttpRedirectCode *string `field:"optional" json:"httpRedirectCode" yaml:"httpRedirectCode"`
	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#protocol S3Bucket#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The object key prefix to use in the redirect request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#replace_key_prefix_with S3Bucket#replace_key_prefix_with}
	ReplaceKeyPrefixWith *string `field:"optional" json:"replaceKeyPrefixWith" yaml:"replaceKeyPrefixWith"`
	// The specific object key to use in the redirect request.d.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#replace_key_with S3Bucket#replace_key_with}
	ReplaceKeyWith *string `field:"optional" json:"replaceKeyWith" yaml:"replaceKeyWith"`
}

