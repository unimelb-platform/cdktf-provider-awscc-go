package backuprestoretestingselection


type BackupRestoreTestingSelectionProtectedResourceConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_restore_testing_selection#string_equals BackupRestoreTestingSelection#string_equals}.
	StringEquals interface{} `field:"optional" json:"stringEquals" yaml:"stringEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_restore_testing_selection#string_not_equals BackupRestoreTestingSelection#string_not_equals}.
	StringNotEquals interface{} `field:"optional" json:"stringNotEquals" yaml:"stringNotEquals"`
}

