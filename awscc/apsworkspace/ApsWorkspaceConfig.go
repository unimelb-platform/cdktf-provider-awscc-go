package apsworkspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApsWorkspaceConfig struct {
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
	// The AMP Workspace alert manager definition data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_workspace#alert_manager_definition ApsWorkspace#alert_manager_definition}
	AlertManagerDefinition *string `field:"optional" json:"alertManagerDefinition" yaml:"alertManagerDefinition"`
	// AMP Workspace alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_workspace#alias ApsWorkspace#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// KMS Key ARN used to encrypt and decrypt AMP workspace data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_workspace#kms_key_arn ApsWorkspace#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Logging configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_workspace#logging_configuration ApsWorkspace#logging_configuration}
	LoggingConfiguration *ApsWorkspaceLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/aps_workspace#tags ApsWorkspace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

