package datasyncstoragesystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncStorageSystemConfig struct {
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
	// The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#agent_arns DatasyncStorageSystem#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The server name and network port required to connect with the management interface of the on-premises storage system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#server_configuration DatasyncStorageSystem#server_configuration}
	ServerConfiguration *DatasyncStorageSystemServerConfiguration `field:"required" json:"serverConfiguration" yaml:"serverConfiguration"`
	// The type of on-premises storage system that DataSync Discovery will analyze.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#system_type DatasyncStorageSystem#system_type}
	SystemType *string `field:"required" json:"systemType" yaml:"systemType"`
	// The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#cloudwatch_log_group_arn DatasyncStorageSystem#cloudwatch_log_group_arn}
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// A familiar name for the on-premises storage system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#name DatasyncStorageSystem#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The username and password for accessing your on-premises storage system's management interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#server_credentials DatasyncStorageSystem#server_credentials}
	ServerCredentials *DatasyncStorageSystemServerCredentials `field:"optional" json:"serverCredentials" yaml:"serverCredentials"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_storage_system#tags DatasyncStorageSystem#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

