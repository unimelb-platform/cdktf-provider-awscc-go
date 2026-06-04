package gameliftcontainerfleet


type GameliftContainerFleetLocationsLocationCapacity struct {
	// The number of EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet#desired_ec2_instances GameliftContainerFleet#desired_ec2_instances}
	DesiredEc2Instances *float64 `field:"optional" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The maximum value that is allowed for the fleet's instance count for a location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet#max_size GameliftContainerFleet#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum value allowed for the fleet's instance count for a location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet#min_size GameliftContainerFleet#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

