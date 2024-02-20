package appconfigconfigurationprofile


type AppconfigConfigurationProfileTags struct {
	// The key-value string map. The tag key can be up to 128 characters and must not start with aws:.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#key AppconfigConfigurationProfile#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value can be up to 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#value AppconfigConfigurationProfile#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

