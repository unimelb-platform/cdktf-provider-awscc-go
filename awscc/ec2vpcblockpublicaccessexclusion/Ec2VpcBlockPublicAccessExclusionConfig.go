package ec2vpcblockpublicaccessexclusion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpcBlockPublicAccessExclusionConfig struct {
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
	// The desired Block Public Access Exclusion Mode for a specific VPC/Subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_block_public_access_exclusion#internet_gateway_exclusion_mode Ec2VpcBlockPublicAccessExclusion#internet_gateway_exclusion_mode}
	InternetGatewayExclusionMode *string `field:"required" json:"internetGatewayExclusionMode" yaml:"internetGatewayExclusionMode"`
	// The ID of the subnet. Required only if you don't specify VpcId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_block_public_access_exclusion#subnet_id Ec2VpcBlockPublicAccessExclusion#subnet_id}
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_block_public_access_exclusion#tags Ec2VpcBlockPublicAccessExclusion#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the vpc. Required only if you don't specify SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_block_public_access_exclusion#vpc_id Ec2VpcBlockPublicAccessExclusion#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

