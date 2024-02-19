package ec2networkinterface


type Ec2NetworkInterfaceConnectionTrackingSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#tcp_established_timeout Ec2NetworkInterface#tcp_established_timeout}.
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#udp_stream_timeout Ec2NetworkInterface#udp_stream_timeout}.
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_network_interface#udp_timeout Ec2NetworkInterface#udp_timeout}.
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

