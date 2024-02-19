package lambdafunction


type LambdaFunctionEphemeralStorage struct {
	// The amount of ephemeral storage that your function has access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#size LambdaFunction#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

