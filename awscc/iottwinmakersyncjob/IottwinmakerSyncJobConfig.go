package iottwinmakersyncjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IottwinmakerSyncJobConfig struct {
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
	// The IAM Role that execute SyncJob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iottwinmaker_sync_job#sync_role IottwinmakerSyncJob#sync_role}
	SyncRole *string `field:"required" json:"syncRole" yaml:"syncRole"`
	// The source of the SyncJob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iottwinmaker_sync_job#sync_source IottwinmakerSyncJob#sync_source}
	SyncSource *string `field:"required" json:"syncSource" yaml:"syncSource"`
	// The ID of the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iottwinmaker_sync_job#workspace_id IottwinmakerSyncJob#workspace_id}
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iottwinmaker_sync_job#tags IottwinmakerSyncJob#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

