package mediaconnectbridgeoutput

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectBridgeOutputConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_output#bridge_arn MediaconnectBridgeOutput#bridge_arn}
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The network output name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_output#name MediaconnectBridgeOutput#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The output of the bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_bridge_output#network_output MediaconnectBridgeOutput#network_output}
	NetworkOutput *MediaconnectBridgeOutputNetworkOutput `field:"required" json:"networkOutput" yaml:"networkOutput"`
}

