package cloudfrontresponseheaderspolicy


type CloudfrontResponseHeadersPolicyResponseHeadersPolicyConfigCorsConfigAccessControlExposeHeaders struct {
	// The list of HTTP headers. You can specify ``*`` to expose all headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_response_headers_policy#items CloudfrontResponseHeadersPolicy#items}
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

