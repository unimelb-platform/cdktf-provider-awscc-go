package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationMssPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackage_packaging_configuration#speke_key_provider MediapackagePackagingConfiguration#speke_key_provider}
	SpekeKeyProvider *MediapackagePackagingConfigurationMssPackageEncryptionSpekeKeyProvider `field:"optional" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

