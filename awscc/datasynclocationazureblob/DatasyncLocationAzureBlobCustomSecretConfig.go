package datasynclocationazureblob


type DatasyncLocationAzureBlobCustomSecretConfig struct {
	// Specifies the ARN for the AWS Identity and Access Management role that DataSync uses to access the secret specified for SecretArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_azure_blob#secret_access_role_arn DatasyncLocationAzureBlob#secret_access_role_arn}
	SecretAccessRoleArn *string `field:"optional" json:"secretAccessRoleArn" yaml:"secretAccessRoleArn"`
	// Specifies the ARN for a customer created AWS Secrets Manager secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_azure_blob#secret_arn DatasyncLocationAzureBlob#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

