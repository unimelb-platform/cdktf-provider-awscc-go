package secretsmanagersecrettargetattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerSecretTargetAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_secret_target_attachment#secret_id SecretsmanagerSecretTargetAttachment#secret_id}.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_secret_target_attachment#target_id SecretsmanagerSecretTargetAttachment#target_id}.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_secret_target_attachment#target_type SecretsmanagerSecretTargetAttachment#target_type}.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

