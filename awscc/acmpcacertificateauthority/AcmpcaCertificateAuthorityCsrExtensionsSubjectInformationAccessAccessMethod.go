package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod struct {
	// Pre-defined enum string for X.509 AccessMethod ObjectIdentifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#access_method_type AcmpcaCertificateAuthority#access_method_type}
	AccessMethodType *string `field:"optional" json:"accessMethodType" yaml:"accessMethodType"`
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#custom_object_identifier AcmpcaCertificateAuthority#custom_object_identifier}
	CustomObjectIdentifier *string `field:"optional" json:"customObjectIdentifier" yaml:"customObjectIdentifier"`
}

