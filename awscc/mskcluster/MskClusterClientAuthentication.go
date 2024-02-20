package mskcluster


type MskClusterClientAuthentication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#sasl MskCluster#sasl}.
	Sasl *MskClusterClientAuthenticationSasl `field:"optional" json:"sasl" yaml:"sasl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#tls MskCluster#tls}.
	Tls *MskClusterClientAuthenticationTls `field:"optional" json:"tls" yaml:"tls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#unauthenticated MskCluster#unauthenticated}.
	Unauthenticated *MskClusterClientAuthenticationUnauthenticated `field:"optional" json:"unauthenticated" yaml:"unauthenticated"`
}

