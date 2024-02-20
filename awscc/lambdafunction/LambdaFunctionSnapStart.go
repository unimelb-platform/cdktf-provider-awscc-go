package lambdafunction


type LambdaFunctionSnapStart struct {
	// Applying SnapStart setting on function resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#apply_on LambdaFunction#apply_on}
	ApplyOn *string `field:"required" json:"applyOn" yaml:"applyOn"`
}

