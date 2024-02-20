package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiB struct {
	// The maximum amount of memory, in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#max Ec2LaunchTemplate#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory, in MiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#min Ec2LaunchTemplate#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

