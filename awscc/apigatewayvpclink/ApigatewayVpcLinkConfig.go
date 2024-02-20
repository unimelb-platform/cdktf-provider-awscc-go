package apigatewayvpclink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigatewayVpcLinkConfig struct {
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
	// A name for the VPC link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_vpc_link#name ApigatewayVpcLink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of network load balancer of the VPC targeted by the VPC link.
	//
	// The network load balancer must be owned by the same AWS account of the API owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_vpc_link#target_arns ApigatewayVpcLink#target_arns}
	TargetArns *[]*string `field:"required" json:"targetArns" yaml:"targetArns"`
	// A description of the VPC link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_vpc_link#description ApigatewayVpcLink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_vpc_link#tags ApigatewayVpcLink#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

