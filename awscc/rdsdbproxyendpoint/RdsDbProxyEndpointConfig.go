package rdsdbproxyendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbProxyEndpointConfig struct {
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
	// The identifier for the DB proxy endpoint.
	//
	// This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_endpoint#db_proxy_endpoint_name RdsDbProxyEndpoint#db_proxy_endpoint_name}
	DbProxyEndpointName *string `field:"required" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_endpoint#db_proxy_name RdsDbProxyEndpoint#db_proxy_name}
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// VPC subnet IDs to associate with the new DB proxy endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_endpoint#vpc_subnet_ids RdsDbProxyEndpoint#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_endpoint#tags RdsDbProxyEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_endpoint#target_role RdsDbProxyEndpoint#target_role}
	TargetRole *string `field:"optional" json:"targetRole" yaml:"targetRole"`
	// VPC security group IDs to associate with the new DB proxy endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy_endpoint#vpc_security_group_ids RdsDbProxyEndpoint#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

