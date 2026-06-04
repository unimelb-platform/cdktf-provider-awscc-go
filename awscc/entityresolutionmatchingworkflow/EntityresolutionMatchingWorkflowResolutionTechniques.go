package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowResolutionTechniques struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#provider_properties EntityresolutionMatchingWorkflow#provider_properties}.
	ProviderProperties *EntityresolutionMatchingWorkflowResolutionTechniquesProviderProperties `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#resolution_type EntityresolutionMatchingWorkflow#resolution_type}.
	ResolutionType *string `field:"optional" json:"resolutionType" yaml:"resolutionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#rule_based_properties EntityresolutionMatchingWorkflow#rule_based_properties}.
	RuleBasedProperties *EntityresolutionMatchingWorkflowResolutionTechniquesRuleBasedProperties `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
}

