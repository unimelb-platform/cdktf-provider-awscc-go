package mskserverlesscluster


type MskServerlessClusterClientAuthenticationSasl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_serverless_cluster#iam MskServerlessCluster#iam}.
	Iam *MskServerlessClusterClientAuthenticationSaslIam `field:"required" json:"iam" yaml:"iam"`
}

