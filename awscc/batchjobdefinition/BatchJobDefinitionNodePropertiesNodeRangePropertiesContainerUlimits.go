package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerUlimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#hard_limit BatchJobDefinition#hard_limit}.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#soft_limit BatchJobDefinition#soft_limit}.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

