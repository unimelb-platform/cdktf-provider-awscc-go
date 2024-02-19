package batchjobdefinition


type BatchJobDefinitionNodeProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#main_node BatchJobDefinition#main_node}.
	MainNode *float64 `field:"required" json:"mainNode" yaml:"mainNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#node_range_properties BatchJobDefinition#node_range_properties}.
	NodeRangeProperties interface{} `field:"required" json:"nodeRangeProperties" yaml:"nodeRangeProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#num_nodes BatchJobDefinition#num_nodes}.
	NumNodes *float64 `field:"required" json:"numNodes" yaml:"numNodes"`
}

