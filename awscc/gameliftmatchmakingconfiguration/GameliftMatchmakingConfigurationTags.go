package gameliftmatchmakingconfiguration


type GameliftMatchmakingConfigurationTags struct {
	// The key name of the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration#key GameliftMatchmakingConfiguration#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_matchmaking_configuration#value GameliftMatchmakingConfiguration#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

