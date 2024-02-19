package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsExtendedKeyUsage struct {
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#extended_key_usage_object_identifier AcmpcaCertificate#extended_key_usage_object_identifier}
	ExtendedKeyUsageObjectIdentifier *string `field:"optional" json:"extendedKeyUsageObjectIdentifier" yaml:"extendedKeyUsageObjectIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#extended_key_usage_type AcmpcaCertificate#extended_key_usage_type}.
	ExtendedKeyUsageType *string `field:"optional" json:"extendedKeyUsageType" yaml:"extendedKeyUsageType"`
}

