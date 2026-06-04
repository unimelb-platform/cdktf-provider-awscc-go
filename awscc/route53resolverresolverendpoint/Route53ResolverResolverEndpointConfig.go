package route53resolverresolverendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53ResolverResolverEndpointConfig struct {
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
	// Indicates whether the Resolver endpoint allows inbound or outbound DNS queries: - INBOUND: allows DNS queries to your VPC from your network  - OUTBOUND: allows DNS queries from your VPC to your network  - INBOUND_DELEGATION: allows DNS queries to your VPC from your network with authoritative answers from private hosted zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#direction Route53ResolverResolverEndpoint#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	//
	// The subnet ID uniquely identifies a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#ip_addresses Route53ResolverResolverEndpoint#ip_addresses}
	IpAddresses interface{} `field:"required" json:"ipAddresses" yaml:"ipAddresses"`
	// The ID of one or more security groups that control access to this VPC.
	//
	// The security group must include one or more inbound rules (for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access. For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#security_group_ids Route53ResolverResolverEndpoint#security_group_ids}
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#name Route53ResolverResolverEndpoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN (Amazon Resource Name) for the Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#outpost_arn Route53ResolverResolverEndpoint#outpost_arn}
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The Amazon EC2 instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#preferred_instance_type Route53ResolverResolverEndpoint#preferred_instance_type}
	PreferredInstanceType *string `field:"optional" json:"preferredInstanceType" yaml:"preferredInstanceType"`
	// Protocols used for the endpoint. DoH-FIPS is applicable for inbound endpoints only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#protocols Route53ResolverResolverEndpoint#protocols}
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The Resolver endpoint IP address type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#resolver_endpoint_type Route53ResolverResolverEndpoint#resolver_endpoint_type}
	ResolverEndpointType *string `field:"optional" json:"resolverEndpointType" yaml:"resolverEndpointType"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53resolver_resolver_endpoint#tags Route53ResolverResolverEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

