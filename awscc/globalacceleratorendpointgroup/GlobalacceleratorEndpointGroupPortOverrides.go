package globalacceleratorendpointgroup


type GlobalacceleratorEndpointGroupPortOverrides struct {
	// A network port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group#endpoint_port GlobalacceleratorEndpointGroup#endpoint_port}
	EndpointPort *float64 `field:"required" json:"endpointPort" yaml:"endpointPort"`
	// A network port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_endpoint_group#listener_port GlobalacceleratorEndpointGroup#listener_port}
	ListenerPort *float64 `field:"required" json:"listenerPort" yaml:"listenerPort"`
}

