package mediapackageoriginendpoint


type MediapackageOriginEndpointDashPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#speke_key_provider MediapackageOriginEndpoint#speke_key_provider}
	SpekeKeyProvider *MediapackageOriginEndpointDashPackageEncryptionSpekeKeyProvider `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// Time (in seconds) between each encryption key rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#key_rotation_interval_seconds MediapackageOriginEndpoint#key_rotation_interval_seconds}
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

