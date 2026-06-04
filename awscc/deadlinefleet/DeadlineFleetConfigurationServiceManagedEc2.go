package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#instance_capabilities DeadlineFleet#instance_capabilities}.
	InstanceCapabilities *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities `field:"optional" json:"instanceCapabilities" yaml:"instanceCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#instance_market_options DeadlineFleet#instance_market_options}.
	InstanceMarketOptions *DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptions `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/deadline_fleet#storage_profile_id DeadlineFleet#storage_profile_id}.
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
}

