package mediaconnectflowvpcinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MediaconnectFlowVpcInterfaceConfig struct {
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
	// The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_vpc_interface#flow_arn MediaconnectFlowVpcInterface#flow_arn}
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// Immutable and has to be a unique against other VpcInterfaces in this Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_vpc_interface#name MediaconnectFlowVpcInterface#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role Arn MediaConnect can assumes to create ENIs in customer's account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_vpc_interface#role_arn MediaconnectFlowVpcInterface#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Security Group IDs to be used on ENI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_vpc_interface#security_group_ids MediaconnectFlowVpcInterface#security_group_ids}
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Subnet must be in the AZ of the Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediaconnect_flow_vpc_interface#subnet_id MediaconnectFlowVpcInterface#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

