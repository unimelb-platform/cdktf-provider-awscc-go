package acmpcacertificateauthority


type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationEdiPartyName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#name_assigner AcmpcaCertificateAuthority#name_assigner}.
	NameAssigner *string `field:"required" json:"nameAssigner" yaml:"nameAssigner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/acmpca_certificate_authority#party_name AcmpcaCertificateAuthority#party_name}.
	PartyName *string `field:"required" json:"partyName" yaml:"partyName"`
}

