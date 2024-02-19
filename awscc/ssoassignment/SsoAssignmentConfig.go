package ssoassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoAssignmentConfig struct {
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
	// The sso instance that the permission set is owned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_assignment#instance_arn SsoAssignment#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The permission set that the assignemt will be assigned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_assignment#permission_set_arn SsoAssignment#permission_set_arn}
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// The assignee's identifier, user id/group id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_assignment#principal_id SsoAssignment#principal_id}
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The assignee's type, user/group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_assignment#principal_type SsoAssignment#principal_type}
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// The account id to be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_assignment#target_id SsoAssignment#target_id}
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// The type of resource to be provsioned to, only aws account now.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_assignment#target_type SsoAssignment#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

