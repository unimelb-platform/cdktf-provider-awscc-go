package kafkaconnectconnector


type KafkaconnectConnectorLogDeliveryWorkerLogDeliveryFirehose struct {
	// Specifies whether the logs get sent to the specified Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#enabled KafkaconnectConnector#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Kinesis Data Firehose delivery stream that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#delivery_stream KafkaconnectConnector#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

