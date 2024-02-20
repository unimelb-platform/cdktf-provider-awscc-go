package acmpcacertificate


type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesEdiPartyName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#name_assigner AcmpcaCertificate#name_assigner}.
	NameAssigner *string `field:"required" json:"nameAssigner" yaml:"nameAssigner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate#party_name AcmpcaCertificate#party_name}.
	PartyName *string `field:"required" json:"partyName" yaml:"partyName"`
}

