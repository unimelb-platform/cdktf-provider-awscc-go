package lambdaversion


type LambdaVersionProvisionedConcurrencyConfig struct {
	// The amount of provisioned concurrency to allocate for the version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_version#provisioned_concurrent_executions LambdaVersion#provisioned_concurrent_executions}
	ProvisionedConcurrentExecutions *float64 `field:"required" json:"provisionedConcurrentExecutions" yaml:"provisionedConcurrentExecutions"`
}

