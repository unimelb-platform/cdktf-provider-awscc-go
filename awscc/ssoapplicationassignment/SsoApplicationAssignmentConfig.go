package ssoapplicationassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoApplicationAssignmentConfig struct {
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
	// The ARN of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application_assignment#application_arn SsoApplicationAssignment#application_arn}
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// An identifier for an object in IAM Identity Center, such as a user or group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application_assignment#principal_id SsoApplicationAssignment#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The entity type for which the assignment will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_application_assignment#principal_type SsoApplicationAssignment#principal_type}
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
}

