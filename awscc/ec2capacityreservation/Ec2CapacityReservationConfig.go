package ec2capacityreservation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2CapacityReservationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#availability_zone Ec2CapacityReservation#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#instance_count Ec2CapacityReservation#instance_count}.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#instance_platform Ec2CapacityReservation#instance_platform}.
	InstancePlatform *string `field:"required" json:"instancePlatform" yaml:"instancePlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#instance_type Ec2CapacityReservation#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#ebs_optimized Ec2CapacityReservation#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#end_date Ec2CapacityReservation#end_date}.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#end_date_type Ec2CapacityReservation#end_date_type}.
	EndDateType *string `field:"optional" json:"endDateType" yaml:"endDateType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#ephemeral_storage Ec2CapacityReservation#ephemeral_storage}.
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#instance_match_criteria Ec2CapacityReservation#instance_match_criteria}.
	InstanceMatchCriteria *string `field:"optional" json:"instanceMatchCriteria" yaml:"instanceMatchCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#out_post_arn Ec2CapacityReservation#out_post_arn}.
	OutPostArn *string `field:"optional" json:"outPostArn" yaml:"outPostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#placement_group_arn Ec2CapacityReservation#placement_group_arn}.
	PlacementGroupArn *string `field:"optional" json:"placementGroupArn" yaml:"placementGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#tag_specifications Ec2CapacityReservation#tag_specifications}.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation#tenancy Ec2CapacityReservation#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

