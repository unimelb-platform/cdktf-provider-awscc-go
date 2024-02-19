package lambdafunction


type LambdaFunctionRuntimeManagementConfig struct {
	// Trigger for runtime update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#update_runtime_on LambdaFunction#update_runtime_on}
	UpdateRuntimeOn *string `field:"required" json:"updateRuntimeOn" yaml:"updateRuntimeOn"`
	// Unique identifier for a runtime version arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#runtime_version_arn LambdaFunction#runtime_version_arn}
	RuntimeVersionArn *string `field:"optional" json:"runtimeVersionArn" yaml:"runtimeVersionArn"`
}

