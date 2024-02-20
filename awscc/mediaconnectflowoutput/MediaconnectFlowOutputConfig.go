package mediaconnectflowoutput

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectFlowOutputConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#flow_arn MediaconnectFlowOutput#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The protocol that is used by the source or output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#protocol MediaconnectFlowOutput#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The range of IP addresses that should be allowed to initiate output requests to this flow.
	//
	// These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#cidr_allow_list MediaconnectFlowOutput#cidr_allow_list}
	CidrAllowList *[]*string `field:"optional" json:"cidrAllowList" yaml:"cidrAllowList"`
	// A description of the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#description MediaconnectFlowOutput#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The address where you want to send the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#destination MediaconnectFlowOutput#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The type of key used for the encryption.
	//
	// If no keyType is provided, the service will use the default setting (static-key).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#encryption MediaconnectFlowOutput#encryption}
	Encryption *MediaconnectFlowOutputEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#max_latency MediaconnectFlowOutput#max_latency}
	MaxLatency *float64 `field:"optional" json:"maxLatency" yaml:"maxLatency"`
	// The minimum latency in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#min_latency MediaconnectFlowOutput#min_latency}
	MinLatency *float64 `field:"optional" json:"minLatency" yaml:"minLatency"`
	// The name of the output. This value must be unique within the current flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#name MediaconnectFlowOutput#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port to use when content is distributed to this output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#port MediaconnectFlowOutput#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The remote ID for the Zixi-pull stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#remote_id MediaconnectFlowOutput#remote_id}
	RemoteId *string `field:"optional" json:"remoteId" yaml:"remoteId"`
	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#smoothing_latency MediaconnectFlowOutput#smoothing_latency}
	SmoothingLatency *float64 `field:"optional" json:"smoothingLatency" yaml:"smoothingLatency"`
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#stream_id MediaconnectFlowOutput#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
	// The name of the VPC interface attachment to use for this output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_output#vpc_interface_attachment MediaconnectFlowOutput#vpc_interface_attachment}
	VpcInterfaceAttachment *MediaconnectFlowOutputVpcInterfaceAttachment `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

