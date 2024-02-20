package iotcoredeviceadvisorsuitedefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotcoredeviceadvisorSuiteDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotcoredeviceadvisor_suite_definition#suite_definition_configuration IotcoredeviceadvisorSuiteDefinition#suite_definition_configuration}.
	SuiteDefinitionConfiguration *IotcoredeviceadvisorSuiteDefinitionSuiteDefinitionConfiguration `field:"required" json:"suiteDefinitionConfiguration" yaml:"suiteDefinitionConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotcoredeviceadvisor_suite_definition#tags IotcoredeviceadvisorSuiteDefinition#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

