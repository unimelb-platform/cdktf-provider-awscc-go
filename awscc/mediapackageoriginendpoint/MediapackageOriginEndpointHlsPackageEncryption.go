package mediapackageoriginendpoint


type MediapackageOriginEndpointHlsPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#speke_key_provider MediapackageOriginEndpoint#speke_key_provider}
	SpekeKeyProvider *MediapackageOriginEndpointHlsPackageEncryptionSpekeKeyProvider `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
	// A constant initialization vector for encryption (optional). When not specified the initialization vector will be periodically rotated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#constant_initialization_vector MediapackageOriginEndpoint#constant_initialization_vector}
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The encryption method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption_method MediapackageOriginEndpoint#encryption_method}
	EncryptionMethod *string `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// Interval (in seconds) between each encryption key rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#key_rotation_interval_seconds MediapackageOriginEndpoint#key_rotation_interval_seconds}
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
	// When enabled, the EXT-X-KEY tag will be repeated in output manifests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#repeat_ext_x_key MediapackageOriginEndpoint#repeat_ext_x_key}
	RepeatExtXKey interface{} `field:"optional" json:"repeatExtXKey" yaml:"repeatExtXKey"`
}

