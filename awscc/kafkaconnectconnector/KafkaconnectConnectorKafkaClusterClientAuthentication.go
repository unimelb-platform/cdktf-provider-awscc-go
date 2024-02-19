package kafkaconnectconnector


type KafkaconnectConnectorKafkaClusterClientAuthentication struct {
	// The type of client authentication used to connect to the Kafka cluster.
	//
	// Value NONE means that no client authentication is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#authentication_type KafkaconnectConnector#authentication_type}
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
}

