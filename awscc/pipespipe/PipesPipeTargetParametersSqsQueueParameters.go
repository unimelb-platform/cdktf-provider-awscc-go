package pipespipe


type PipesPipeTargetParametersSqsQueueParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#message_deduplication_id PipesPipe#message_deduplication_id}.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#message_group_id PipesPipe#message_group_id}.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

