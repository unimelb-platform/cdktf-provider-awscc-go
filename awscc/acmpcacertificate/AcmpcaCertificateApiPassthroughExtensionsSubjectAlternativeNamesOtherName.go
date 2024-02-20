package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesOtherName struct {
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#type_id AcmpcaCertificate#type_id}
	TypeId *string `field:"required" json:"typeId" yaml:"typeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

