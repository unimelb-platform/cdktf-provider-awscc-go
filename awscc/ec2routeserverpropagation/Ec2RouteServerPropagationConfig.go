package ec2routeserverpropagation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2RouteServerPropagationConfig struct {
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
	// Route Server ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_propagation#route_server_id Ec2RouteServerPropagation#route_server_id}
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// Route Table ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_propagation#route_table_id Ec2RouteServerPropagation#route_table_id}
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

