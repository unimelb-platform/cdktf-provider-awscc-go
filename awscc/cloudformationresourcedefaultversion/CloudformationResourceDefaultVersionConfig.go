package cloudformationresourcedefaultversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationResourceDefaultVersionConfig struct {
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
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_resource_default_version#type_name CloudformationResourceDefaultVersion#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the type version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_resource_default_version#type_version_arn CloudformationResourceDefaultVersion#type_version_arn}
	TypeVersionArn *string `field:"optional" json:"typeVersionArn" yaml:"typeVersionArn"`
	// The ID of an existing version of the resource to set as the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_resource_default_version#version_id CloudformationResourceDefaultVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

