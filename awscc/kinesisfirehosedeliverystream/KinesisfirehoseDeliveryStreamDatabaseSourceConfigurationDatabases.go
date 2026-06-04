package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabases struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#exclude KinesisfirehoseDeliveryStream#exclude}.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisfirehose_delivery_stream#include KinesisfirehoseDeliveryStream#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

