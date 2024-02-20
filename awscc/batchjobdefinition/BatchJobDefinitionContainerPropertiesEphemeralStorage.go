package batchjobdefinition


type BatchJobDefinitionContainerPropertiesEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#size_in_gi_b BatchJobDefinition#size_in_gi_b}.
	SizeInGiB *float64 `field:"required" json:"sizeInGiB" yaml:"sizeInGiB"`
}

