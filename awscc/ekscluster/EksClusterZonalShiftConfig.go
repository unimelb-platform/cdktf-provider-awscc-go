package ekscluster


type EksClusterZonalShiftConfig struct {
	// Set this value to true to enable zonal shift for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#enabled EksCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

