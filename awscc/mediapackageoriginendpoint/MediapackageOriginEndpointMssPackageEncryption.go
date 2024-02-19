package mediapackageoriginendpoint


type MediapackageOriginEndpointMssPackageEncryption struct {
	// A configuration for accessing an external Secure Packager and Encoder Key Exchange (SPEKE) service that will provide encryption keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#speke_key_provider MediapackageOriginEndpoint#speke_key_provider}
	SpekeKeyProvider *MediapackageOriginEndpointMssPackageEncryptionSpekeKeyProvider `field:"required" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}

