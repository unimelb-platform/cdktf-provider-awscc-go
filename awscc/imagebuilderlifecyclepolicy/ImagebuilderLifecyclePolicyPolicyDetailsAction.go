package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailsAction struct {
	// The action type of the policy detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#type ImagebuilderLifecyclePolicy#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The included resources of the policy detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#include_resources ImagebuilderLifecyclePolicy#include_resources}
	IncludeResources *ImagebuilderLifecyclePolicyPolicyDetailsActionIncludeResources `field:"optional" json:"includeResources" yaml:"includeResources"`
}

