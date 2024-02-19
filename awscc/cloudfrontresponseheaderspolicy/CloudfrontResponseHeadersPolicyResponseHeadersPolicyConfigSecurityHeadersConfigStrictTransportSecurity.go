package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigSecurityHeadersConfigStrictTransportSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec CloudfrontResponseHeadersPolicy#access_control_max_age_sec}.
	AccessControlMaxAgeSec *float64 `field:"required" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#override CloudfrontResponseHeadersPolicy#override}.
	Override interface{} `field:"required" json:"override" yaml:"override"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#include_subdomains CloudfrontResponseHeadersPolicy#include_subdomains}.
	IncludeSubdomains interface{} `field:"optional" json:"includeSubdomains" yaml:"includeSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#preload CloudfrontResponseHeadersPolicy#preload}.
	Preload interface{} `field:"optional" json:"preload" yaml:"preload"`
}

