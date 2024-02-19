package ec2launchtemplate


type Ec2LaunchTemplateTagSpecificationsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#key Ec2LaunchTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#value Ec2LaunchTemplate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

