package ec2launchtemplate


type Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions struct {
	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#enable_resource_name_dns_aaaa_record Ec2LaunchTemplate#enable_resource_name_dns_aaaa_record}
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#enable_resource_name_dns_a_record Ec2LaunchTemplate#enable_resource_name_dns_a_record}
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// The type of hostname for EC2 instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_launch_template#hostname_type Ec2LaunchTemplate#hostname_type}
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

