package secretsmanagersecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecretsmanagerSecretConfig struct {
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
	// (Optional) Specifies a user-provided description of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#description SecretsmanagerSecret#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// (Optional) Specifies text data that you want to encrypt and store in this new version of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#generate_secret_string SecretsmanagerSecret#generate_secret_string}
	GenerateSecretString *SecretsmanagerSecretGenerateSecretString `field:"optional" json:"generateSecretString" yaml:"generateSecretString"`
	// (Optional) Specifies the ARN, Key ID, or alias of the AWS KMS customer master key (CMK) used to encrypt the SecretString.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#kms_key_id SecretsmanagerSecret#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The friendly name of the secret. You can use forward slashes in the name to represent a path hierarchy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#name SecretsmanagerSecret#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// (Optional) A list of ReplicaRegion objects.
	//
	// The ReplicaRegion type consists of a Region (required) and the KmsKeyId which can be an ARN, Key ID, or Alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#replica_regions SecretsmanagerSecret#replica_regions}
	ReplicaRegions interface{} `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// (Optional) Specifies text data that you want to encrypt and store in this new version of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#secret_string SecretsmanagerSecret#secret_string}
	SecretString *string `field:"optional" json:"secretString" yaml:"secretString"`
	// The list of user-defined tags associated with the secret.
	//
	// Use tags to manage your AWS resources. For additional information about tags, see TagResource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#tags SecretsmanagerSecret#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

