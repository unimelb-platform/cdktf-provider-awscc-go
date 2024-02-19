package rdsdbclusterparametergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbClusterParameterGroupConfig struct {
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
	// A friendly description for this DB cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster_parameter_group#description RdsDbClusterParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The DB cluster parameter group family name.
	//
	// A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a DB engine and engine version compatible with that DB cluster parameter group family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster_parameter_group#family RdsDbClusterParameterGroup#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster_parameter_group#parameters RdsDbClusterParameterGroup#parameters}
	Parameters *string `field:"required" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster_parameter_group#db_cluster_parameter_group_name RdsDbClusterParameterGroup#db_cluster_parameter_group_name}.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// The list of tags for the cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_cluster_parameter_group#tags RdsDbClusterParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

