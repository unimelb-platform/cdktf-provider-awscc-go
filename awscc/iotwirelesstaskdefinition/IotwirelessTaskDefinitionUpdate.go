package iotwirelesstaskdefinition


type IotwirelessTaskDefinitionUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#lo_ra_wan IotwirelessTaskDefinition#lo_ra_wan}.
	LoRaWan *IotwirelessTaskDefinitionUpdateLoRaWan `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#update_data_role IotwirelessTaskDefinition#update_data_role}.
	UpdateDataRole *string `field:"optional" json:"updateDataRole" yaml:"updateDataRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#update_data_source IotwirelessTaskDefinition#update_data_source}.
	UpdateDataSource *string `field:"optional" json:"updateDataSource" yaml:"updateDataSource"`
}

