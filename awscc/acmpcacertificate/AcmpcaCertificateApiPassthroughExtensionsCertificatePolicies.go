package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsCertificatePolicies struct {
	// String that contains X.509 ObjectIdentifier information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#cert_policy_id AcmpcaCertificate#cert_policy_id}
	CertPolicyId *string `field:"required" json:"certPolicyId" yaml:"certPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#policy_qualifiers AcmpcaCertificate#policy_qualifiers}.
	PolicyQualifiers interface{} `field:"optional" json:"policyQualifiers" yaml:"policyQualifiers"`
}

