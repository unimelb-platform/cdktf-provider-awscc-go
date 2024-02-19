package ec2capacityreservation


type Ec2CapacityReservationTagSpecificationsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#key Ec2CapacityReservation#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#value Ec2CapacityReservation#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

