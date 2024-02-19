package logsaccountpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsAccountPolicyConfig struct {
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
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per PolicyType.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_account_policy#policy_document LogsAccountPolicy#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the account policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_account_policy#policy_name LogsAccountPolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Type of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_account_policy#policy_type LogsAccountPolicy#policy_type}
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Scope for policy application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_account_policy#scope LogsAccountPolicy#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Log group  selection criteria to apply policy only to a subset of log groups.
	//
	// SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/logs_account_policy#selection_criteria LogsAccountPolicy#selection_criteria}
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

