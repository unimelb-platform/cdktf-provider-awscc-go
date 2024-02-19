package efsfilesystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EfsFileSystemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#availability_zone_name EfsFileSystem#availability_zone_name}.
	AvailabilityZoneName *string `field:"optional" json:"availabilityZoneName" yaml:"availabilityZoneName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#backup_policy EfsFileSystem#backup_policy}.
	BackupPolicy *EfsFileSystemBackupPolicy `field:"optional" json:"backupPolicy" yaml:"backupPolicy"`
	// Whether to bypass the FileSystemPolicy lockout safety check.
	//
	// The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#bypass_policy_lockout_safety_check EfsFileSystem#bypass_policy_lockout_safety_check}
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#encrypted EfsFileSystem#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#file_system_policy EfsFileSystem#file_system_policy}.
	FileSystemPolicy *string `field:"optional" json:"fileSystemPolicy" yaml:"fileSystemPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#file_system_protection EfsFileSystem#file_system_protection}.
	FileSystemProtection *EfsFileSystemFileSystemProtection `field:"optional" json:"fileSystemProtection" yaml:"fileSystemProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#file_system_tags EfsFileSystem#file_system_tags}.
	FileSystemTags interface{} `field:"optional" json:"fileSystemTags" yaml:"fileSystemTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#kms_key_id EfsFileSystem#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#lifecycle_policies EfsFileSystem#lifecycle_policies}.
	LifecyclePolicies interface{} `field:"optional" json:"lifecyclePolicies" yaml:"lifecyclePolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#performance_mode EfsFileSystem#performance_mode}.
	PerformanceMode *string `field:"optional" json:"performanceMode" yaml:"performanceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#provisioned_throughput_in_mibps EfsFileSystem#provisioned_throughput_in_mibps}.
	ProvisionedThroughputInMibps *float64 `field:"optional" json:"provisionedThroughputInMibps" yaml:"provisionedThroughputInMibps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#replication_configuration EfsFileSystem#replication_configuration}.
	ReplicationConfiguration *EfsFileSystemReplicationConfiguration `field:"optional" json:"replicationConfiguration" yaml:"replicationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#throughput_mode EfsFileSystem#throughput_mode}.
	ThroughputMode *string `field:"optional" json:"throughputMode" yaml:"throughputMode"`
}

