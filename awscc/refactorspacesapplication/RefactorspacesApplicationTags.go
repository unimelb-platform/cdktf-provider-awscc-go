package refactorspacesapplication


type RefactorspacesApplicationTags struct {
	// A string used to identify this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/refactorspaces_application#key RefactorspacesApplication#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string containing the value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/refactorspaces_application#value RefactorspacesApplication#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

