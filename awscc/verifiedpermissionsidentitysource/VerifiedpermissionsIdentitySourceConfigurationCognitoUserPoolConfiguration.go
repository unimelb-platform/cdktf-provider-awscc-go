package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_identity_source#user_pool_arn VerifiedpermissionsIdentitySource#user_pool_arn}.
	UserPoolArn *string `field:"required" json:"userPoolArn" yaml:"userPoolArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_identity_source#client_ids VerifiedpermissionsIdentitySource#client_ids}.
	ClientIds *[]*string `field:"optional" json:"clientIds" yaml:"clientIds"`
}

