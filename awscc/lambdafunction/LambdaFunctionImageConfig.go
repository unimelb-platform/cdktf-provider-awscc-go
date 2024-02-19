package lambdafunction


type LambdaFunctionImageConfig struct {
	// Command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#command LambdaFunction#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// EntryPoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#entry_point LambdaFunction#entry_point}
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// WorkingDirectory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#working_directory LambdaFunction#working_directory}
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

