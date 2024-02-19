package eksidentityproviderconfig


type EksIdentityProviderConfigOidcRequiredClaims struct {
	// The key of the requiredClaims.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_identity_provider_config#key EksIdentityProviderConfig#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the requiredClaims.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_identity_provider_config#value EksIdentityProviderConfig#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

