package appconfigconfigurationprofile


type AppconfigConfigurationProfileValidators struct {
	// Either the JSON Schema content or the Amazon Resource Name (ARN) of an Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#content AppconfigConfigurationProfile#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// AWS AppConfig supports validators of type JSON_SCHEMA and LAMBDA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appconfig_configuration_profile#type AppconfigConfigurationProfile#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

