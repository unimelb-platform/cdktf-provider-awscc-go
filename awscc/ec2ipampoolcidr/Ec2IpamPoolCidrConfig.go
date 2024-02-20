package ec2ipampoolcidr

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2IpamPoolCidrConfig struct {
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
	// Id of the IPAM Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool_cidr#ipam_pool_id Ec2IpamPoolCidr#ipam_pool_id}
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Represents a single IPv4 or IPv6 CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool_cidr#cidr Ec2IpamPoolCidr#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// The desired netmask length of the provision.
	//
	// If set, IPAM will choose a block of free space with this size and return the CIDR representing it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ipam_pool_cidr#netmask_length Ec2IpamPoolCidr#netmask_length}
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

