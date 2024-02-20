package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumesSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#secret_name BatchJobDefinition#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#optional BatchJobDefinition#optional}.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

