package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptions struct {
	// Deprecated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#block_duration_minutes Ec2LaunchTemplate#block_duration_minutes}
	BlockDurationMinutes *float64 `field:"optional" json:"blockDurationMinutes" yaml:"blockDurationMinutes"`
	// The behavior when a Spot Instance is interrupted. The default is terminate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#instance_interruption_behavior Ec2LaunchTemplate#instance_interruption_behavior}
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// The maximum hourly price you're willing to pay for the Spot Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#max_price Ec2LaunchTemplate#max_price}
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// The Spot Instance request type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#spot_instance_type Ec2LaunchTemplate#spot_instance_type}
	SpotInstanceType *string `field:"optional" json:"spotInstanceType" yaml:"spotInstanceType"`
	// The end date of the request, in UTC format (YYYY-MM-DDTHH:MM:SSZ). Supported only for persistent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#valid_until Ec2LaunchTemplate#valid_until}
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

