package redshiftclusterparametergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftClusterParameterGroupConfig struct {
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
	// A description of the parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_parameter_group#description RedshiftClusterParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The Amazon Redshift engine version to which the cluster parameter group applies.
	//
	// The cluster engine version determines the set of parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_parameter_group#parameter_group_family RedshiftClusterParameterGroup#parameter_group_family}
	ParameterGroupFamily *string `field:"required" json:"parameterGroupFamily" yaml:"parameterGroupFamily"`
	// The name of the cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_parameter_group#parameter_group_name RedshiftClusterParameterGroup#parameter_group_name}
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_parameter_group#parameters RedshiftClusterParameterGroup#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_parameter_group#tags RedshiftClusterParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

