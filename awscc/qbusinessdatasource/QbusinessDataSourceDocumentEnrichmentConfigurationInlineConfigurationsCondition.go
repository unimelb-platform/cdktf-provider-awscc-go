package qbusinessdatasource


type QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_data_source#key QbusinessDataSource#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_data_source#operator QbusinessDataSource#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_data_source#value QbusinessDataSource#value}.
	Value *QbusinessDataSourceDocumentEnrichmentConfigurationInlineConfigurationsConditionValue `field:"optional" json:"value" yaml:"value"`
}

