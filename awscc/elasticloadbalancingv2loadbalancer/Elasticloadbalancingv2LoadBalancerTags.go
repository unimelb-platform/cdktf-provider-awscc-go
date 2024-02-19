package elasticloadbalancingv2loadbalancer


type Elasticloadbalancingv2LoadBalancerTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#key Elasticloadbalancingv2LoadBalancer#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_load_balancer#value Elasticloadbalancingv2LoadBalancer#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

