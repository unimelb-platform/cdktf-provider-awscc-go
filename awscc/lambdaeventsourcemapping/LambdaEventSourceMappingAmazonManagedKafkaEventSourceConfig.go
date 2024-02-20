package lambdaeventsourcemapping


type LambdaEventSourceMappingAmazonManagedKafkaEventSourceConfig struct {
	// The identifier for the Kafka Consumer Group to join.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#consumer_group_id LambdaEventSourceMapping#consumer_group_id}
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
}

