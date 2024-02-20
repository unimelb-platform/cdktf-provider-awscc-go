package lambdaeventinvokeconfig


type LambdaEventInvokeConfigDestinationConfigOnFailure struct {
	// The Amazon Resource Name (ARN) of the destination resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_invoke_config#destination LambdaEventInvokeConfig#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
}

