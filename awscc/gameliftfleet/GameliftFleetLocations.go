package gameliftfleet


type GameliftFleetLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#location GameliftFleet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#location_capacity GameliftFleet#location_capacity}
	LocationCapacity *GameliftFleetLocationsLocationCapacity `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
}

