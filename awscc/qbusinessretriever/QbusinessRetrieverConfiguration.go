package qbusinessretriever


type QbusinessRetrieverConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_retriever#kendra_index_configuration QbusinessRetriever#kendra_index_configuration}.
	KendraIndexConfiguration *QbusinessRetrieverConfigurationKendraIndexConfiguration `field:"optional" json:"kendraIndexConfiguration" yaml:"kendraIndexConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/qbusiness_retriever#native_index_configuration QbusinessRetriever#native_index_configuration}.
	NativeIndexConfiguration *QbusinessRetrieverConfigurationNativeIndexConfiguration `field:"optional" json:"nativeIndexConfiguration" yaml:"nativeIndexConfiguration"`
}

