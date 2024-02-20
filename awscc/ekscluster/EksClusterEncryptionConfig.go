package ekscluster


type EksClusterEncryptionConfig struct {
	// The encryption provider for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_cluster#provider EksCluster#provider}
	Provider *EksClusterEncryptionConfigProvider `field:"optional" json:"provider" yaml:"provider"`
	// Specifies the resources to be encrypted. The only supported value is "secrets".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_cluster#resources EksCluster#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

