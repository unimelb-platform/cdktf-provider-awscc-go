package s3multiregionaccesspoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3MultiRegionAccessPointConfig struct {
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
	// The list of buckets that you want to associate this Multi Region Access Point with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point#regions S3MultiRegionAccessPoint#regions}
	Regions interface{} `field:"required" json:"regions" yaml:"regions"`
	// The name you want to assign to this Multi Region Access Point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point#name S3MultiRegionAccessPoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point.
	//
	// You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point#public_access_block_configuration S3MultiRegionAccessPoint#public_access_block_configuration}
	PublicAccessBlockConfiguration *S3MultiRegionAccessPointPublicAccessBlockConfiguration `field:"optional" json:"publicAccessBlockConfiguration" yaml:"publicAccessBlockConfiguration"`
}

