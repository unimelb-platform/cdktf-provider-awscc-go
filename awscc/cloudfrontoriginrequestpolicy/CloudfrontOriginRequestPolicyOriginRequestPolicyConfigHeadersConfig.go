package cloudfrontoriginrequestpolicy


type CloudfrontOriginRequestPolicyOriginRequestPolicyConfigHeadersConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#header_behavior CloudfrontOriginRequestPolicy#header_behavior}.
	HeaderBehavior *string `field:"required" json:"headerBehavior" yaml:"headerBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_request_policy#headers CloudfrontOriginRequestPolicy#headers}.
	Headers *[]*string `field:"optional" json:"headers" yaml:"headers"`
}

