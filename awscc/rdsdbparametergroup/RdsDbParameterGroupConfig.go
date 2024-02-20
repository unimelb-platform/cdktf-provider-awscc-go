package rdsdbparametergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbParameterGroupConfig struct {
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
	// Provides the customer-specified description for this DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_parameter_group#description RdsDbParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB parameter group family name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_parameter_group#family RdsDbParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// Specifies the name of the DB parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_parameter_group#db_parameter_group_name RdsDbParameterGroup#db_parameter_group_name}
	DbParameterGroupName *string `field:"optional" json:"dbParameterGroupName" yaml:"dbParameterGroupName"`
	// An array of parameter names and values for the parameter update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_parameter_group#parameters RdsDbParameterGroup#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_parameter_group#tags RdsDbParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

