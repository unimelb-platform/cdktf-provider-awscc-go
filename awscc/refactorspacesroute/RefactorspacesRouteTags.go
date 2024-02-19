package refactorspacesroute


type RefactorspacesRouteTags struct {
	// A string used to identify this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_route#key RefactorspacesRoute#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/refactorspaces_route#value RefactorspacesRoute#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

