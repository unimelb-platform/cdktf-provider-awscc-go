package securityhubpolicyassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubPolicyAssociationConfig struct {
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
	// The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_policy_association#configuration_policy_id SecurityhubPolicyAssociation#configuration_policy_id}
	ConfigurationPolicyId *string `field:"required" json:"configurationPolicyId" yaml:"configurationPolicyId"`
	// The identifier of the target account, organizational unit, or the root.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_policy_association#target_id SecurityhubPolicyAssociation#target_id}
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Indicates whether the target is an AWS account, organizational unit, or the organization root.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_policy_association#target_type SecurityhubPolicyAssociation#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

