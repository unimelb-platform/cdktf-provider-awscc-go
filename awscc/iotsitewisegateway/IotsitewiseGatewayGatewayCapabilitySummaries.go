package iotsitewisegateway


type IotsitewiseGatewayGatewayCapabilitySummaries struct {
	// The namespace of the capability configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#capability_namespace IotsitewiseGateway#capability_namespace}
	CapabilityNamespace *string `field:"required" json:"capabilityNamespace" yaml:"capabilityNamespace"`
	// The JSON document that defines the gateway capability's configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#capability_configuration IotsitewiseGateway#capability_configuration}
	CapabilityConfiguration *string `field:"optional" json:"capabilityConfiguration" yaml:"capabilityConfiguration"`
}

