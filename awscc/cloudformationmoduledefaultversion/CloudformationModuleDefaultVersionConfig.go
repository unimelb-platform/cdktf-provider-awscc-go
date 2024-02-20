package cloudformationmoduledefaultversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationModuleDefaultVersionConfig struct {
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
	// The Amazon Resource Name (ARN) of the module version to set as the default version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_module_default_version#arn CloudformationModuleDefaultVersion#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name of a module existing in the registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_module_default_version#module_name CloudformationModuleDefaultVersion#module_name}
	ModuleName *string `field:"optional" json:"moduleName" yaml:"moduleName"`
	// The ID of an existing version of the named module to set as the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_module_default_version#version_id CloudformationModuleDefaultVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

