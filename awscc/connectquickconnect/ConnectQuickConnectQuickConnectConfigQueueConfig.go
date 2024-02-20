package connectquickconnect


type ConnectQuickConnectQuickConnectConfigQueueConfig struct {
	// The identifier of the contact flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#contact_flow_arn ConnectQuickConnect#contact_flow_arn}
	ContactFlowArn *string `field:"required" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The identifier for the queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_quick_connect#queue_arn ConnectQuickConnect#queue_arn}
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
}

