package mskreplicator


type MskReplicatorKafkaClusters struct {
	// Details of an Amazon MSK cluster. Exactly one of AmazonMskCluster is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#amazon_msk_cluster MskReplicator#amazon_msk_cluster}
	AmazonMskCluster *MskReplicatorKafkaClustersAmazonMskCluster `field:"required" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#vpc_config MskReplicator#vpc_config}
	VpcConfig *MskReplicatorKafkaClustersVpcConfig `field:"required" json:"vpcConfig" yaml:"vpcConfig"`
}

