package ec2vpcpeeringconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpcPeeringConnectionConfig struct {
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
	// The ID of the VPC with which you are creating the VPC peering connection.
	//
	// You must specify this parameter in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_peering_connection#peer_vpc_id Ec2VpcPeeringConnection#peer_vpc_id}
	PeerVpcId *string `field:"required" json:"peerVpcId" yaml:"peerVpcId"`
	// The ID of the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_peering_connection#vpc_id Ec2VpcPeeringConnection#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The AWS account ID of the owner of the accepter VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_peering_connection#peer_owner_id Ec2VpcPeeringConnection#peer_owner_id}
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	// The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_peering_connection#peer_region Ec2VpcPeeringConnection#peer_region}
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
	// The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in another AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_peering_connection#peer_role_arn Ec2VpcPeeringConnection#peer_role_arn}
	PeerRoleArn *string `field:"optional" json:"peerRoleArn" yaml:"peerRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpc_peering_connection#tags Ec2VpcPeeringConnection#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

