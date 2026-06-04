package ssoinstanceaccesscontrolattributeconfiguration


type SsoInstanceAccessControlAttributeConfigurationInstanceAccessControlAttributeConfigurationAccessControlAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_instance_access_control_attribute_configuration#key SsoInstanceAccessControlAttributeConfiguration#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sso_instance_access_control_attribute_configuration#value SsoInstanceAccessControlAttributeConfiguration#value}.
	Value *SsoInstanceAccessControlAttributeConfigurationInstanceAccessControlAttributeConfigurationAccessControlAttributesValue `field:"optional" json:"value" yaml:"value"`
}

