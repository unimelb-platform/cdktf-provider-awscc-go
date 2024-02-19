package cloudfrontcloudfrontoriginaccessidentity

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfrontCloudfrontOriginAccessIdentityConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_cloudfront_origin_access_identity#cloudfront_origin_access_identity_config CloudfrontCloudfrontOriginAccessIdentity#cloudfront_origin_access_identity_config}.
	CloudfrontOriginAccessIdentityConfig *CloudfrontCloudfrontOriginAccessIdentityCloudfrontOriginAccessIdentityConfig `field:"required" json:"cloudfrontOriginAccessIdentityConfig" yaml:"cloudfrontOriginAccessIdentityConfig"`
}

