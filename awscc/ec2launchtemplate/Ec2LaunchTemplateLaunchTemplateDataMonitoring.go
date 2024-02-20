package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataMonitoring struct {
	// Specify true to enable detailed monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#enabled Ec2LaunchTemplate#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

