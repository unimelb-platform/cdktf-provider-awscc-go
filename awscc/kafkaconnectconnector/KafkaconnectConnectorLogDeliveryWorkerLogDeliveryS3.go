package kafkaconnectconnector


type KafkaconnectConnectorLogDeliveryWorkerLogDeliveryS3 struct {
	// Specifies whether the logs get sent to the specified Amazon S3 destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#enabled KafkaconnectConnector#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the S3 bucket that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#bucket KafkaconnectConnector#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#prefix KafkaconnectConnector#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

