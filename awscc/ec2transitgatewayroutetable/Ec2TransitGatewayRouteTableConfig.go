package ec2transitgatewayroutetable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayRouteTableConfig struct {
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
	// The ID of the transit gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_route_table#transit_gateway_id Ec2TransitGatewayRouteTable#transit_gateway_id}
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Tags are composed of a Key/Value pair.
	//
	// You can use tags to categorize and track each parameter group. The tag value null is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway_route_table#tags Ec2TransitGatewayRouteTable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

