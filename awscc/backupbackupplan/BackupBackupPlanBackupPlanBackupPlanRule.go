package backupbackupplan


type BackupBackupPlanBackupPlanBackupPlanRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#rule_name BackupBackupPlan#rule_name}.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#target_backup_vault BackupBackupPlan#target_backup_vault}.
	TargetBackupVault *string `field:"required" json:"targetBackupVault" yaml:"targetBackupVault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#completion_window_minutes BackupBackupPlan#completion_window_minutes}.
	CompletionWindowMinutes *float64 `field:"optional" json:"completionWindowMinutes" yaml:"completionWindowMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#copy_actions BackupBackupPlan#copy_actions}.
	CopyActions interface{} `field:"optional" json:"copyActions" yaml:"copyActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#enable_continuous_backup BackupBackupPlan#enable_continuous_backup}.
	EnableContinuousBackup interface{} `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#lifecycle BackupBackupPlan#lifecycle}.
	Lifecycle *BackupBackupPlanBackupPlanBackupPlanRuleLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#recovery_point_tags BackupBackupPlan#recovery_point_tags}.
	RecoveryPointTags *map[string]*string `field:"optional" json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#schedule_expression BackupBackupPlan#schedule_expression}.
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#schedule_expression_timezone BackupBackupPlan#schedule_expression_timezone}.
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_plan#start_window_minutes BackupBackupPlan#start_window_minutes}.
	StartWindowMinutes *float64 `field:"optional" json:"startWindowMinutes" yaml:"startWindowMinutes"`
}

