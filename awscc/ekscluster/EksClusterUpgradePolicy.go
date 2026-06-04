package ekscluster


type EksClusterUpgradePolicy struct {
	// Specify the support type for your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#support_type EksCluster#support_type}
	SupportType *string `field:"optional" json:"supportType" yaml:"supportType"`
}

