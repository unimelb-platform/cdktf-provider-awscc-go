package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_credentials CloudfrontResponseHeadersPolicy#access_control_allow_credentials}.
	AccessControlAllowCredentials interface{} `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_headers CloudfrontResponseHeadersPolicy#access_control_allow_headers}.
	AccessControlAllowHeaders *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowHeaders `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_methods CloudfrontResponseHeadersPolicy#access_control_allow_methods}.
	AccessControlAllowMethods *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowMethods `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_allow_origins CloudfrontResponseHeadersPolicy#access_control_allow_origins}.
	AccessControlAllowOrigins *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlAllowOrigins `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#origin_override CloudfrontResponseHeadersPolicy#origin_override}.
	OriginOverride interface{} `field:"required" json:"originOverride" yaml:"originOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_expose_headers CloudfrontResponseHeadersPolicy#access_control_expose_headers}.
	AccessControlExposeHeaders *CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeaders `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_response_headers_policy#access_control_max_age_sec CloudfrontResponseHeadersPolicy#access_control_max_age_sec}.
	AccessControlMaxAgeSec *float64 `field:"optional" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

