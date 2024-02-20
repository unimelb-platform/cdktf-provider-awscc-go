package kafkaconnectconnector


type KafkaconnectConnectorLogDelivery struct {
	// Specifies where worker logs are delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kafkaconnect_connector#worker_log_delivery KafkaconnectConnector#worker_log_delivery}
	WorkerLogDelivery *KafkaconnectConnectorLogDeliveryWorkerLogDelivery `field:"required" json:"workerLogDelivery" yaml:"workerLogDelivery"`
}

