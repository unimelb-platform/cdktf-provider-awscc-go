package lambdafunction


type LambdaFunctionTracingConfig struct {
	// The tracing mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#mode LambdaFunction#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

