package kafkaconnectconnector


type KafkaconnectConnectorKafkaCluster struct {
	// Details of how to connect to an Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#apache_kafka_cluster KafkaconnectConnector#apache_kafka_cluster}
	ApacheKafkaCluster *KafkaconnectConnectorKafkaClusterApacheKafkaCluster `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

