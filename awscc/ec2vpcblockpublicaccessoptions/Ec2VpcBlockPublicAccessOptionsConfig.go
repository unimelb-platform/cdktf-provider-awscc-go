package ec2vpcblockpublicaccessoptions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpcBlockPublicAccessOptionsConfig struct {
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
	// The desired Block Public Access mode for Internet Gateways in your account.
	//
	// We do not allow to create in a off mode as this is the default value
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_block_public_access_options#internet_gateway_block_mode Ec2VpcBlockPublicAccessOptions#internet_gateway_block_mode}
	InternetGatewayBlockMode *string `field:"required" json:"internetGatewayBlockMode" yaml:"internetGatewayBlockMode"`
}

