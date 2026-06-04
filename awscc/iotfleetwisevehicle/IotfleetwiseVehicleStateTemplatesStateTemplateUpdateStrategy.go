package iotfleetwisevehicle


type IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_vehicle#on_change IotfleetwiseVehicle#on_change}.
	OnChange *string `field:"optional" json:"onChange" yaml:"onChange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotfleetwise_vehicle#periodic IotfleetwiseVehicle#periodic}.
	Periodic *IotfleetwiseVehicleStateTemplatesStateTemplateUpdateStrategyPeriodic `field:"optional" json:"periodic" yaml:"periodic"`
}

