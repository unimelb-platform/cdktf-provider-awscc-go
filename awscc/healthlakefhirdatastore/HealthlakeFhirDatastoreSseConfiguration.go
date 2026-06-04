package healthlakefhirdatastore


type HealthlakeFhirDatastoreSseConfiguration struct {
	// The customer-managed-key (CMK) used when creating a Data Store.
	//
	// If a customer owned key is not specified, an AWS owned key will be used for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/healthlake_fhir_datastore#kms_encryption_config HealthlakeFhirDatastore#kms_encryption_config}
	KmsEncryptionConfig *HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig `field:"optional" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
}

