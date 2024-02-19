package ec2verifiedaccesstrustprovider


type Ec2VerifiedAccessTrustProviderOidcOptions struct {
	// The OIDC authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#authorization_endpoint Ec2VerifiedAccessTrustProvider#authorization_endpoint}
	AuthorizationEndpoint *string `field:"optional" json:"authorizationEndpoint" yaml:"authorizationEndpoint"`
	// The client identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#client_id Ec2VerifiedAccessTrustProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#client_secret Ec2VerifiedAccessTrustProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The OIDC issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#issuer Ec2VerifiedAccessTrustProvider#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user.
	//
	// Each scope returns a specific set of user attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#scope Ec2VerifiedAccessTrustProvider#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The OIDC token endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#token_endpoint Ec2VerifiedAccessTrustProvider#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// The OIDC user info endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_trust_provider#user_info_endpoint Ec2VerifiedAccessTrustProvider#user_info_endpoint}
	UserInfoEndpoint *string `field:"optional" json:"userInfoEndpoint" yaml:"userInfoEndpoint"`
}

