package systemsmanagersapapplication


type SystemsmanagersapApplicationCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#credential_type SystemsmanagersapApplication#credential_type}.
	CredentialType *string `field:"optional" json:"credentialType" yaml:"credentialType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#database_name SystemsmanagersapApplication#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#secret_id SystemsmanagersapApplication#secret_id}.
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
}

