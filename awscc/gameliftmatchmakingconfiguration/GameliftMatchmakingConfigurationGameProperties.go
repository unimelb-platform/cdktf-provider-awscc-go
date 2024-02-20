package gameliftmatchmakingconfiguration


type GameliftMatchmakingConfigurationGameProperties struct {
	// The game property identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration#key GameliftMatchmakingConfiguration#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The game property value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration#value GameliftMatchmakingConfiguration#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

