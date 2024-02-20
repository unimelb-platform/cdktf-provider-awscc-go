package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionConditionsThreshold struct {
	// The operator of the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_calculated_attribute_definition#operator CustomerprofilesCalculatedAttributeDefinition#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value of the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_calculated_attribute_definition#value CustomerprofilesCalculatedAttributeDefinition#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

