package lambdafunction


type LambdaFunctionTags struct {
	// The key for this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_function#key LambdaFunction#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for this tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_function#value LambdaFunction#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

