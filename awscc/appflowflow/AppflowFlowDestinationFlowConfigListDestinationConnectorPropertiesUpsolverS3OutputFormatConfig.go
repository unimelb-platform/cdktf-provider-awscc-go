package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverS3OutputFormatConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#prefix_config AppflowFlow#prefix_config}.
	PrefixConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfig `field:"required" json:"prefixConfig" yaml:"prefixConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#aggregation_config AppflowFlow#aggregation_config}.
	AggregationConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfig `field:"optional" json:"aggregationConfig" yaml:"aggregationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#file_type AppflowFlow#file_type}.
	FileType *string `field:"optional" json:"fileType" yaml:"fileType"`
}

