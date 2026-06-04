package wisdommessagetemplate


type WisdomMessageTemplateDefaultAttributesSystemAttributes struct {
	// The CustomerEndpoint attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#customer_endpoint WisdomMessageTemplate#customer_endpoint}
	CustomerEndpoint *WisdomMessageTemplateDefaultAttributesSystemAttributesCustomerEndpoint `field:"optional" json:"customerEndpoint" yaml:"customerEndpoint"`
	// The name of the task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#name WisdomMessageTemplate#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The SystemEndpoint attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_message_template#system_endpoint WisdomMessageTemplate#system_endpoint}
	SystemEndpoint *WisdomMessageTemplateDefaultAttributesSystemAttributesSystemEndpoint `field:"optional" json:"systemEndpoint" yaml:"systemEndpoint"`
}

