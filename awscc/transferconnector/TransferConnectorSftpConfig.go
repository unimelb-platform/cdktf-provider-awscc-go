package transferconnector


type TransferConnectorSftpConfig struct {
	// List of public host keys, for the external server to which you are connecting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#trusted_host_keys TransferConnector#trusted_host_keys}
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// ARN or name of the secret in AWS Secrets Manager which contains the SFTP user's private keys or passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_connector#user_secret_id TransferConnector#user_secret_id}
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}

