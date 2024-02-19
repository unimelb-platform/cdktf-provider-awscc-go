package sescontactlist


type SesContactListTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_contact_list#key SesContactList#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_contact_list#value SesContactList#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

