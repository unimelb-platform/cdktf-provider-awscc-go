package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#key_arn KinesisfirehoseDeliveryStream#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#key_type KinesisfirehoseDeliveryStream#key_type}.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

