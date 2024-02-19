package ssmincidentsreplicationset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmincidentsReplicationSetConfig struct {
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
	// The ReplicationSet configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_replication_set#regions SsmincidentsReplicationSet#regions}
	Regions interface{} `field:"required" json:"regions" yaml:"regions"`
	// Configures the ReplicationSet deletion protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_replication_set#deletion_protected SsmincidentsReplicationSet#deletion_protected}
	DeletionProtected interface{} `field:"optional" json:"deletionProtected" yaml:"deletionProtected"`
	// The tags to apply to the replication set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ssmincidents_replication_set#tags SsmincidentsReplicationSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

