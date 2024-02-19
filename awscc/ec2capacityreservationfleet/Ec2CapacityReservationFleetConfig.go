package ec2capacityreservationfleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2CapacityReservationFleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#allocation_strategy Ec2CapacityReservationFleet#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#end_date Ec2CapacityReservationFleet#end_date}.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#instance_match_criteria Ec2CapacityReservationFleet#instance_match_criteria}.
	InstanceMatchCriteria *string `field:"optional" json:"instanceMatchCriteria" yaml:"instanceMatchCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#instance_type_specifications Ec2CapacityReservationFleet#instance_type_specifications}.
	InstanceTypeSpecifications interface{} `field:"optional" json:"instanceTypeSpecifications" yaml:"instanceTypeSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#no_remove_end_date Ec2CapacityReservationFleet#no_remove_end_date}.
	NoRemoveEndDate interface{} `field:"optional" json:"noRemoveEndDate" yaml:"noRemoveEndDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#remove_end_date Ec2CapacityReservationFleet#remove_end_date}.
	RemoveEndDate interface{} `field:"optional" json:"removeEndDate" yaml:"removeEndDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#tag_specifications Ec2CapacityReservationFleet#tag_specifications}.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#tenancy Ec2CapacityReservationFleet#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet#total_target_capacity Ec2CapacityReservationFleet#total_target_capacity}.
	TotalTargetCapacity *float64 `field:"optional" json:"totalTargetCapacity" yaml:"totalTargetCapacity"`
}

