package iotsitewisegateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotsitewiseGatewayConfig struct {
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
	// A unique, friendly name for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#gateway_name IotsitewiseGateway#gateway_name}
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The gateway's platform. You can only specify one platform in a gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#gateway_platform IotsitewiseGateway#gateway_platform}
	GatewayPlatform *IotsitewiseGatewayGatewayPlatform `field:"required" json:"gatewayPlatform" yaml:"gatewayPlatform"`
	// A list of gateway capability summaries that each contain a namespace and status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#gateway_capability_summaries IotsitewiseGateway#gateway_capability_summaries}
	GatewayCapabilitySummaries interface{} `field:"optional" json:"gatewayCapabilitySummaries" yaml:"gatewayCapabilitySummaries"`
	// A list of key-value pairs that contain metadata for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_gateway#tags IotsitewiseGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

