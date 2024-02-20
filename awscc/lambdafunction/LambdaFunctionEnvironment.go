package lambdafunction


type LambdaFunctionEnvironment struct {
	// Environment variable key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#variables LambdaFunction#variables}
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

