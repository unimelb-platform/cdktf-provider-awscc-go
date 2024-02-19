package gameliftgamesessionqueue


type GameliftGameSessionQueueTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue#key GameliftGameSessionQueue#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_game_session_queue#value GameliftGameSessionQueue#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

