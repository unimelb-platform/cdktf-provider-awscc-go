package s3multiregionaccesspointpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3MultiRegionAccessPointPolicyConfig struct {
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
	// The name of the Multi Region Access Point to apply policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point_policy#mrap_name S3MultiRegionAccessPointPolicy#mrap_name}
	MrapName *string `field:"required" json:"mrapName" yaml:"mrapName"`
	// Policy document to apply to a Multi Region Access Point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point_policy#policy S3MultiRegionAccessPointPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
}

