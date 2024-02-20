package ec2transitgatewayconnect

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayConnectConfig struct {
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
	// The Connect attachment options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_connect#options Ec2TransitGatewayConnect#options}
	Options *Ec2TransitGatewayConnectOptions `field:"required" json:"options" yaml:"options"`
	// The ID of the attachment from which the Connect attachment was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_connect#transport_transit_gateway_attachment_id Ec2TransitGatewayConnect#transport_transit_gateway_attachment_id}
	TransportTransitGatewayAttachmentId *string `field:"required" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// The tags for the attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_connect#tags Ec2TransitGatewayConnect#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

