package secretsmanagersecret


type SecretsmanagerSecretTags struct {
	// The value for the tag. You can specify a value that's 1 to 256 characters in length.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#key SecretsmanagerSecret#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The key name of the tag.
	//
	// You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#value SecretsmanagerSecret#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

