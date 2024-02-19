package s3objectlambdaaccesspointpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3ObjectlambdaAccessPointPolicyConfig struct {
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
	// The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point_policy#object_lambda_access_point S3ObjectlambdaAccessPointPolicy#object_lambda_access_point}
	ObjectLambdaAccessPoint *string `field:"required" json:"objectLambdaAccessPoint" yaml:"objectLambdaAccessPoint"`
	// A policy document containing permissions to add to the specified ObjectLambdaAccessPoint.
	//
	// For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3objectlambda_access_point_policy#policy_document S3ObjectlambdaAccessPointPolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

