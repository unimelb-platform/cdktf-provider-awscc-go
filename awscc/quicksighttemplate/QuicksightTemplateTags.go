package quicksighttemplate


type QuicksightTemplateTags struct {
	// <p>Tag key.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#key QuicksightTemplate#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// <p>Tag value.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_template#value QuicksightTemplate#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

