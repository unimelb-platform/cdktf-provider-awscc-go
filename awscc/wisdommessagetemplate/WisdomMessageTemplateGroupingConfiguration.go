package wisdommessagetemplate


type WisdomMessageTemplateGroupingConfiguration struct {
	// The criteria used for grouping Amazon Q in Connect users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#criteria WisdomMessageTemplate#criteria}
	Criteria *string `field:"optional" json:"criteria" yaml:"criteria"`
	// The list of values that define different groups of Amazon Q in Connect users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#values WisdomMessageTemplate#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

