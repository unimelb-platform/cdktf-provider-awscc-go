package ec2routetable


type Ec2RouteTableTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_route_table#key Ec2RouteTable#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_route_table#value Ec2RouteTable#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

