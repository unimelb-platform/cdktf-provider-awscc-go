package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionTriggerConfigTriggerPropertiesScheduled struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#schedule_expression CustomerprofilesIntegration#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#data_pull_mode CustomerprofilesIntegration#data_pull_mode}.
	DataPullMode *string `field:"optional" json:"dataPullMode" yaml:"dataPullMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#first_execution_from CustomerprofilesIntegration#first_execution_from}.
	FirstExecutionFrom *float64 `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#schedule_end_time CustomerprofilesIntegration#schedule_end_time}.
	ScheduleEndTime *float64 `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#schedule_offset CustomerprofilesIntegration#schedule_offset}.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#schedule_start_time CustomerprofilesIntegration#schedule_start_time}.
	ScheduleStartTime *float64 `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#timezone CustomerprofilesIntegration#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

