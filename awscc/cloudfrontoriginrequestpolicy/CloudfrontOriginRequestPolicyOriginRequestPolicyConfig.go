package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyOriginRequestPolicyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#cookies_config CloudfrontOriginRequestPolicy#cookies_config}.
	CookiesConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfigCookiesConfig `field:"required" json:"cookiesConfig" yaml:"cookiesConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#headers_config CloudfrontOriginRequestPolicy#headers_config}.
	HeadersConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfigHeadersConfig `field:"required" json:"headersConfig" yaml:"headersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#name CloudfrontOriginRequestPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#query_strings_config CloudfrontOriginRequestPolicy#query_strings_config}.
	QueryStringsConfig *CloudfrontOriginRequestPolicyOriginRequestPolicyConfigQueryStringsConfig `field:"required" json:"queryStringsConfig" yaml:"queryStringsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#comment CloudfrontOriginRequestPolicy#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

