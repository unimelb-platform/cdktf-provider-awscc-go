package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#name CloudfrontResponseHeadersPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#comment CloudfrontResponseHeadersPolicy#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#cors_config CloudfrontResponseHeadersPolicy#cors_config}.
	CorsConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfig `field:"optional" json:"corsConfig" yaml:"corsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#custom_headers_config CloudfrontResponseHeadersPolicy#custom_headers_config}.
	CustomHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCustomHeadersConfig `field:"optional" json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#remove_headers_config CloudfrontResponseHeadersPolicy#remove_headers_config}.
	RemoveHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigRemoveHeadersConfig `field:"optional" json:"removeHeadersConfig" yaml:"removeHeadersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#security_headers_config CloudfrontResponseHeadersPolicy#security_headers_config}.
	SecurityHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfig `field:"optional" json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#server_timing_headers_config CloudfrontResponseHeadersPolicy#server_timing_headers_config}.
	ServerTimingHeadersConfig *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigServerTimingHeadersConfig `field:"optional" json:"serverTimingHeadersConfig" yaml:"serverTimingHeadersConfig"`
}

