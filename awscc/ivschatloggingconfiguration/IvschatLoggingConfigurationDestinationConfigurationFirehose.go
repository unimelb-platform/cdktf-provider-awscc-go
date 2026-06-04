package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfigurationFirehose struct {
	// Name of the Amazon Kinesis Firehose delivery stream where chat activity will be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_logging_configuration#delivery_stream_name IvschatLoggingConfiguration#delivery_stream_name}
	DeliveryStreamName *string `field:"optional" json:"deliveryStreamName" yaml:"deliveryStreamName"`
}

