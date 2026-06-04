package ec2routeserverendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2RouteServerEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_endpoint#route_server_id Ec2RouteServerEndpoint#route_server_id}
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// Subnet ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_endpoint#subnet_id Ec2RouteServerEndpoint#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_route_server_endpoint#tags Ec2RouteServerEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

