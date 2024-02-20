package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyOriginRequestPolicyConfigCookiesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#cookie_behavior CloudfrontOriginRequestPolicy#cookie_behavior}.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#cookies CloudfrontOriginRequestPolicy#cookies}.
	Cookies *[]*string `field:"optional" json:"cookies" yaml:"cookies"`
}

