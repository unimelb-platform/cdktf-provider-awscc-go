package acmpcacertificate


type AcmpcaCertificateApiPassthrough struct {
	// Structure that contains X.500 extensions for a Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#extensions AcmpcaCertificate#extensions}
	Extensions *AcmpcaCertificateApiPassthroughExtensions `field:"optional" json:"extensions" yaml:"extensions"`
	// Structure that contains X.500 distinguished name information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#subject AcmpcaCertificate#subject}
	Subject *AcmpcaCertificateApiPassthroughSubject `field:"optional" json:"subject" yaml:"subject"`
}

