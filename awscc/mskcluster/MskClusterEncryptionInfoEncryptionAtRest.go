package mskcluster


type MskClusterEncryptionInfoEncryptionAtRest struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#data_volume_kms_key_id MskCluster#data_volume_kms_key_id}.
	DataVolumeKmsKeyId *string `field:"required" json:"dataVolumeKmsKeyId" yaml:"dataVolumeKmsKeyId"`
}

