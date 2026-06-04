package batchjobdefinition


type BatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#log_driver BatchJobDefinition#log_driver}.
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#options BatchJobDefinition#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#secret_options BatchJobDefinition#secret_options}.
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

