package secretsmanagerresourcepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerResourcePolicyConfig struct {
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
	// A JSON-formatted string for an AWS resource-based policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_resource_policy#resource_policy SecretsmanagerResourcePolicy#resource_policy}
	ResourcePolicy *string `field:"required" json:"resourcePolicy" yaml:"resourcePolicy"`
	// The ARN or name of the secret to attach the resource-based policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_resource_policy#secret_id SecretsmanagerResourcePolicy#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Specifies whether to block resource-based policies that allow broad access to the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/secretsmanager_resource_policy#block_public_policy SecretsmanagerResourcePolicy#block_public_policy}
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
}

