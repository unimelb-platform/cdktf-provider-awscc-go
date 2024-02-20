package backupbackupplan


type BackupBackupPlanBackupPlanAdvancedBackupSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#backup_options BackupBackupPlan#backup_options}.
	BackupOptions *string `field:"required" json:"backupOptions" yaml:"backupOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#resource_type BackupBackupPlan#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

