package opensearchservicedomain


type OpensearchserviceDomainCognitoOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain#enabled OpensearchserviceDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain#identity_pool_id OpensearchserviceDomain#identity_pool_id}.
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain#role_arn OpensearchserviceDomain#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain#user_pool_id OpensearchserviceDomain#user_pool_id}.
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

