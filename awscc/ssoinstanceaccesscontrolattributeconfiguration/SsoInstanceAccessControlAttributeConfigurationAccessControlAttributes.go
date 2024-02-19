package ssoinstanceaccesscontrolattributeconfiguration


type SsoInstanceAccessControlAttributeConfigurationAccessControlAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_instance_access_control_attribute_configuration#key SsoInstanceAccessControlAttributeConfiguration#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_instance_access_control_attribute_configuration#value SsoInstanceAccessControlAttributeConfiguration#value}.
	Value *SsoInstanceAccessControlAttributeConfigurationAccessControlAttributesValue `field:"required" json:"value" yaml:"value"`
}

