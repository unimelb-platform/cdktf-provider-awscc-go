package wisdommessagetemplate


type WisdomMessageTemplateContentEmailMessageTemplateContentHeaders struct {
	// The name of the email header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#name WisdomMessageTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the email header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#value WisdomMessageTemplate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

