package customerprofilesobjecttype


type CustomerprofilesObjectTypeFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_object_type#name CustomerprofilesObjectType#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Represents a field in a ProfileObjectType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_object_type#object_type_field CustomerprofilesObjectType#object_type_field}
	ObjectTypeField *CustomerprofilesObjectTypeFieldsObjectTypeField `field:"optional" json:"objectTypeField" yaml:"objectTypeField"`
}

