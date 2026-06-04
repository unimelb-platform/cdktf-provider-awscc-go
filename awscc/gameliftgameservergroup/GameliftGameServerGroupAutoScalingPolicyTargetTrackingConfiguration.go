package gameliftgameservergroup


type GameliftGameServerGroupAutoScalingPolicyTargetTrackingConfiguration struct {
	// Desired value to use with a game server group target-based scaling policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/gamelift_game_server_group#target_value GameliftGameServerGroup#target_value}
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

