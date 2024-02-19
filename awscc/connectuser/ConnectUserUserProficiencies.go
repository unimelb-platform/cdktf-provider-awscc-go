package connectuser


type ConnectUserUserProficiencies struct {
	// The name of user's proficiency. You must use name of predefined attribute present in the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#attribute_name ConnectUser#attribute_name}
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// The value of user's proficiency. You must use value of predefined attribute present in the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#attribute_value ConnectUser#attribute_value}
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
	// The level of the proficiency. The valid values are 1, 2, 3, 4 and 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_user#level ConnectUser#level}
	Level *float64 `field:"required" json:"level" yaml:"level"`
}

