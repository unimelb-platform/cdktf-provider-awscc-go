package wisdommessagetemplate


type WisdomMessageTemplateDefaultAttributesAgentAttributes struct {
	// The agent?s first name as entered in their Amazon Connect user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#first_name WisdomMessageTemplate#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// The agent?s last name as entered in their Amazon Connect user account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#last_name WisdomMessageTemplate#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
}

