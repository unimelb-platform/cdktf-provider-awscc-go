package batchjobdefinition


type BatchJobDefinitionEksProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#pod_properties BatchJobDefinition#pod_properties}.
	PodProperties *BatchJobDefinitionEksPropertiesPodProperties `field:"optional" json:"podProperties" yaml:"podProperties"`
}

