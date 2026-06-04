package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionTriggerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_integration#trigger_properties CustomerprofilesIntegration#trigger_properties}.
	TriggerProperties *CustomerprofilesIntegrationFlowDefinitionTriggerConfigTriggerProperties `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/customerprofiles_integration#trigger_type CustomerprofilesIntegration#trigger_type}.
	TriggerType *string `field:"optional" json:"triggerType" yaml:"triggerType"`
}

