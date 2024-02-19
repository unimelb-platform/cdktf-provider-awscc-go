package cloudformationstackset


type CloudformationStackSetParameters struct {
	// The key associated with the parameter.
	//
	// If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#parameter_key CloudformationStackSet#parameter_key}
	ParameterKey *string `field:"required" json:"parameterKey" yaml:"parameterKey"`
	// The input value associated with the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#parameter_value CloudformationStackSet#parameter_value}
	ParameterValue *string `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

