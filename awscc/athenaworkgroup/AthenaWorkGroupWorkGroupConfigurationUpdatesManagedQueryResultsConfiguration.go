package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/athena_work_group#enabled AthenaWorkGroup#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates the encryption configuration for Athena Managed Storage.
	//
	// If not setting this field, Managed Storage will encrypt the query results with Athena's encryption key
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/athena_work_group#encryption_configuration AthenaWorkGroup#encryption_configuration}
	EncryptionConfiguration *AthenaWorkGroupWorkGroupConfigurationUpdatesManagedQueryResultsConfigurationEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

