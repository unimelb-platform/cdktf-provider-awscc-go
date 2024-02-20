package mskcluster


type MskClusterBrokerNodeGroupInfoStorageInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#ebs_storage_info MskCluster#ebs_storage_info}.
	EbsStorageInfo *MskClusterBrokerNodeGroupInfoStorageInfoEbsStorageInfo `field:"optional" json:"ebsStorageInfo" yaml:"ebsStorageInfo"`
}

