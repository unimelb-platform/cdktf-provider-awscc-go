package ec2capacityreservationfleet


type Ec2CapacityReservationFleetTagSpecificationsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#key Ec2CapacityReservationFleet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#value Ec2CapacityReservationFleet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

