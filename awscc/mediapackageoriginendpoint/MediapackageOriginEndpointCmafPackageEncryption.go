package mediapackageoriginendpoint


type MediapackageOriginEndpointCmafPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#speke_key_provider MediapackageOriginEndpoint#speke_key_provider}
	SpekeKeyProvider *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// An optional 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting blocks.
	//
	// If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#constant_initialization_vector MediapackageOriginEndpoint#constant_initialization_vector}
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The encryption method used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption_method MediapackageOriginEndpoint#encryption_method}
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Time (in seconds) between each encryption key rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#key_rotation_interval_seconds MediapackageOriginEndpoint#key_rotation_interval_seconds}
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
}

