package iotanalyticsdataset


type IotanalyticsDatasetContentDeliveryRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_dataset#destination IotanalyticsDataset#destination}.
	Destination *IotanalyticsDatasetContentDeliveryRulesDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_dataset#entry_name IotanalyticsDataset#entry_name}.
	EntryName *string `field:"optional" json:"entryName" yaml:"entryName"`
}

