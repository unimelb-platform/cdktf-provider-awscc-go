package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSourceEndpoints struct {
	// A list of Kafka server endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#kafka_bootstrap_servers LambdaEventSourceMapping#kafka_bootstrap_servers}
	KafkaBootstrapServers *[]*string `field:"optional" json:"kafkaBootstrapServers" yaml:"kafkaBootstrapServers"`
}

