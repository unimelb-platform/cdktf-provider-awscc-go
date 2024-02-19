package ec2vpcgatewayattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpcGatewayAttachmentConfig struct {
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
	// The ID of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_gateway_attachment#vpc_id Ec2VpcGatewayAttachment#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of the internet gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_gateway_attachment#internet_gateway_id Ec2VpcGatewayAttachment#internet_gateway_id}
	InternetGatewayId *string `field:"optional" json:"internetGatewayId" yaml:"internetGatewayId"`
	// The ID of the virtual private gateway. You must specify either InternetGatewayId or VpnGatewayId, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_gateway_attachment#vpn_gateway_id Ec2VpcGatewayAttachment#vpn_gateway_id}
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

