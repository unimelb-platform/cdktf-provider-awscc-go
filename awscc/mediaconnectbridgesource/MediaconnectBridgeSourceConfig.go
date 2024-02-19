package mediaconnectbridgesource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectBridgeSourceConfig struct {
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
	// The Amazon Resource Number (ARN) of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#bridge_arn MediaconnectBridgeSource#bridge_arn}
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#name MediaconnectBridgeSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#flow_source MediaconnectBridgeSource#flow_source}
	FlowSource *MediaconnectBridgeSourceFlowSource `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the bridge. A network source originates at your premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_source#network_source MediaconnectBridgeSource#network_source}
	NetworkSource *MediaconnectBridgeSourceNetworkSource `field:"optional" json:"networkSource" yaml:"networkSource"`
}

