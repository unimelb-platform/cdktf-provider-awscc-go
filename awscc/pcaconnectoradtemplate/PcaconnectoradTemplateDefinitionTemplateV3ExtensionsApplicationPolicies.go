package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV3ExtensionsApplicationPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#policies PcaconnectoradTemplate#policies}.
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#critical PcaconnectoradTemplate#critical}.
	Critical interface{} `field:"optional" json:"critical" yaml:"critical"`
}

