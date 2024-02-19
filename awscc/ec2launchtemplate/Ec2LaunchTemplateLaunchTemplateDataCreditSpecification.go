package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataCreditSpecification struct {
	// The user data to make available to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#cpu_credits Ec2LaunchTemplate#cpu_credits}
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
}

