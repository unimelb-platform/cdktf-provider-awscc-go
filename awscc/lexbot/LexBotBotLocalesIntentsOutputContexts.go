package lexbot


type LexBotBotLocalesIntentsOutputContexts struct {
	// Unique name for a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#name LexBot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The amount of time, in seconds, that the output context should remain active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#time_to_live_in_seconds LexBot#time_to_live_in_seconds}
	TimeToLiveInSeconds *float64 `field:"required" json:"timeToLiveInSeconds" yaml:"timeToLiveInSeconds"`
	// The number of conversation turns that the output context should remain active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#turns_to_live LexBot#turns_to_live}
	TurnsToLive *float64 `field:"required" json:"turnsToLive" yaml:"turnsToLive"`
}

