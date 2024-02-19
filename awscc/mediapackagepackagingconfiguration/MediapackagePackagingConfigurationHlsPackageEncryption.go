package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationHlsPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#speke_key_provider MediapackagePackagingConfiguration#speke_key_provider}
	SpekeKeyProvider *MediapackagePackagingConfigurationHlsPackageEncryptionSpekeKeyProvider `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// An HTTP Live Streaming (HLS) encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#constant_initialization_vector MediapackagePackagingConfiguration#constant_initialization_vector}
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The encryption method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#encryption_method MediapackagePackagingConfiguration#encryption_method}
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
}

