package backupbackupplan


type BackupBackupPlanBackupPlanBackupPlanRuleCopyActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#destination_backup_vault_arn BackupBackupPlan#destination_backup_vault_arn}.
	DestinationBackupVaultArn *string `field:"required" json:"destinationBackupVaultArn" yaml:"destinationBackupVaultArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#lifecycle BackupBackupPlan#lifecycle}.
	Lifecycle *BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

