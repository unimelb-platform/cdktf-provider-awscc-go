package securityhubsecuritycontrol


type SecurityhubSecurityControlParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_security_control#value SecurityhubSecurityControl#value}.
	Value *SecurityhubSecurityControlParametersValue `field:"optional" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securityhub_security_control#value_type SecurityhubSecurityControl#value_type}.
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

