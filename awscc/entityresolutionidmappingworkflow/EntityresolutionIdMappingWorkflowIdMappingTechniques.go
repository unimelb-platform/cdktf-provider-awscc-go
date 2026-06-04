package entityresolutionidmappingworkflow


type EntityresolutionIdMappingWorkflowIdMappingTechniques struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_id_mapping_workflow#id_mapping_type EntityresolutionIdMappingWorkflow#id_mapping_type}.
	IdMappingType *string `field:"optional" json:"idMappingType" yaml:"idMappingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_id_mapping_workflow#provider_properties EntityresolutionIdMappingWorkflow#provider_properties}.
	ProviderProperties *EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderProperties `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_id_mapping_workflow#rule_based_properties EntityresolutionIdMappingWorkflow#rule_based_properties}.
	RuleBasedProperties *EntityresolutionIdMappingWorkflowIdMappingTechniquesRuleBasedProperties `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
}

