package ec2transitgatewayattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#subnet_ids Ec2TransitGatewayAttachment#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#transit_gateway_id Ec2TransitGatewayAttachment#transit_gateway_id}.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#vpc_id Ec2TransitGatewayAttachment#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The options for the transit gateway vpc attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#options Ec2TransitGatewayAttachment#options}
	Options *Ec2TransitGatewayAttachmentOptions `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_attachment#tags Ec2TransitGatewayAttachment#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

