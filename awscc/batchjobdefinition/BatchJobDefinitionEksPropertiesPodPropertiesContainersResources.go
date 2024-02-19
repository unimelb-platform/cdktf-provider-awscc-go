package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesContainersResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#limits BatchJobDefinition#limits}.
	Limits *string `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#requests BatchJobDefinition#requests}.
	Requests *string `field:"optional" json:"requests" yaml:"requests"`
}

