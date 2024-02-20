package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamMskSourceConfigurationAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#connectivity KinesisfirehoseDeliveryStream#connectivity}.
	Connectivity *string `field:"required" json:"connectivity" yaml:"connectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisfirehose_delivery_stream#role_arn KinesisfirehoseDeliveryStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

