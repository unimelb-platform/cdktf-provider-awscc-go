package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#max DeadlineFleet#max}.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#min DeadlineFleet#min}.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#name DeadlineFleet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

