package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#target_nodes BatchJobDefinition#target_nodes}.
	TargetNodes *string `field:"required" json:"targetNodes" yaml:"targetNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#container BatchJobDefinition#container}.
	Container *BatchJobDefinitionNodePropertiesNodeRangePropertiesContainer `field:"optional" json:"container" yaml:"container"`
}

