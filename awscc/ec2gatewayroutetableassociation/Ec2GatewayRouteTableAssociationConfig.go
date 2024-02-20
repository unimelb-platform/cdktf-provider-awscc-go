package ec2gatewayroutetableassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2GatewayRouteTableAssociationConfig struct {
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
	// The ID of the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_gateway_route_table_association#gateway_id Ec2GatewayRouteTableAssociation#gateway_id}
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The ID of the route table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_gateway_route_table_association#route_table_id Ec2GatewayRouteTableAssociation#route_table_id}
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

