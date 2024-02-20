package organizationsorganizationalunit


type OrganizationsOrganizationalUnitTags struct {
	// The key identifier, or name, of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_organizational_unit#key OrganizationsOrganizationalUnit#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The string value that's associated with the key of the tag.
	//
	// You can set the value of a tag to an empty string, but you can't set the value of a tag to null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/organizations_organizational_unit#value OrganizationsOrganizationalUnit#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

