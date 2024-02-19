package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappings struct {
	// The user data to make available to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#device_name Ec2LaunchTemplate#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Parameters for a block device for an EBS volume in an Amazon EC2 launch template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ebs Ec2LaunchTemplate#ebs}
	Ebs *Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// To omit the device from the block device mapping, specify an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#no_device Ec2LaunchTemplate#no_device}
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// The virtual device name (ephemeralN).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#virtual_name Ec2LaunchTemplate#virtual_name}
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

