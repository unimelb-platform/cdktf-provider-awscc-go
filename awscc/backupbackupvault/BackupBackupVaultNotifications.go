package backupbackupvault


type BackupBackupVaultNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#backup_vault_events BackupBackupVault#backup_vault_events}.
	BackupVaultEvents *[]*string `field:"required" json:"backupVaultEvents" yaml:"backupVaultEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#sns_topic_arn BackupBackupVault#sns_topic_arn}.
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

