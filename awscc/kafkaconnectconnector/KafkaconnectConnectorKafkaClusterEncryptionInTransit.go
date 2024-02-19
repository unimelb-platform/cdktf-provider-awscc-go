package kafkaconnectconnector


type KafkaconnectConnectorKafkaClusterEncryptionInTransit struct {
	// The type of encryption in transit to the Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#encryption_type KafkaconnectConnector#encryption_type}
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
}

