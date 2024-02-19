package eksnodegroup


type EksNodegroupScalingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#desired_size EksNodegroup#desired_size}.
	DesiredSize *float64 `field:"optional" json:"desiredSize" yaml:"desiredSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#max_size EksNodegroup#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#min_size EksNodegroup#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
}

