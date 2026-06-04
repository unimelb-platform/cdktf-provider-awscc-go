package gameliftcontainerfleet


type GameliftContainerFleetLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet#location GameliftContainerFleet#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Current resource capacity settings in a specified fleet or location.
	//
	// The location value might refer to a fleet's remote location or its home Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet#location_capacity GameliftContainerFleet#location_capacity}
	LocationCapacity *GameliftContainerFleetLocationsLocationCapacity `field:"optional" json:"locationCapacity" yaml:"locationCapacity"`
	// A list of fleet actions that have been suspended in the fleet location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_container_fleet#stopped_actions GameliftContainerFleet#stopped_actions}
	StoppedActions *[]*string `field:"optional" json:"stoppedActions" yaml:"stoppedActions"`
}

