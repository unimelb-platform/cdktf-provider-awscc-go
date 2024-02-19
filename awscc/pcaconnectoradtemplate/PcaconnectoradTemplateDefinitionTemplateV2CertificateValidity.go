package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#renewal_period PcaconnectoradTemplate#renewal_period}.
	RenewalPeriod *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityRenewalPeriod `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#validity_period PcaconnectoradTemplate#validity_period}.
	ValidityPeriod *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityValidityPeriod `field:"required" json:"validityPeriod" yaml:"validityPeriod"`
}

