package iotprovisioningtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotProvisioningTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#provisioning_role_arn IotProvisioningTemplate#provisioning_role_arn}.
	ProvisioningRoleArn *string `field:"required" json:"provisioningRoleArn" yaml:"provisioningRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#template_body IotProvisioningTemplate#template_body}.
	TemplateBody *string `field:"required" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#description IotProvisioningTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#enabled IotProvisioningTemplate#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#pre_provisioning_hook IotProvisioningTemplate#pre_provisioning_hook}.
	PreProvisioningHook *IotProvisioningTemplatePreProvisioningHook `field:"optional" json:"preProvisioningHook" yaml:"preProvisioningHook"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#tags IotProvisioningTemplate#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#template_name IotProvisioningTemplate#template_name}.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_provisioning_template#template_type IotProvisioningTemplate#template_type}.
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}

