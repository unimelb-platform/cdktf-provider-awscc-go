package backupbackupselection


type BackupBackupSelectionBackupSelectionListOfTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#condition_key BackupBackupSelection#condition_key}.
	ConditionKey *string `field:"required" json:"conditionKey" yaml:"conditionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#condition_type BackupBackupSelection#condition_type}.
	ConditionType *string `field:"required" json:"conditionType" yaml:"conditionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#condition_value BackupBackupSelection#condition_value}.
	ConditionValue *string `field:"required" json:"conditionValue" yaml:"conditionValue"`
}

