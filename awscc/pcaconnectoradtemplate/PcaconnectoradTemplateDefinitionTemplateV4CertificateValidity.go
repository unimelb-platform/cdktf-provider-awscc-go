package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4CertificateValidity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#renewal_period PcaconnectoradTemplate#renewal_period}.
	RenewalPeriod *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidityRenewalPeriod `field:"required" json:"renewalPeriod" yaml:"renewalPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#validity_period PcaconnectoradTemplate#validity_period}.
	ValidityPeriod *PcaconnectoradTemplateDefinitionTemplateV4CertificateValidityValidityPeriod `field:"required" json:"validityPeriod" yaml:"validityPeriod"`
}

