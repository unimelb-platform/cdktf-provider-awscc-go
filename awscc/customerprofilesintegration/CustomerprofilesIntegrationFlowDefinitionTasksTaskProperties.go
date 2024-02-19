package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionTasksTaskProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#operator_property_key CustomerprofilesIntegration#operator_property_key}.
	OperatorPropertyKey *string `field:"required" json:"operatorPropertyKey" yaml:"operatorPropertyKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#property CustomerprofilesIntegration#property}.
	Property *string `field:"required" json:"property" yaml:"property"`
}

