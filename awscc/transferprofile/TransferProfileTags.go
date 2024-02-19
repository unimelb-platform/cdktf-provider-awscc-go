package transferprofile


type TransferProfileTags struct {
	// The name assigned to the tag that you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_profile#key TransferProfile#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Contains one or more values that you assigned to the key name you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_profile#value TransferProfile#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

