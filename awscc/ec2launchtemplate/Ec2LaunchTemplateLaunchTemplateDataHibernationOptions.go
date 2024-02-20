package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataHibernationOptions struct {
	// TIf you set this parameter to true, the instance is enabled for hibernation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#configured Ec2LaunchTemplate#configured}
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

