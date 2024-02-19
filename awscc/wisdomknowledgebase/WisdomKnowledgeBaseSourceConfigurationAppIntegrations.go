package wisdomknowledgebase


type WisdomKnowledgeBaseSourceConfigurationAppIntegrations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#app_integration_arn WisdomKnowledgeBase#app_integration_arn}.
	AppIntegrationArn *string `field:"required" json:"appIntegrationArn" yaml:"appIntegrationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wisdom_knowledge_base#object_fields WisdomKnowledgeBase#object_fields}.
	ObjectFields *[]*string `field:"optional" json:"objectFields" yaml:"objectFields"`
}

