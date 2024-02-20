package emrserverlessapplication


type EmrserverlessApplicationTags struct {
	// The value for the tag.
	//
	// You can specify a value that is 1 to 128 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#key EmrserverlessApplication#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrserverless_application#value EmrserverlessApplication#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

