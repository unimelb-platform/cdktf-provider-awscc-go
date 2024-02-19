package gameliftgamesessionqueue


type GameliftGameSessionQueuePriorityConfiguration struct {
	// The prioritization order to use for fleet locations, when the PriorityOrder property includes LOCATION.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue#location_order GameliftGameSessionQueue#location_order}
	LocationOrder *[]*string `field:"optional" json:"locationOrder" yaml:"locationOrder"`
	// The recommended sequence to use when prioritizing where to place new game sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue#priority_order GameliftGameSessionQueue#priority_order}
	PriorityOrder *[]*string `field:"optional" json:"priorityOrder" yaml:"priorityOrder"`
}

