package ec2vpcendpoint


type Ec2VpcEndpointDnsOptions struct {
	// The DNS records created for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint#dns_record_ip_type Ec2VpcEndpoint#dns_record_ip_type}
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Indicates whether to enable private DNS only for inbound endpoints.
	//
	// This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint#private_dns_only_for_inbound_resolver_endpoint Ec2VpcEndpoint#private_dns_only_for_inbound_resolver_endpoint}
	PrivateDnsOnlyForInboundResolverEndpoint *string `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
}

