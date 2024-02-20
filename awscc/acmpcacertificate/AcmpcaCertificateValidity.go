package acmpcacertificate


type AcmpcaCertificateValidity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#type AcmpcaCertificate#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#value AcmpcaCertificate#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

