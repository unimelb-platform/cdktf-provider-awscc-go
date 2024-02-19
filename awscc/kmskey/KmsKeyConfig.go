package kmskey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsKeyConfig struct {
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
	// Skips ("bypasses") the key policy lockout safety check. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#bypass_policy_lockout_safety_check KmsKey#bypass_policy_lockout_safety_check}
	BypassPolicyLockoutSafetyCheck interface{} `field:"optional" json:"bypassPolicyLockoutSafetyCheck" yaml:"bypassPolicyLockoutSafetyCheck"`
	// A description of the AWS KMS key.
	//
	// Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#description KmsKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#enabled KmsKey#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Enables automatic rotation of the key material for the specified AWS KMS key.
	//
	// By default, automation key rotation is not enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#enable_key_rotation KmsKey#enable_key_rotation}
	EnableKeyRotation interface{} `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#key_policy KmsKey#key_policy}
	KeyPolicy *string `field:"optional" json:"keyPolicy" yaml:"keyPolicy"`
	// Specifies the type of AWS KMS key to create.
	//
	// The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#key_spec KmsKey#key_spec}
	KeySpec *string `field:"optional" json:"keySpec" yaml:"keySpec"`
	// Determines the cryptographic operations for which you can use the AWS KMS key.
	//
	// The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#key_usage KmsKey#key_usage}
	KeyUsage *string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Specifies whether the AWS KMS key should be Multi-Region.
	//
	// You can't change the MultiRegion value after the AWS KMS key is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#multi_region KmsKey#multi_region}
	MultiRegion interface{} `field:"optional" json:"multiRegion" yaml:"multiRegion"`
	// The source of the key material for the KMS key.
	//
	// You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#origin KmsKey#origin}
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack.
	//
	// Enter a value between 7 and 30 days. The default value is 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#pending_window_in_days KmsKey#pending_window_in_days}
	PendingWindowInDays *float64 `field:"optional" json:"pendingWindowInDays" yaml:"pendingWindowInDays"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kms_key#tags KmsKey#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

