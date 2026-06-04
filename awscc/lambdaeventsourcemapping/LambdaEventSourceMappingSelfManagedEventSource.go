package lambdaeventsourcemapping


type LambdaEventSourceMappingSelfManagedEventSource struct {
	// The list of bootstrap servers for your Kafka brokers in the following format: ``"KafkaBootstrapServers": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"]``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#endpoints LambdaEventSourceMapping#endpoints}
	Endpoints *LambdaEventSourceMappingSelfManagedEventSourceEndpoints `field:"optional" json:"endpoints" yaml:"endpoints"`
}

