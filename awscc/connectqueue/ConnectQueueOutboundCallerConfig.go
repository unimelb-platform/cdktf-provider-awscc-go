package connectqueue


type ConnectQueueOutboundCallerConfig struct {
	// The caller ID name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_queue#outbound_caller_id_name ConnectQueue#outbound_caller_id_name}
	OutboundCallerIdName *string `field:"optional" json:"outboundCallerIdName" yaml:"outboundCallerIdName"`
	// The caller ID number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_queue#outbound_caller_id_number_arn ConnectQueue#outbound_caller_id_number_arn}
	OutboundCallerIdNumberArn *string `field:"optional" json:"outboundCallerIdNumberArn" yaml:"outboundCallerIdNumberArn"`
	// The outbound whisper flow to be used during an outbound call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/connect_queue#outbound_flow_arn ConnectQueue#outbound_flow_arn}
	OutboundFlowArn *string `field:"optional" json:"outboundFlowArn" yaml:"outboundFlowArn"`
}

