package databrewproject


type DatabrewProjectTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#key DatabrewProject#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_project#value DatabrewProject#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

