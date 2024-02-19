package pcaconnectoradtemplate


type PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#key_spec PcaconnectoradTemplate#key_spec}.
	KeySpec *string `field:"required" json:"keySpec" yaml:"keySpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#minimal_key_length PcaconnectoradTemplate#minimal_key_length}.
	MinimalKeyLength *float64 `field:"required" json:"minimalKeyLength" yaml:"minimalKeyLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#algorithm PcaconnectoradTemplate#algorithm}.
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#crypto_providers PcaconnectoradTemplate#crypto_providers}.
	CryptoProviders *[]*string `field:"optional" json:"cryptoProviders" yaml:"cryptoProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pcaconnectorad_template#key_usage_property PcaconnectoradTemplate#key_usage_property}.
	KeyUsageProperty *PcaconnectoradTemplateDefinitionTemplateV4PrivateKeyAttributesKeyUsageProperty `field:"optional" json:"keyUsageProperty" yaml:"keyUsageProperty"`
}

