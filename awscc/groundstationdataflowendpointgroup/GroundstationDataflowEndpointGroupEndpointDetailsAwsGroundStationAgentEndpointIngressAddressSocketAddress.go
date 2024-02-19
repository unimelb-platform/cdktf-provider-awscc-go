package groundstationdataflowendpointgroup


type GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointIngressAddressSocketAddress struct {
	// IPv4 socket address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#name GroundstationDataflowEndpointGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Port range of a socket address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#port_range GroundstationDataflowEndpointGroup#port_range}
	PortRange *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointIngressAddressSocketAddressPortRange `field:"optional" json:"portRange" yaml:"portRange"`
}

