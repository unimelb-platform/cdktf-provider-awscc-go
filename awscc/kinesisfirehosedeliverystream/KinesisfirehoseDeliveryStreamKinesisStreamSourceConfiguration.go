package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamKinesisStreamSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#kinesis_stream_arn KinesisfirehoseDeliveryStream#kinesis_stream_arn}.
	KinesisStreamArn *string `field:"required" json:"kinesisStreamArn" yaml:"kinesisStreamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#role_arn KinesisfirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

