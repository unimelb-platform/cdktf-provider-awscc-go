package mskcluster


type MskClusterClientAuthenticationSasl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#iam MskCluster#iam}.
	Iam *MskClusterClientAuthenticationSaslIam `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#scram MskCluster#scram}.
	Scram *MskClusterClientAuthenticationSaslScram `field:"optional" json:"scram" yaml:"scram"`
}

