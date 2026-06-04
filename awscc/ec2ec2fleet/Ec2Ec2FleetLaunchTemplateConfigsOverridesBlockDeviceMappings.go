package ec2ec2fleet


type Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_ec2_fleet#device_name Ec2Ec2Fleet#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_ec2_fleet#ebs Ec2Ec2Fleet#ebs}.
	Ebs *Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_ec2_fleet#no_device Ec2Ec2Fleet#no_device}.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_ec2_fleet#virtual_name Ec2Ec2Fleet#virtual_name}.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

