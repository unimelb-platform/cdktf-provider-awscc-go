package ec2dhcpoptions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2DhcpOptionsConfig struct {
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
	// This value is used to complete unqualified DNS hostnames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#domain_name Ec2DhcpOptions#domain_name}
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The IPv4 addresses of up to four domain name servers, or AmazonProvidedDNS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#domain_name_servers Ec2DhcpOptions#domain_name_servers}
	DomainNameServers *[]*string `field:"optional" json:"domainNameServers" yaml:"domainNameServers"`
	// The IPv4 addresses of up to four NetBIOS name servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#netbios_name_servers Ec2DhcpOptions#netbios_name_servers}
	NetbiosNameServers *[]*string `field:"optional" json:"netbiosNameServers" yaml:"netbiosNameServers"`
	// The NetBIOS node type (1, 2, 4, or 8).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#netbios_node_type Ec2DhcpOptions#netbios_node_type}
	NetbiosNodeType *float64 `field:"optional" json:"netbiosNodeType" yaml:"netbiosNodeType"`
	// The IPv4 addresses of up to four Network Time Protocol (NTP) servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#ntp_servers Ec2DhcpOptions#ntp_servers}
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
	// Any tags assigned to the DHCP options set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_dhcp_options#tags Ec2DhcpOptions#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

