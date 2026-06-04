package athenacapacityreservation


type AthenaCapacityReservationCapacityAssignmentConfiguration struct {
	// List of capacity assignments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/athena_capacity_reservation#capacity_assignments AthenaCapacityReservation#capacity_assignments}
	CapacityAssignments interface{} `field:"optional" json:"capacityAssignments" yaml:"capacityAssignments"`
}

