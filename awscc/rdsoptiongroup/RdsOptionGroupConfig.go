package rdsoptiongroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsOptionGroupConfig struct {
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
	// Indicates the name of the engine that this option group can be applied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#engine_name RdsOptionGroup#engine_name}
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// Indicates the major engine version associated with this option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#major_engine_version RdsOptionGroup#major_engine_version}
	MajorEngineVersion *string `field:"required" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// Provides a description of the option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#option_group_description RdsOptionGroup#option_group_description}
	OptionGroupDescription *string `field:"required" json:"optionGroupDescription" yaml:"optionGroupDescription"`
	// Indicates what options are available in the option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#option_configurations RdsOptionGroup#option_configurations}
	OptionConfigurations interface{} `field:"optional" json:"optionConfigurations" yaml:"optionConfigurations"`
	// Specifies the name of the option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#option_group_name RdsOptionGroup#option_group_name}
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_option_group#tags RdsOptionGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

