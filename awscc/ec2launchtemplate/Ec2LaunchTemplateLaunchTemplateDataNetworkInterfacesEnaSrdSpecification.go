package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecification struct {
	// Enables TCP ENA-SRD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ena_srd_enabled Ec2LaunchTemplate#ena_srd_enabled}
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// Allows customer to specify ENA-SRD (UDP) options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ena_srd_udp_specification Ec2LaunchTemplate#ena_srd_udp_specification}
	EnaSrdUdpSpecification *Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

