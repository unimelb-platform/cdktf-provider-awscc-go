package gameliftfleet


type GameliftFleetLocationsLocationCapacity struct {
	// The number of EC2 instances you want to maintain in the specified fleet location.
	//
	// This value must fall between the minimum and maximum size limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#desired_ec2_instances GameliftFleet#desired_ec2_instances}
	DesiredEc2Instances *float64 `field:"required" json:"desiredEc2Instances" yaml:"desiredEc2Instances"`
	// The maximum value that is allowed for the fleet's instance count for a location.
	//
	// When creating a new fleet, GameLift automatically sets this value to "1". Once the fleet is active, you can change this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#max_size GameliftFleet#max_size}
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// The minimum value allowed for the fleet's instance count for a location.
	//
	// When creating a new fleet, GameLift automatically sets this value to "0". After the fleet is active, you can change this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#min_size GameliftFleet#min_size}
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
}

