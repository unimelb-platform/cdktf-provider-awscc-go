package s3outpostsaccesspoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3OutpostsAccessPointConfig struct {
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
	// The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_access_point#bucket S3OutpostsAccessPoint#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A name for the AccessPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_access_point#name S3OutpostsAccessPoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_access_point#vpc_configuration S3OutpostsAccessPoint#vpc_configuration}
	VpcConfiguration *S3OutpostsAccessPointVpcConfiguration `field:"required" json:"vpcConfiguration" yaml:"vpcConfiguration"`
	// The access point policy associated with this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_access_point#policy S3OutpostsAccessPoint#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

