package ec2instance


type Ec2InstancePrivateDnsNameOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#enable_resource_name_dns_aaaa_record Ec2Instance#enable_resource_name_dns_aaaa_record}.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#enable_resource_name_dns_a_record Ec2Instance#enable_resource_name_dns_a_record}.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#hostname_type Ec2Instance#hostname_type}.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}

