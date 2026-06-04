package entityresolutionidnamespace


type EntityresolutionIdNamespaceIdMappingWorkflowProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_id_namespace#id_mapping_type EntityresolutionIdNamespace#id_mapping_type}.
	IdMappingType *string `field:"optional" json:"idMappingType" yaml:"idMappingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_id_namespace#provider_properties EntityresolutionIdNamespace#provider_properties}.
	ProviderProperties *EntityresolutionIdNamespaceIdMappingWorkflowPropertiesProviderProperties `field:"optional" json:"providerProperties" yaml:"providerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_id_namespace#rule_based_properties EntityresolutionIdNamespace#rule_based_properties}.
	RuleBasedProperties *EntityresolutionIdNamespaceIdMappingWorkflowPropertiesRuleBasedProperties `field:"optional" json:"ruleBasedProperties" yaml:"ruleBasedProperties"`
}

