package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamMskSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#authentication_configuration KinesisfirehoseDeliveryStream#authentication_configuration}.
	AuthenticationConfiguration *KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#msk_cluster_arn KinesisfirehoseDeliveryStream#msk_cluster_arn}.
	MskClusterArn *string `field:"required" json:"mskClusterArn" yaml:"mskClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#topic_name KinesisfirehoseDeliveryStream#topic_name}.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

