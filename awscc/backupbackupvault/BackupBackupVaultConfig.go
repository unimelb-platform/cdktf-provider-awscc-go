package backupbackupvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupBackupVaultConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#backup_vault_name BackupBackupVault#backup_vault_name}.
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#access_policy BackupBackupVault#access_policy}.
	AccessPolicy *string `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#backup_vault_tags BackupBackupVault#backup_vault_tags}.
	BackupVaultTags *map[string]*string `field:"optional" json:"backupVaultTags" yaml:"backupVaultTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#encryption_key_arn BackupBackupVault#encryption_key_arn}.
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#lock_configuration BackupBackupVault#lock_configuration}.
	LockConfiguration *BackupBackupVaultLockConfiguration `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/backup_backup_vault#notifications BackupBackupVault#notifications}.
	Notifications *BackupBackupVaultNotifications `field:"optional" json:"notifications" yaml:"notifications"`
}

