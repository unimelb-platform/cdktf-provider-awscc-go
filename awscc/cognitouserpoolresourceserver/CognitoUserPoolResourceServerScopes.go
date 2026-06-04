package cognitouserpoolresourceserver


type CognitoUserPoolResourceServerScopes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_resource_server#scope_description CognitoUserPoolResourceServer#scope_description}.
	ScopeDescription *string `field:"optional" json:"scopeDescription" yaml:"scopeDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_resource_server#scope_name CognitoUserPoolResourceServer#scope_name}.
	ScopeName *string `field:"optional" json:"scopeName" yaml:"scopeName"`
}

