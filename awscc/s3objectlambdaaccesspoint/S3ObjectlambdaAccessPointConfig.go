package s3objectlambdaaccesspoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ObjectlambdaAccessPointConfig struct {
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
	// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point#object_lambda_configuration S3ObjectlambdaAccessPoint#object_lambda_configuration}
	ObjectLambdaConfiguration *S3ObjectlambdaAccessPointObjectLambdaConfiguration `field:"required" json:"objectLambdaConfiguration" yaml:"objectLambdaConfiguration"`
	// The name you want to assign to this Object lambda Access Point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point#name S3ObjectlambdaAccessPoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

