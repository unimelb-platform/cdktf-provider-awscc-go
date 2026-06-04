package cecostcategory


type CeCostCategoryTags struct {
	// The key name for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ce_cost_category#key CeCostCategory#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ce_cost_category#value CeCostCategory#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

