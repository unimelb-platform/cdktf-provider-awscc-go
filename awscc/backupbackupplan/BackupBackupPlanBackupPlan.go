package backupbackupplan


type BackupBackupPlanBackupPlan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#backup_plan_name BackupBackupPlan#backup_plan_name}.
	BackupPlanName *string `field:"required" json:"backupPlanName" yaml:"backupPlanName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#backup_plan_rule BackupBackupPlan#backup_plan_rule}.
	BackupPlanRule interface{} `field:"required" json:"backupPlanRule" yaml:"backupPlanRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#advanced_backup_settings BackupBackupPlan#advanced_backup_settings}.
	AdvancedBackupSettings interface{} `field:"optional" json:"advancedBackupSettings" yaml:"advancedBackupSettings"`
}

