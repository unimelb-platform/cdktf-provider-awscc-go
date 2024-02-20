package s3storagelens


type S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevel struct {
	// Selection criteria for Storage Lens Group level metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#storage_lens_group_selection_criteria S3StorageLens#storage_lens_group_selection_criteria}
	StorageLensGroupSelectionCriteria *S3StorageLensStorageLensConfigurationAccountLevelStorageLensGroupLevelStorageLensGroupSelectionCriteria `field:"optional" json:"storageLensGroupSelectionCriteria" yaml:"storageLensGroupSelectionCriteria"`
}

