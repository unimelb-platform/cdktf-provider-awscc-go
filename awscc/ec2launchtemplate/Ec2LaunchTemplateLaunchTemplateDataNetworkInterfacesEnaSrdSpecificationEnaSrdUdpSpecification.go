package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification struct {
	// Enables UDP ENA-SRD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#ena_srd_udp_enabled Ec2LaunchTemplate#ena_srd_udp_enabled}
	EnaSrdUdpEnabled interface{} `field:"optional" json:"enaSrdUdpEnabled" yaml:"enaSrdUdpEnabled"`
}

