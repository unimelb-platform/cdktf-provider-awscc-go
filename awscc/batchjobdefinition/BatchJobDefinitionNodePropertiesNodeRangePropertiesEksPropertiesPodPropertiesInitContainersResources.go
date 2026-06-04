package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#limits BatchJobDefinition#limits}.
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#requests BatchJobDefinition#requests}.
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

