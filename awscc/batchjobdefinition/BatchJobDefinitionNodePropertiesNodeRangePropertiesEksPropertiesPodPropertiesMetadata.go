package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#annotations BatchJobDefinition#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#labels BatchJobDefinition#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#namespace BatchJobDefinition#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

