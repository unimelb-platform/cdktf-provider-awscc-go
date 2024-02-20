package codestarconnectionssyncconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodestarconnectionsSyncConfigurationConfig struct {
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
	// The name of the branch of the repository from which resources are to be synchronized,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_sync_configuration#branch CodestarconnectionsSyncConfiguration#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The source provider repository path of the sync configuration file of the respective SyncType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_sync_configuration#config_file CodestarconnectionsSyncConfiguration#config_file}
	ConfigFile *string `field:"required" json:"configFile" yaml:"configFile"`
	// A UUID that uniquely identifies the RepositoryLink that the SyncConfig is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_sync_configuration#repository_link_id CodestarconnectionsSyncConfiguration#repository_link_id}
	RepositoryLinkId *string `field:"required" json:"repositoryLinkId" yaml:"repositoryLinkId"`
	// The name of the resource that is being synchronized to the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_sync_configuration#resource_name CodestarconnectionsSyncConfiguration#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The IAM Role that allows AWS to update CloudFormation stacks based on content in the specified repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_sync_configuration#role_arn CodestarconnectionsSyncConfiguration#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of resource synchronization service that is to be configured, for example, CFN_STACK_SYNC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/codestarconnections_sync_configuration#sync_type CodestarconnectionsSyncConfiguration#sync_type}
	SyncType *string `field:"required" json:"syncType" yaml:"syncType"`
}

