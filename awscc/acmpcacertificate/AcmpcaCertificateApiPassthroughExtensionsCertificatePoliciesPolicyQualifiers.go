package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#policy_qualifier_id AcmpcaCertificate#policy_qualifier_id}.
	PolicyQualifierId *string `field:"required" json:"policyQualifierId" yaml:"policyQualifierId"`
	// Structure that contains a X.509 policy qualifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#qualifier AcmpcaCertificate#qualifier}
	Qualifier *AcmpcaCertificateApiPassthroughExtensionsCertificatePoliciesPolicyQualifiersQualifier `field:"required" json:"qualifier" yaml:"qualifier"`
}

