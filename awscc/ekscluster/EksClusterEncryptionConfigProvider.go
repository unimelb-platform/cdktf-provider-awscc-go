package ekscluster


type EksClusterEncryptionConfigProvider struct {
	// Amazon Resource Name (ARN) or alias of the KMS key.
	//
	// The KMS key must be symmetric, created in the same region as the cluster, and if the KMS key was created in a different account, the user must have access to the KMS key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_cluster#key_arn EksCluster#key_arn}
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

