package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification struct {
	// Integer value for TCP Established Timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#tcp_established_timeout Ec2LaunchTemplate#tcp_established_timeout}
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Integer value for UDP Stream Timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#udp_stream_timeout Ec2LaunchTemplate#udp_stream_timeout}
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Integer value for UDP Timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#udp_timeout Ec2LaunchTemplate#udp_timeout}
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

