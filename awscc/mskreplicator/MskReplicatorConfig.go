package mskreplicator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskReplicatorConfig struct {
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
	// Specifies a list of Kafka clusters which are targets of the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#kafka_clusters MskReplicator#kafka_clusters}
	KafkaClusters interface{} `field:"required" json:"kafkaClusters" yaml:"kafkaClusters"`
	// A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#replication_info_list MskReplicator#replication_info_list}
	ReplicationInfoList interface{} `field:"required" json:"replicationInfoList" yaml:"replicationInfoList"`
	// The name of the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#replicator_name MskReplicator#replicator_name}
	ReplicatorName *string `field:"required" json:"replicatorName" yaml:"replicatorName"`
	// The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#service_execution_role_arn MskReplicator#service_execution_role_arn}
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// The current version of the MSK replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#current_version MskReplicator#current_version}
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// A summary description of the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#description MskReplicator#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#tags MskReplicator#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

