package ec2subnetroutetableassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2SubnetRouteTableAssociationConfig struct {
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
	// The ID of the route table.  The physical ID changes when the route table ID is changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_subnet_route_table_association#route_table_id Ec2SubnetRouteTableAssociation#route_table_id}
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The ID of the subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_subnet_route_table_association#subnet_id Ec2SubnetRouteTableAssociation#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

