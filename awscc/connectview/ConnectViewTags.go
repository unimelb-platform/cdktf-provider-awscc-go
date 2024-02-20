package connectview


type ConnectViewTags struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_view#key ConnectView#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag. . You can specify a value that is maximum of 256 Unicode characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_view#value ConnectView#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

