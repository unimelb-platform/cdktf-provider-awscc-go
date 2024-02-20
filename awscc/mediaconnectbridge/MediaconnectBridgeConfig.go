package mediaconnectbridge

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectBridgeConfig struct {
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
	// The name of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The placement Amazon Resource Number (ARN) of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#placement_arn MediaconnectBridge#placement_arn}
	PlacementArn *string `field:"required" json:"placementArn" yaml:"placementArn"`
	// The sources on this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#sources MediaconnectBridge#sources}
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#egress_gateway_bridge MediaconnectBridge#egress_gateway_bridge}.
	EgressGatewayBridge *MediaconnectBridgeEgressGatewayBridge `field:"optional" json:"egressGatewayBridge" yaml:"egressGatewayBridge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#ingress_gateway_bridge MediaconnectBridge#ingress_gateway_bridge}.
	IngressGatewayBridge *MediaconnectBridgeIngressGatewayBridge `field:"optional" json:"ingressGatewayBridge" yaml:"ingressGatewayBridge"`
	// The outputs on this bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#outputs MediaconnectBridge#outputs}
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The settings for source failover.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge#source_failover_config MediaconnectBridge#source_failover_config}
	SourceFailoverConfig *MediaconnectBridgeSourceFailoverConfig `field:"optional" json:"sourceFailoverConfig" yaml:"sourceFailoverConfig"`
}

