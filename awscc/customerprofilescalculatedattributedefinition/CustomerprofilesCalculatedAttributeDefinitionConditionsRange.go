package customerprofilescalculatedattributedefinition


type CustomerprofilesCalculatedAttributeDefinitionConditionsRange struct {
	// The unit of time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_calculated_attribute_definition#unit CustomerprofilesCalculatedAttributeDefinition#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The amount of time of the specified unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_calculated_attribute_definition#value CustomerprofilesCalculatedAttributeDefinition#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

