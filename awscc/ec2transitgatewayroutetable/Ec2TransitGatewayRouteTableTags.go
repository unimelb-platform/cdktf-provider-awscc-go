package ec2transitgatewayroutetable


type Ec2TransitGatewayRouteTableTags struct {
	// The key of the associated tag key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_route_table#key Ec2TransitGatewayRouteTable#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the associated tag key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_route_table#value Ec2TransitGatewayRouteTable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

