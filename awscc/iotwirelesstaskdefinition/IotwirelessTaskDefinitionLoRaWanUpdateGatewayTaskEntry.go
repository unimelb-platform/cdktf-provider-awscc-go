package iotwirelesstaskdefinition


type IotwirelessTaskDefinitionLoRaWanUpdateGatewayTaskEntry struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#current_version IotwirelessTaskDefinition#current_version}.
	CurrentVersion *IotwirelessTaskDefinitionLoRaWanUpdateGatewayTaskEntryCurrentVersion `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotwireless_task_definition#update_version IotwirelessTaskDefinition#update_version}.
	UpdateVersion *IotwirelessTaskDefinitionLoRaWanUpdateGatewayTaskEntryUpdateVersion `field:"optional" json:"updateVersion" yaml:"updateVersion"`
}

