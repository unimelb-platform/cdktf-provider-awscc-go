package cloudformationmoduleversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationModuleVersionConfig struct {
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
	// The name of the module being registered.
	//
	// Recommended module naming pattern: company_or_organization::service::type::MODULE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_module_version#module_name CloudformationModuleVersion#module_name}
	ModuleName *string `field:"required" json:"moduleName" yaml:"moduleName"`
	// The url to the S3 bucket containing the schema and template fragment for the module you want to register.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_module_version#module_package CloudformationModuleVersion#module_package}
	ModulePackage *string `field:"required" json:"modulePackage" yaml:"modulePackage"`
}

