package ec2capacityreservation


type Ec2CapacityReservationTagSpecificationsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_capacity_reservation#key Ec2CapacityReservation#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_capacity_reservation#value Ec2CapacityReservation#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

