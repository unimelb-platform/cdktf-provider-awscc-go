package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions struct {
	// Disables the automatic recovery behavior of your instance or sets it to default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_launch_template#auto_recovery Ec2LaunchTemplate#auto_recovery}
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

