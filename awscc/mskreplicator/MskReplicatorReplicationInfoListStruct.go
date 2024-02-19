package mskreplicator


type MskReplicatorReplicationInfoListStruct struct {
	// Configuration relating to consumer group replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#consumer_group_replication MskReplicator#consumer_group_replication}
	ConsumerGroupReplication *MskReplicatorReplicationInfoListConsumerGroupReplication `field:"required" json:"consumerGroupReplication" yaml:"consumerGroupReplication"`
	// Amazon Resource Name of the source Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#source_kafka_cluster_arn MskReplicator#source_kafka_cluster_arn}
	SourceKafkaClusterArn *string `field:"required" json:"sourceKafkaClusterArn" yaml:"sourceKafkaClusterArn"`
	// The type of compression to use writing records to target Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#target_compression_type MskReplicator#target_compression_type}
	TargetCompressionType *string `field:"required" json:"targetCompressionType" yaml:"targetCompressionType"`
	// Amazon Resource Name of the target Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#target_kafka_cluster_arn MskReplicator#target_kafka_cluster_arn}
	TargetKafkaClusterArn *string `field:"required" json:"targetKafkaClusterArn" yaml:"targetKafkaClusterArn"`
	// Configuration relating to topic replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_replicator#topic_replication MskReplicator#topic_replication}
	TopicReplication *MskReplicatorReplicationInfoListTopicReplication `field:"required" json:"topicReplication" yaml:"topicReplication"`
}

