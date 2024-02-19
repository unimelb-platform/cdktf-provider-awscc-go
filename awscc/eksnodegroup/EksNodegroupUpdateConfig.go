package eksnodegroup


type EksNodegroupUpdateConfig struct {
	// The maximum number of nodes unavailable at once during a version update.
	//
	// Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#max_unavailable EksNodegroup#max_unavailable}
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// The maximum percentage of nodes unavailable during a version update.
	//
	// This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#max_unavailable_percentage EksNodegroup#max_unavailable_percentage}
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
}

