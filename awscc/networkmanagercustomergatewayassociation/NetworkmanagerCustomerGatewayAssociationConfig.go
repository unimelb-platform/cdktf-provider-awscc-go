package networkmanagercustomergatewayassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerCustomerGatewayAssociationConfig struct {
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
	// The Amazon Resource Name (ARN) of the customer gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_customer_gateway_association#customer_gateway_arn NetworkmanagerCustomerGatewayAssociation#customer_gateway_arn}
	CustomerGatewayArn *string `field:"required" json:"customerGatewayArn" yaml:"customerGatewayArn"`
	// The ID of the device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_customer_gateway_association#device_id NetworkmanagerCustomerGatewayAssociation#device_id}
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// The ID of the global network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_customer_gateway_association#global_network_id NetworkmanagerCustomerGatewayAssociation#global_network_id}
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkmanager_customer_gateway_association#link_id NetworkmanagerCustomerGatewayAssociation#link_id}
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
}

