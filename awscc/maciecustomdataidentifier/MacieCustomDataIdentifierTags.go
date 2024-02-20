package maciecustomdataidentifier


type MacieCustomDataIdentifierTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_custom_data_identifier#key MacieCustomDataIdentifier#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_custom_data_identifier#value MacieCustomDataIdentifier#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

