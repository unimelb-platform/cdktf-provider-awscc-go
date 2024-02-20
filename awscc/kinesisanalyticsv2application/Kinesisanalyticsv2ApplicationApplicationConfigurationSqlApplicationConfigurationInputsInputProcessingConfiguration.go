package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputProcessingConfiguration struct {
	// The InputLambdaProcessor that is used to preprocess the records in the stream before being processed by your application code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#input_lambda_processor Kinesisanalyticsv2Application#input_lambda_processor}
	InputLambdaProcessor *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputProcessingConfigurationInputLambdaProcessor `field:"optional" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}

