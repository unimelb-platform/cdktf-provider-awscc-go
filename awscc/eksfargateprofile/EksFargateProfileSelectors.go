package eksfargateprofile


type EksFargateProfileSelectors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_fargate_profile#namespace EksFargateProfile#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_fargate_profile#labels EksFargateProfile#labels}.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

