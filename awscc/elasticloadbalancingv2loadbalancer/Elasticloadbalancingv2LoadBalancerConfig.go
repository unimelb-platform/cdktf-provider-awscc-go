package elasticloadbalancingv2loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Elasticloadbalancingv2LoadBalancerConfig struct {
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
	// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through PrivateLink.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#enforce_security_group_inbound_rules_on_private_link_traffic Elasticloadbalancingv2LoadBalancer#enforce_security_group_inbound_rules_on_private_link_traffic}
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `field:"optional" json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// The type of IP addresses used by the subnets for your load balancer.
	//
	// The possible values are ipv4 (for IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#ip_address_type Elasticloadbalancingv2LoadBalancer#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The load balancer attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#load_balancer_attributes Elasticloadbalancingv2LoadBalancer#load_balancer_attributes}
	LoadBalancerAttributes interface{} `field:"optional" json:"loadBalancerAttributes" yaml:"loadBalancerAttributes"`
	// The name of the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#name Elasticloadbalancingv2LoadBalancer#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The nodes of an Internet-facing load balancer have public IP addresses.
	//
	// The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet. The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer. The default is an Internet-facing load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#scheme Elasticloadbalancingv2LoadBalancer#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// The IDs of the security groups for the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#security_groups Elasticloadbalancingv2LoadBalancer#security_groups}
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the public subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#subnet_mappings Elasticloadbalancingv2LoadBalancer#subnet_mappings}
	SubnetMappings interface{} `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
	// The IDs of the public subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#subnets Elasticloadbalancingv2LoadBalancer#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags to assign to the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#tags Elasticloadbalancingv2LoadBalancer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of load balancer. The default is application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#type Elasticloadbalancingv2LoadBalancer#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

