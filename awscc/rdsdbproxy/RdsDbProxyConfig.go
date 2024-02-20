package rdsdbproxy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsDbProxyConfig struct {
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
	// The authorization mechanism that the proxy uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#auth RdsDbProxy#auth}
	Auth interface{} `field:"required" json:"auth" yaml:"auth"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#db_proxy_name RdsDbProxy#db_proxy_name}
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The kinds of databases that the proxy can connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#engine_family RdsDbProxy#engine_family}
	EngineFamily *string `field:"required" json:"engineFamily" yaml:"engineFamily"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#role_arn RdsDbProxy#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// VPC subnet IDs to associate with the new proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#vpc_subnet_ids RdsDbProxy#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#debug_logging RdsDbProxy#debug_logging}
	DebugLogging interface{} `field:"optional" json:"debugLogging" yaml:"debugLogging"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#idle_client_timeout RdsDbProxy#idle_client_timeout}
	IdleClientTimeout *float64 `field:"optional" json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#require_tls RdsDbProxy#require_tls}
	RequireTls interface{} `field:"optional" json:"requireTls" yaml:"requireTls"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#tags RdsDbProxy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// VPC security group IDs to associate with the new proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_proxy#vpc_security_group_ids RdsDbProxy#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

