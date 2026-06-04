package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumesSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#optional BatchJobDefinition#optional}.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#secret_name BatchJobDefinition#secret_name}.
	SecretName *string `field:"optional" json:"secretName" yaml:"secretName"`
}

