package backupbackupselection


type BackupBackupSelectionBackupSelectionConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#string_equals BackupBackupSelection#string_equals}.
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#string_like BackupBackupSelection#string_like}.
	StringLike interface{} `field:"optional" json:"stringLike" yaml:"stringLike"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#string_not_equals BackupBackupSelection#string_not_equals}.
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_selection#string_not_like BackupBackupSelection#string_not_like}.
	StringNotLike interface{} `field:"optional" json:"stringNotLike" yaml:"stringNotLike"`
}

