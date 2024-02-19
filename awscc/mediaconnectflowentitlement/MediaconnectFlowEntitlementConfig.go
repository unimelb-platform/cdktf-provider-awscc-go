package mediaconnectflowentitlement

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectFlowEntitlementConfig struct {
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
	// A description of the entitlement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#description MediaconnectFlowEntitlement#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The ARN of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#flow_arn MediaconnectFlowEntitlement#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the entitlement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#name MediaconnectFlowEntitlement#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS account IDs that you want to share your content with.
	//
	// The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#subscribers MediaconnectFlowEntitlement#subscribers}
	Subscribers *[]*string `field:"required" json:"subscribers" yaml:"subscribers"`
	// Percentage from 0-100 of the data transfer cost to be billed to the subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#data_transfer_subscriber_fee_percent MediaconnectFlowEntitlement#data_transfer_subscriber_fee_percent}
	DataTransferSubscriberFeePercent *float64 `field:"optional" json:"dataTransferSubscriberFeePercent" yaml:"dataTransferSubscriberFeePercent"`
	// The type of encryption that will be used on the output that is associated with this entitlement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#encryption MediaconnectFlowEntitlement#encryption}
	Encryption *MediaconnectFlowEntitlementEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// An indication of whether the entitlement is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_entitlement#entitlement_status MediaconnectFlowEntitlement#entitlement_status}
	EntitlementStatus *string `field:"optional" json:"entitlementStatus" yaml:"entitlementStatus"`
}

