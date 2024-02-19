package gameliftfleet


type GameliftFleetAnywhereConfiguration struct {
	// Cost of compute can be specified on Anywhere Fleets to prioritize placement across Queue destinations based on Cost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#cost GameliftFleet#cost}
	Cost *string `field:"required" json:"cost" yaml:"cost"`
}

