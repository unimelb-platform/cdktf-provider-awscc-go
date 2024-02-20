package customerprofilesobjecttype


type CustomerprofilesObjectTypeFieldsObjectTypeField struct {
	// The content type of the field. Used for determining equality when searching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_object_type#content_type CustomerprofilesObjectType#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// A field of a ProfileObject.
	//
	// For example: _source.FirstName, where "_source" is a ProfileObjectType of a Zendesk user and "FirstName" is a field in that ObjectType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_object_type#source CustomerprofilesObjectType#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The location of the data in the standard ProfileObject model. For example: _profile.Address.PostalCode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_object_type#target CustomerprofilesObjectType#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

