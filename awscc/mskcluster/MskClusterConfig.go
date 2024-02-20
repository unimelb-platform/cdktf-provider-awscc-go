package mskcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MskClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#broker_node_group_info MskCluster#broker_node_group_info}.
	BrokerNodeGroupInfo *MskClusterBrokerNodeGroupInfo `field:"required" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#cluster_name MskCluster#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#kafka_version MskCluster#kafka_version}.
	KafkaVersion *string `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#number_of_broker_nodes MskCluster#number_of_broker_nodes}.
	NumberOfBrokerNodes *float64 `field:"required" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#client_authentication MskCluster#client_authentication}.
	ClientAuthentication *MskClusterClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#configuration_info MskCluster#configuration_info}.
	ConfigurationInfo *MskClusterConfigurationInfo `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// The current version of the MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#current_version MskCluster#current_version}
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#encryption_info MskCluster#encryption_info}.
	EncryptionInfo *MskClusterEncryptionInfo `field:"optional" json:"encryptionInfo" yaml:"encryptionInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#enhanced_monitoring MskCluster#enhanced_monitoring}.
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#logging_info MskCluster#logging_info}.
	LoggingInfo *MskClusterLoggingInfo `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#open_monitoring MskCluster#open_monitoring}.
	OpenMonitoring *MskClusterOpenMonitoring `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#storage_mode MskCluster#storage_mode}.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#tags MskCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

