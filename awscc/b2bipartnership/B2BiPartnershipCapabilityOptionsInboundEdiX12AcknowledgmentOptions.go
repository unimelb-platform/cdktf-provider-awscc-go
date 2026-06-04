package b2bipartnership


type B2BiPartnershipCapabilityOptionsInboundEdiX12AcknowledgmentOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_partnership#functional_acknowledgment B2BiPartnership#functional_acknowledgment}.
	FunctionalAcknowledgment *string `field:"optional" json:"functionalAcknowledgment" yaml:"functionalAcknowledgment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/b2bi_partnership#technical_acknowledgment B2BiPartnership#technical_acknowledgment}.
	TechnicalAcknowledgment *string `field:"optional" json:"technicalAcknowledgment" yaml:"technicalAcknowledgment"`
}

