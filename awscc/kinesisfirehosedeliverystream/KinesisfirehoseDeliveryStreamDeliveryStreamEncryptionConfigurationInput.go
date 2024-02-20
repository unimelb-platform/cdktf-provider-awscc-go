package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#key_type KinesisfirehoseDeliveryStream#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#key_arn KinesisfirehoseDeliveryStream#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

