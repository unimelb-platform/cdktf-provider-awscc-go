package mskreplicator


type MskReplicatorKafkaClustersAmazonMskCluster struct {
	// The ARN of an Amazon MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#msk_cluster_arn MskReplicator#msk_cluster_arn}
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
}

