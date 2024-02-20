package mskcluster


type MskClusterBrokerNodeGroupInfoConnectivityInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#public_access MskCluster#public_access}.
	PublicAccess *MskClusterBrokerNodeGroupInfoConnectivityInfoPublicAccess `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#vpc_connectivity MskCluster#vpc_connectivity}.
	VpcConnectivity *MskClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivity `field:"optional" json:"vpcConnectivity" yaml:"vpcConnectivity"`
}

