package pipespipe


type PipesPipeTargetParametersKinesisStreamParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#partition_key PipesPipe#partition_key}.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}

