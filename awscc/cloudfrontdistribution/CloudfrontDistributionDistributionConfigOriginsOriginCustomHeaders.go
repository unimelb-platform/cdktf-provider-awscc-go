package cloudfrontdistribution


type CloudfrontDistributionDistributionConfigOriginsOriginCustomHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#header_name CloudfrontDistribution#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_distribution#header_value CloudfrontDistribution#header_value}.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
}

