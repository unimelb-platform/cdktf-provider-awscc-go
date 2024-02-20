package ssoinstanceaccesscontrolattributeconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsoInstanceAccessControlAttributeConfigurationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ARN of the AWS SSO instance under which the operation will be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_instance_access_control_attribute_configuration#instance_arn SsoInstanceAccessControlAttributeConfiguration#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_instance_access_control_attribute_configuration#access_control_attributes SsoInstanceAccessControlAttributeConfiguration#access_control_attributes}.
	AccessControlAttributes interface{} `field:"optional" json:"accessControlAttributes" yaml:"accessControlAttributes"`
	// The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes.
	//
	// We recomend that you use  AccessControlAttributes property instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sso_instance_access_control_attribute_configuration#instance_access_control_attribute_configuration SsoInstanceAccessControlAttributeConfiguration#instance_access_control_attribute_configuration}
	InstanceAccessControlAttributeConfiguration *SsoInstanceAccessControlAttributeConfigurationInstanceAccessControlAttributeConfiguration `field:"optional" json:"instanceAccessControlAttributeConfiguration" yaml:"instanceAccessControlAttributeConfiguration"`
}

