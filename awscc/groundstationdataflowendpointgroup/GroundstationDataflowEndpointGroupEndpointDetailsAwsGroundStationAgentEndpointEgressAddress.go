package groundstationdataflowendpointgroup


type GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddress struct {
	// Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#mtu GroundstationDataflowEndpointGroup#mtu}
	Mtu *float64 `field:"optional" json:"mtu" yaml:"mtu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_dataflow_endpoint_group#socket_address GroundstationDataflowEndpointGroup#socket_address}.
	SocketAddress *GroundstationDataflowEndpointGroupEndpointDetailsAwsGroundStationAgentEndpointEgressAddressSocketAddress `field:"optional" json:"socketAddress" yaml:"socketAddress"`
}

