package maciefindingsfilter


type MacieFindingsFilterTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#key MacieFindingsFilter#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#value MacieFindingsFilter#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

