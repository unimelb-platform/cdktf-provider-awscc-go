package supportappslackworkspaceconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SupportappSlackWorkspaceConfigurationConfig struct {
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
	// The team ID in Slack, which uniquely identifies a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_workspace_configuration#team_id SupportappSlackWorkspaceConfiguration#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_workspace_configuration#version_id SupportappSlackWorkspaceConfiguration#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

