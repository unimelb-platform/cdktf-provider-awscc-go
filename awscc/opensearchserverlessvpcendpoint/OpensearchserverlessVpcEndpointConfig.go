package opensearchserverlessvpcendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpensearchserverlessVpcEndpointConfig struct {
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
	// The name of the VPC Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_vpc_endpoint#name OpensearchserverlessVpcEndpoint#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of one or more subnets in which to create an endpoint network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_vpc_endpoint#subnet_ids OpensearchserverlessVpcEndpoint#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC in which the endpoint will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_vpc_endpoint#vpc_id OpensearchserverlessVpcEndpoint#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of one or more security groups to associate with the endpoint network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/opensearchserverless_vpc_endpoint#security_group_ids OpensearchserverlessVpcEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

