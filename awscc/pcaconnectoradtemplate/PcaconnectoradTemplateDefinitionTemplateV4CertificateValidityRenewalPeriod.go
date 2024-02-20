package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4CertificateValidityRenewalPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#period PcaconnectoradTemplate#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#period_type PcaconnectoradTemplate#period_type}.
	PeriodType *string `field:"required" json:"periodType" yaml:"periodType"`
}

