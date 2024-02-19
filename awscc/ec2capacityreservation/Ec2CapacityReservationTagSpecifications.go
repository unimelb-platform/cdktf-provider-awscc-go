package ec2capacityreservation


type Ec2CapacityReservationTagSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#resource_type Ec2CapacityReservation#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#tags Ec2CapacityReservation#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

