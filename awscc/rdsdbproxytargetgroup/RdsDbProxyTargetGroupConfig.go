package rdsdbproxytargetgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbProxyTargetGroupConfig struct {
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
	// The identifier for the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_target_group#db_proxy_name RdsDbProxyTargetGroup#db_proxy_name}
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The identifier for the DBProxyTargetGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_target_group#target_group_name RdsDbProxyTargetGroup#target_group_name}
	TargetGroupName *string `field:"required" json:"targetGroupName" yaml:"targetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_target_group#connection_pool_configuration_info RdsDbProxyTargetGroup#connection_pool_configuration_info}.
	ConnectionPoolConfigurationInfo *RdsDbProxyTargetGroupConnectionPoolConfigurationInfo `field:"optional" json:"connectionPoolConfigurationInfo" yaml:"connectionPoolConfigurationInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_target_group#db_cluster_identifiers RdsDbProxyTargetGroup#db_cluster_identifiers}.
	DbClusterIdentifiers *[]*string `field:"optional" json:"dbClusterIdentifiers" yaml:"dbClusterIdentifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_target_group#db_instance_identifiers RdsDbProxyTargetGroup#db_instance_identifiers}.
	DbInstanceIdentifiers *[]*string `field:"optional" json:"dbInstanceIdentifiers" yaml:"dbInstanceIdentifiers"`
}

