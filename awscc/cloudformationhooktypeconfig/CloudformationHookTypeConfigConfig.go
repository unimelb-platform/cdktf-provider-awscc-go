package cloudformationhooktypeconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationHookTypeConfigConfig struct {
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
	// The configuration data for the extension, in this account and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_hook_type_config#configuration CloudformationHookTypeConfig#configuration}
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// An alias by which to refer to this extension configuration data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_hook_type_config#configuration_alias CloudformationHookTypeConfig#configuration_alias}
	ConfigurationAlias *string `field:"optional" json:"configurationAlias" yaml:"configurationAlias"`
	// The Amazon Resource Name (ARN) of the type without version number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_hook_type_config#type_arn CloudformationHookTypeConfig#type_arn}
	TypeArn *string `field:"optional" json:"typeArn" yaml:"typeArn"`
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_hook_type_config#type_name CloudformationHookTypeConfig#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
}

