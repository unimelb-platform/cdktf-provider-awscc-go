package ekscluster


type EksClusterLogging struct {
	// The cluster control plane logging configuration for your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_cluster#cluster_logging EksCluster#cluster_logging}
	ClusterLogging *EksClusterLoggingClusterLogging `field:"optional" json:"clusterLogging" yaml:"clusterLogging"`
}

