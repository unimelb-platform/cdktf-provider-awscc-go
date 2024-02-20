package eksnodegroup


type EksNodegroupTaints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#effect EksNodegroup#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#key EksNodegroup#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup#value EksNodegroup#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

