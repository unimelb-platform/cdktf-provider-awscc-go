package appconfigenvironment


type AppconfigEnvironmentTags struct {
	// The key-value string map.
	//
	// The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#key AppconfigEnvironment#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The tag value can be up to 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_environment#value AppconfigEnvironment#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

