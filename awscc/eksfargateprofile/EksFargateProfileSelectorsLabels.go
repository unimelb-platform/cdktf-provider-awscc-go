package eksfargateprofile


type EksFargateProfileSelectorsLabels struct {
	// The key name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_fargate_profile#key EksFargateProfile#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_fargate_profile#value EksFargateProfile#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

