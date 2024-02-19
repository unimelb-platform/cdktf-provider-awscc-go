package acmpcacertificate


type AcmpcaCertificateApiPassthroughSubjectCustomAttributes struct {
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#object_identifier AcmpcaCertificate#object_identifier}
	ObjectIdentifier *string `field:"required" json:"objectIdentifier" yaml:"objectIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

