package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAccelerators struct {
	// The number of elastic inference accelerators to attach to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#count Ec2LaunchTemplate#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The type of elastic inference accelerator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#type Ec2LaunchTemplate#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

