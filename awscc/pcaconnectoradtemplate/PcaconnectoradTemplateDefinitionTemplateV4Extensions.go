package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4Extensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#key_usage PcaconnectoradTemplate#key_usage}.
	KeyUsage *PcaconnectoradTemplateDefinitionTemplateV4ExtensionsKeyUsage `field:"required" json:"keyUsage" yaml:"keyUsage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#application_policies PcaconnectoradTemplate#application_policies}.
	ApplicationPolicies *PcaconnectoradTemplateDefinitionTemplateV4ExtensionsApplicationPolicies `field:"optional" json:"applicationPolicies" yaml:"applicationPolicies"`
}

