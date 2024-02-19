package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsFastLaunchConfigurationsSnapshotConfiguration struct {
	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#target_resource_count ImagebuilderDistributionConfiguration#target_resource_count}
	TargetResourceCount *float64 `field:"optional" json:"targetResourceCount" yaml:"targetResourceCount"`
}

