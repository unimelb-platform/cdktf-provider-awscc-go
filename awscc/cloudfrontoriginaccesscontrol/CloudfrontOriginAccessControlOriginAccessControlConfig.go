package cloudfrontoriginaccesscontrol


type CloudfrontOriginAccessControlOriginAccessControlConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_access_control#name CloudfrontOriginAccessControl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_access_control#origin_access_control_origin_type CloudfrontOriginAccessControl#origin_access_control_origin_type}.
	OriginAccessControlOriginType *string `field:"required" json:"originAccessControlOriginType" yaml:"originAccessControlOriginType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_access_control#signing_behavior CloudfrontOriginAccessControl#signing_behavior}.
	SigningBehavior *string `field:"required" json:"signingBehavior" yaml:"signingBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_access_control#signing_protocol CloudfrontOriginAccessControl#signing_protocol}.
	SigningProtocol *string `field:"required" json:"signingProtocol" yaml:"signingProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_origin_access_control#description CloudfrontOriginAccessControl#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

