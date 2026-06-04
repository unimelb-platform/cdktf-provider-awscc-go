package backuplogicallyairgappedbackupvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupLogicallyAirGappedBackupVaultConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault#backup_vault_name BackupLogicallyAirGappedBackupVault#backup_vault_name}.
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault#max_retention_days BackupLogicallyAirGappedBackupVault#max_retention_days}.
	MaxRetentionDays *float64 `field:"required" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault#min_retention_days BackupLogicallyAirGappedBackupVault#min_retention_days}.
	MinRetentionDays *float64 `field:"required" json:"minRetentionDays" yaml:"minRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault#access_policy BackupLogicallyAirGappedBackupVault#access_policy}.
	AccessPolicy *string `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault#backup_vault_tags BackupLogicallyAirGappedBackupVault#backup_vault_tags}.
	BackupVaultTags *map[string]*string `field:"optional" json:"backupVaultTags" yaml:"backupVaultTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/backup_logically_air_gapped_backup_vault#notifications BackupLogicallyAirGappedBackupVault#notifications}.
	Notifications *BackupLogicallyAirGappedBackupVaultNotifications `field:"optional" json:"notifications" yaml:"notifications"`
}

