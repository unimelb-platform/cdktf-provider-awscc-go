package kendraindex


type KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kendra_index#group_attribute_field KendraIndex#group_attribute_field}.
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kendra_index#user_name_attribute_field KendraIndex#user_name_attribute_field}.
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

