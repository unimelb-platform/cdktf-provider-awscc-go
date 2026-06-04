package appsyncgraphqlapi


type AppsyncGraphQlApiAdditionalAuthenticationProvidersUserPoolConfig struct {
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_graph_ql_api#app_id_client_regex AppsyncGraphQlApi#app_id_client_regex}
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// The AWS Region in which the user pool was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_graph_ql_api#aws_region AppsyncGraphQlApi#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The user pool ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_graph_ql_api#user_pool_id AppsyncGraphQlApi#user_pool_id}
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

