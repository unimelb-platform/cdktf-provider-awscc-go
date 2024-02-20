package ec2securitygroupegress

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2SecurityGroupEgressConfig struct {
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
	// The ID of the security group.
	//
	// You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#group_id Ec2SecurityGroupEgress#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// [VPC only] Use -1 to specify all protocols.
	//
	// When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#ip_protocol Ec2SecurityGroupEgress#ip_protocol}
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#cidr_ip Ec2SecurityGroupEgress#cidr_ip}
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// [VPC only] The IPv6 ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#cidr_ipv_6 Ec2SecurityGroupEgress#cidr_ipv_6}
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// Resource Type definition for an egress (outbound) security group rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#description Ec2SecurityGroupEgress#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// [EC2-VPC only] The ID of a prefix list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#destination_prefix_list_id Ec2SecurityGroupEgress#destination_prefix_list_id}
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// You must specify a destination security group (DestinationPrefixListId or DestinationSecurityGroupId) or a CIDR range (CidrIp or CidrIpv6).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#destination_security_group_id Ec2SecurityGroupEgress#destination_security_group_id}
	DestinationSecurityGroupId *string `field:"optional" json:"destinationSecurityGroupId" yaml:"destinationSecurityGroupId"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.
	//
	// A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#from_port Ec2SecurityGroupEgress#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	//
	// A value of -1 indicates all ICMP/ICMPv6 codes. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_security_group_egress#to_port Ec2SecurityGroupEgress#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

