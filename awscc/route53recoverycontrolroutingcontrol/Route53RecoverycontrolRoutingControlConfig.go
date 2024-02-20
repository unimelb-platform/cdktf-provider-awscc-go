package route53recoverycontrolroutingcontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoverycontrolRoutingControlConfig struct {
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
	// The name of the routing control. You can use any non-white space character in the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_routing_control#name Route53RecoverycontrolRoutingControl#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Arn associated with Control Panel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_routing_control#cluster_arn Route53RecoverycontrolRoutingControl#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The Amazon Resource Name (ARN) of the control panel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoverycontrol_routing_control#control_panel_arn Route53RecoverycontrolRoutingControl#control_panel_arn}
	ControlPanelArn *string `field:"optional" json:"controlPanelArn" yaml:"controlPanelArn"`
}

