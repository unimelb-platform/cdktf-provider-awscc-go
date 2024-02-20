package ec2subnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2SubnetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#vpc_id Ec2Subnet#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#assign_ipv_6_address_on_creation Ec2Subnet#assign_ipv_6_address_on_creation}.
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#availability_zone Ec2Subnet#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#availability_zone_id Ec2Subnet#availability_zone_id}.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#cidr_block Ec2Subnet#cidr_block}.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#enable_dns_64 Ec2Subnet#enable_dns_64}.
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this subnet's CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_4_ipam_pool_id Ec2Subnet#ipv_4_ipam_pool_id}
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_4_netmask_length Ec2Subnet#ipv_4_netmask_length}
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_6_cidr_block Ec2Subnet#ipv_6_cidr_block}.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_6_cidr_blocks Ec2Subnet#ipv_6_cidr_blocks}.
	Ipv6CidrBlocks *[]*string `field:"optional" json:"ipv6CidrBlocks" yaml:"ipv6CidrBlocks"`
	// The ID of an IPv6 IPAM pool you want to use for allocating this subnet's CIDR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_6_ipam_pool_id Ec2Subnet#ipv_6_ipam_pool_id}
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_6_native Ec2Subnet#ipv_6_native}.
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// The netmask length of the IPv6 CIDR you want to allocate to this subnet from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#ipv_6_netmask_length Ec2Subnet#ipv_6_netmask_length}
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#map_public_ip_on_launch Ec2Subnet#map_public_ip_on_launch}.
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#outpost_arn Ec2Subnet#outpost_arn}.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#private_dns_name_options_on_launch Ec2Subnet#private_dns_name_options_on_launch}.
	PrivateDnsNameOptionsOnLaunch *Ec2SubnetPrivateDnsNameOptionsOnLaunch `field:"optional" json:"privateDnsNameOptionsOnLaunch" yaml:"privateDnsNameOptionsOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_subnet#tags Ec2Subnet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

