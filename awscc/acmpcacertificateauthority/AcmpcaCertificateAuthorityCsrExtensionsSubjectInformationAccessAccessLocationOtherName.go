package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOtherName struct {
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#type_id AcmpcaCertificateAuthority#type_id}
	TypeId *string `field:"required" json:"typeId" yaml:"typeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#value AcmpcaCertificateAuthority#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

