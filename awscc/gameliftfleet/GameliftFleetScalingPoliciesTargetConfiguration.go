package gameliftfleet


type GameliftFleetScalingPoliciesTargetConfiguration struct {
	// Desired value to use with a target-based scaling policy.
	//
	// The value must be relevant for whatever metric the scaling policy is using. For example, in a policy using the metric PercentAvailableGameSessions, the target value should be the preferred size of the fleet's buffer (the percent of capacity that should be idle and ready for new game sessions).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/gamelift_fleet#target_value GameliftFleet#target_value}
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
}

