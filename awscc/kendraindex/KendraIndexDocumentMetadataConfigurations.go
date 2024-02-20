package kendraindex


type KendraIndexDocumentMetadataConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_index#name KendraIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_index#type KendraIndex#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_index#relevance KendraIndex#relevance}.
	Relevance *KendraIndexDocumentMetadataConfigurationsRelevance `field:"optional" json:"relevance" yaml:"relevance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_index#search KendraIndex#search}.
	Search *KendraIndexDocumentMetadataConfigurationsSearch `field:"optional" json:"search" yaml:"search"`
}

