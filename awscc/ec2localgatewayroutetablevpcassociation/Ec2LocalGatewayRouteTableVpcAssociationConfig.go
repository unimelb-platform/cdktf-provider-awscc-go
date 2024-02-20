package ec2localgatewayroutetablevpcassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2LocalGatewayRouteTableVpcAssociationConfig struct {
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
	// The ID of the local gateway route table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_local_gateway_route_table_vpc_association#local_gateway_route_table_id Ec2LocalGatewayRouteTableVpcAssociation#local_gateway_route_table_id}
	LocalGatewayRouteTableId *string `field:"required" json:"localGatewayRouteTableId" yaml:"localGatewayRouteTableId"`
	// The ID of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_local_gateway_route_table_vpc_association#vpc_id Ec2LocalGatewayRouteTableVpcAssociation#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The tags for the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_local_gateway_route_table_vpc_association#tags Ec2LocalGatewayRouteTableVpcAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

