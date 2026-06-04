package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCount struct {
	// The maximum number of vCPUs. To specify no maximum limit, omit this parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_launch_template#max Ec2LaunchTemplate#max}
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of vCPUs. To specify no minimum limit, specify ``0``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_launch_template#min Ec2LaunchTemplate#min}
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

