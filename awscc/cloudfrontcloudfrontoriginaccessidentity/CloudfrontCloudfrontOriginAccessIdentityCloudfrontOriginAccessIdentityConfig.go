package cloudfrontcloudfrontoriginaccessidentity


type CloudfrontCloudfrontOriginAccessIdentityCloudfrontOriginAccessIdentityConfig struct {
	// A comment to describe the origin access identity. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cloudfront_cloudfront_origin_access_identity#comment CloudfrontCloudfrontOriginAccessIdentity#comment}
	Comment *string `field:"required" json:"comment" yaml:"comment"`
}

