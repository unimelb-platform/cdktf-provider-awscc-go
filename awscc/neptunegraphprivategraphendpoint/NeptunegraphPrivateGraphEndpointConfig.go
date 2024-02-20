package neptunegraphprivategraphendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NeptunegraphPrivateGraphEndpointConfig struct {
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
	// The auto-generated Graph Id assigned by the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptunegraph_private_graph_endpoint#graph_identifier NeptunegraphPrivateGraphEndpoint#graph_identifier}
	GraphIdentifier *string `field:"required" json:"graphIdentifier" yaml:"graphIdentifier"`
	// The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptunegraph_private_graph_endpoint#vpc_id NeptunegraphPrivateGraphEndpoint#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptunegraph_private_graph_endpoint#security_group_ids NeptunegraphPrivateGraphEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/neptunegraph_private_graph_endpoint#subnet_ids NeptunegraphPrivateGraphEndpoint#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

