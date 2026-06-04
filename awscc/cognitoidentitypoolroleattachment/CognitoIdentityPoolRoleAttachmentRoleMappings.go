package cognitoidentitypoolroleattachment


type CognitoIdentityPoolRoleAttachmentRoleMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#ambiguous_role_resolution CognitoIdentityPoolRoleAttachment#ambiguous_role_resolution}.
	AmbiguousRoleResolution *string `field:"optional" json:"ambiguousRoleResolution" yaml:"ambiguousRoleResolution"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#identity_provider CognitoIdentityPoolRoleAttachment#identity_provider}.
	IdentityProvider *string `field:"optional" json:"identityProvider" yaml:"identityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#rules_configuration CognitoIdentityPoolRoleAttachment#rules_configuration}.
	RulesConfiguration *CognitoIdentityPoolRoleAttachmentRoleMappingsRulesConfiguration `field:"optional" json:"rulesConfiguration" yaml:"rulesConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_identity_pool_role_attachment#type CognitoIdentityPoolRoleAttachment#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

