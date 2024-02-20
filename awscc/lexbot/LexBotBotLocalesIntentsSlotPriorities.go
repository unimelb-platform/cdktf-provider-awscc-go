package lexbot


type LexBotBotLocalesIntentsSlotPriorities struct {
	// The priority that a slot should be elicited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#priority LexBot#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// The name of the slot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#slot_name LexBot#slot_name}
	SlotName *string `field:"required" json:"slotName" yaml:"slotName"`
}

