package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfigurationSecurityControlCustomParametersParameters struct {
	// An object that includes the data type of a security control parameter and its current value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_configuration_policy#value SecurityhubConfigurationPolicy#value}
	Value *SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfigurationSecurityControlCustomParametersParametersValue `field:"optional" json:"value" yaml:"value"`
	// Identifies whether a control parameter uses a custom user-defined value or subscribes to the default AWS Security Hub behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_configuration_policy#value_type SecurityhubConfigurationPolicy#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

