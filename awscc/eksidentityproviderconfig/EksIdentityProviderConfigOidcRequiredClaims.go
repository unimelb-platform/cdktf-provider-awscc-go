package eksidentityproviderconfig


type EksIdentityProviderConfigOidcRequiredClaims struct {
	// The key of the requiredClaims.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_identity_provider_config#key EksIdentityProviderConfig#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the requiredClaims.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_identity_provider_config#value EksIdentityProviderConfig#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

