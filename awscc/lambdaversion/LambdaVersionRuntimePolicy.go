package lambdaversion


type LambdaVersionRuntimePolicy struct {
	// The runtime update mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_version#update_runtime_on LambdaVersion#update_runtime_on}
	UpdateRuntimeOn *string `field:"required" json:"updateRuntimeOn" yaml:"updateRuntimeOn"`
	// The ARN of the runtime the function is configured to use.
	//
	// If the runtime update mode is manual, the ARN is returned, otherwise null is returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_version#runtime_version_arn LambdaVersion#runtime_version_arn}
	RuntimeVersionArn *string `field:"optional" json:"runtimeVersionArn" yaml:"runtimeVersionArn"`
}

