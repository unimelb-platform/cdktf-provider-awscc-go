package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#certificate_policies AcmpcaCertificate#certificate_policies}.
	CertificatePolicies interface{} `field:"optional" json:"certificatePolicies" yaml:"certificatePolicies"`
	// Array of X.509 extensions for a certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#custom_extensions AcmpcaCertificate#custom_extensions}
	CustomExtensions interface{} `field:"optional" json:"customExtensions" yaml:"customExtensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#extended_key_usage AcmpcaCertificate#extended_key_usage}.
	ExtendedKeyUsage interface{} `field:"optional" json:"extendedKeyUsage" yaml:"extendedKeyUsage"`
	// Structure that contains X.509 KeyUsage information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#key_usage AcmpcaCertificate#key_usage}
	KeyUsage *AcmpcaCertificateApiPassthroughExtensionsKeyUsage `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#subject_alternative_names AcmpcaCertificate#subject_alternative_names}.
	SubjectAlternativeNames interface{} `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}

