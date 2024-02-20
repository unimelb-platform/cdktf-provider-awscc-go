package dmsreplicationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsReplicationConfigConfig struct {
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
	// Configuration parameters for provisioning a AWS DMS Serverless replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#compute_config DmsReplicationConfig#compute_config}
	ComputeConfig *DmsReplicationConfigComputeConfig `field:"optional" json:"computeConfig" yaml:"computeConfig"`
	// The Amazon Resource Name (ARN) of the Replication Config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#replication_config_arn DmsReplicationConfig#replication_config_arn}
	ReplicationConfigArn *string `field:"optional" json:"replicationConfigArn" yaml:"replicationConfigArn"`
	// A unique identifier of replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#replication_config_identifier DmsReplicationConfig#replication_config_identifier}
	ReplicationConfigIdentifier *string `field:"optional" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// JSON settings for Servereless replications that are provisioned using this replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#replication_settings DmsReplicationConfig#replication_settings}
	ReplicationSettings *string `field:"optional" json:"replicationSettings" yaml:"replicationSettings"`
	// The type of AWS DMS Serverless replication to provision using this replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#replication_type DmsReplicationConfig#replication_type}
	ReplicationType *string `field:"optional" json:"replicationType" yaml:"replicationType"`
	// A unique value or name that you get set for a given resource that can be used to construct an Amazon Resource Name (ARN) for that resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#resource_identifier DmsReplicationConfig#resource_identifier}
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// The Amazon Resource Name (ARN) of the source endpoint for this AWS DMS Serverless replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#source_endpoint_arn DmsReplicationConfig#source_endpoint_arn}
	SourceEndpointArn *string `field:"optional" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// JSON settings for specifying supplemental data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#supplemental_settings DmsReplicationConfig#supplemental_settings}
	SupplementalSettings *string `field:"optional" json:"supplementalSettings" yaml:"supplementalSettings"`
	// JSON table mappings for AWS DMS Serverless replications that are provisioned using this replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#table_mappings DmsReplicationConfig#table_mappings}
	TableMappings *string `field:"optional" json:"tableMappings" yaml:"tableMappings"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#tags DmsReplicationConfig#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the target endpoint for this AWS DMS Serverless replication configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_replication_config#target_endpoint_arn DmsReplicationConfig#target_endpoint_arn}
	TargetEndpointArn *string `field:"optional" json:"targetEndpointArn" yaml:"targetEndpointArn"`
}

