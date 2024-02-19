package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV3PrivateKeyFlags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#client_version PcaconnectoradTemplate#client_version}.
	ClientVersion *string `field:"required" json:"clientVersion" yaml:"clientVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#exportable_key PcaconnectoradTemplate#exportable_key}.
	ExportableKey interface{} `field:"optional" json:"exportableKey" yaml:"exportableKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#require_alternate_signature_algorithm PcaconnectoradTemplate#require_alternate_signature_algorithm}.
	RequireAlternateSignatureAlgorithm interface{} `field:"optional" json:"requireAlternateSignatureAlgorithm" yaml:"requireAlternateSignatureAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#strong_key_protection_required PcaconnectoradTemplate#strong_key_protection_required}.
	StrongKeyProtectionRequired interface{} `field:"optional" json:"strongKeyProtectionRequired" yaml:"strongKeyProtectionRequired"`
}

