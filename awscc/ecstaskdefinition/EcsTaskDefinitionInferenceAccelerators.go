package ecstaskdefinition


type EcsTaskDefinitionInferenceAccelerators struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#device_name EcsTaskDefinition#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#device_type EcsTaskDefinition#device_type}.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

