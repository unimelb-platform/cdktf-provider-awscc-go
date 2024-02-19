package batchjobdefinition


type BatchJobDefinitionContainerPropertiesLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#log_driver BatchJobDefinition#log_driver}.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#options BatchJobDefinition#options}.
	Options *string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#secret_options BatchJobDefinition#secret_options}.
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

