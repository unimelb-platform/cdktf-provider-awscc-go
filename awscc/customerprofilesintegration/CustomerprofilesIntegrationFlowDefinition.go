package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#flow_name CustomerprofilesIntegration#flow_name}.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#kms_arn CustomerprofilesIntegration#kms_arn}.
	KmsArn *string `field:"required" json:"kmsArn" yaml:"kmsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#source_flow_config CustomerprofilesIntegration#source_flow_config}.
	SourceFlowConfig *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig `field:"required" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#tasks CustomerprofilesIntegration#tasks}.
	Tasks interface{} `field:"required" json:"tasks" yaml:"tasks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#trigger_config CustomerprofilesIntegration#trigger_config}.
	TriggerConfig *CustomerprofilesIntegrationFlowDefinitionTriggerConfig `field:"required" json:"triggerConfig" yaml:"triggerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#description CustomerprofilesIntegration#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

