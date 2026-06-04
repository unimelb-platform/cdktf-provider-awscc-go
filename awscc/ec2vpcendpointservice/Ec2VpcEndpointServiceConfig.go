package ec2vpcendpointservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2VpcEndpointServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#acceptance_required Ec2VpcEndpointService#acceptance_required}.
	AcceptanceRequired interface{} `field:"optional" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#contributor_insights_enabled Ec2VpcEndpointService#contributor_insights_enabled}.
	ContributorInsightsEnabled interface{} `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#gateway_load_balancer_arns Ec2VpcEndpointService#gateway_load_balancer_arns}.
	GatewayLoadBalancerArns *[]*string `field:"optional" json:"gatewayLoadBalancerArns" yaml:"gatewayLoadBalancerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#network_load_balancer_arns Ec2VpcEndpointService#network_load_balancer_arns}.
	NetworkLoadBalancerArns *[]*string `field:"optional" json:"networkLoadBalancerArns" yaml:"networkLoadBalancerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#payer_responsibility Ec2VpcEndpointService#payer_responsibility}.
	PayerResponsibility *string `field:"optional" json:"payerResponsibility" yaml:"payerResponsibility"`
	// Specify which Ip Address types are supported for VPC endpoint service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#supported_ip_address_types Ec2VpcEndpointService#supported_ip_address_types}
	SupportedIpAddressTypes *[]*string `field:"optional" json:"supportedIpAddressTypes" yaml:"supportedIpAddressTypes"`
	// The Regions from which service consumers can access the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#supported_regions Ec2VpcEndpointService#supported_regions}
	SupportedRegions *[]*string `field:"optional" json:"supportedRegions" yaml:"supportedRegions"`
	// The tags to add to the VPC endpoint service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpc_endpoint_service#tags Ec2VpcEndpointService#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

