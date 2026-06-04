package iotanalyticsdataset


type IotanalyticsDatasetContentDeliveryRulesDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_dataset#iot_events_destination_configuration IotanalyticsDataset#iot_events_destination_configuration}.
	IotEventsDestinationConfiguration *IotanalyticsDatasetContentDeliveryRulesDestinationIotEventsDestinationConfiguration `field:"optional" json:"iotEventsDestinationConfiguration" yaml:"iotEventsDestinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_dataset#s3_destination_configuration IotanalyticsDataset#s3_destination_configuration}.
	S3DestinationConfiguration *IotanalyticsDatasetContentDeliveryRulesDestinationS3DestinationConfiguration `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
}

