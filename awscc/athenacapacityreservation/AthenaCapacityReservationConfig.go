package athenacapacityreservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AthenaCapacityReservationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The reservation name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_capacity_reservation#name AthenaCapacityReservation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of DPUs to request to be allocated to the reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_capacity_reservation#target_dpus AthenaCapacityReservation#target_dpus}
	TargetDpus *float64 `field:"required" json:"targetDpus" yaml:"targetDpus"`
	// Assignment configuration to assign workgroups to a reservation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_capacity_reservation#capacity_assignment_configuration AthenaCapacityReservation#capacity_assignment_configuration}
	CapacityAssignmentConfiguration *AthenaCapacityReservationCapacityAssignmentConfiguration `field:"optional" json:"capacityAssignmentConfiguration" yaml:"capacityAssignmentConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_capacity_reservation#tags AthenaCapacityReservation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

