package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfigurationServerlessConfiguration struct {
	// Configurations for Redshift query engine serverless auth setup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#auth_configuration BedrockKnowledgeBase#auth_configuration}
	AuthConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryEngineConfigurationServerlessConfigurationAuthConfiguration `field:"optional" json:"authConfiguration" yaml:"authConfiguration"`
	// Workgroup arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#workgroup_arn BedrockKnowledgeBase#workgroup_arn}
	WorkgroupArn *string `field:"optional" json:"workgroupArn" yaml:"workgroupArn"`
}

