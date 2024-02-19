package elasticloadbalancingv2loadbalancer


type Elasticloadbalancingv2LoadBalancerSubnetMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#subnet_id Elasticloadbalancingv2LoadBalancer#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#allocation_id Elasticloadbalancingv2LoadBalancer#allocation_id}.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#i_pv_6_address Elasticloadbalancingv2LoadBalancer#i_pv_6_address}.
	IPv6Address *string `field:"optional" json:"iPv6Address" yaml:"iPv6Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#private_i_pv_4_address Elasticloadbalancingv2LoadBalancer#private_i_pv_4_address}.
	PrivateIPv4Address *string `field:"optional" json:"privateIPv4Address" yaml:"privateIPv4Address"`
}

