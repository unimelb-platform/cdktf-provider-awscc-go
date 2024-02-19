package resiliencehubresiliencypolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResiliencehubResiliencyPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#policy ResiliencehubResiliencyPolicy#policy}.
	Policy interface{} `field:"required" json:"policy" yaml:"policy"`
	// Name of Resiliency Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#policy_name ResiliencehubResiliencyPolicy#policy_name}
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Resiliency Policy Tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#tier ResiliencehubResiliencyPolicy#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Data Location Constraint of the Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#data_location_constraint ResiliencehubResiliencyPolicy#data_location_constraint}
	DataLocationConstraint *string `field:"optional" json:"dataLocationConstraint" yaml:"dataLocationConstraint"`
	// Description of Resiliency Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#policy_description ResiliencehubResiliencyPolicy#policy_description}
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/resiliencehub_resiliency_policy#tags ResiliencehubResiliencyPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

