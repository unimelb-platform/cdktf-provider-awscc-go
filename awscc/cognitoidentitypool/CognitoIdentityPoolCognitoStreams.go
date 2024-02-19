package cognitoidentitypool


type CognitoIdentityPoolCognitoStreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool#role_arn CognitoIdentityPool#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool#streaming_status CognitoIdentityPool#streaming_status}.
	StreamingStatus *string `field:"optional" json:"streamingStatus" yaml:"streamingStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_identity_pool#stream_name CognitoIdentityPool#stream_name}.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

