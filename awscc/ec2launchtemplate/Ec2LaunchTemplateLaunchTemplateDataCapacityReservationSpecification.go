package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification struct {
	// Indicates the instance's Capacity Reservation preferences.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#capacity_reservation_preference Ec2LaunchTemplate#capacity_reservation_preference}
	CapacityReservationPreference *string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
	// Specifies a target Capacity Reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#capacity_reservation_target Ec2LaunchTemplate#capacity_reservation_target}
	CapacityReservationTarget *Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTarget `field:"optional" json:"capacityReservationTarget" yaml:"capacityReservationTarget"`
}

