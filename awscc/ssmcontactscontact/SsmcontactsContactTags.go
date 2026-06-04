package ssmcontactscontact


type SsmcontactsContactTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ssmcontacts_contact#key SsmcontactsContact#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ssmcontacts_contact#value SsmcontactsContact#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

