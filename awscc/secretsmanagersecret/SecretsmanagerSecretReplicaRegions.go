package secretsmanagersecret


type SecretsmanagerSecretReplicaRegions struct {
	// (Optional) A string that represents a Region, for example "us-east-1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#region SecretsmanagerSecret#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN, key ID, or alias of the KMS key to encrypt the secret.
	//
	// If you don't include this field, Secrets Manager uses aws/secretsmanager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/secretsmanager_secret#kms_key_id SecretsmanagerSecret#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

