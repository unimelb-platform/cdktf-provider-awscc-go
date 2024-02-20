package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions struct {
	// The market type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#market_type Ec2LaunchTemplate#market_type}
	MarketType *string `field:"optional" json:"marketType" yaml:"marketType"`
	// Specifies options for Spot Instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#spot_options Ec2LaunchTemplate#spot_options}
	SpotOptions *Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
}

