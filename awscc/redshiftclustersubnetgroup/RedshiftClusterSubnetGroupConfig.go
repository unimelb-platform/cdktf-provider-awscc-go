package redshiftclustersubnetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedshiftClusterSubnetGroupConfig struct {
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
	// The description of the parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_subnet_group#description RedshiftClusterSubnetGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The list of VPC subnet IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_subnet_group#subnet_ids RedshiftClusterSubnetGroup#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The list of tags for the cluster parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/redshift_cluster_subnet_group#tags RedshiftClusterSubnetGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

