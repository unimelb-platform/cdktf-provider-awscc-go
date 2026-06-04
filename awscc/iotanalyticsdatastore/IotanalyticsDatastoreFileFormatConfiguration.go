package iotanalyticsdatastore


type IotanalyticsDatastoreFileFormatConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_datastore#json_configuration IotanalyticsDatastore#json_configuration}.
	JsonConfiguration *string `field:"optional" json:"jsonConfiguration" yaml:"jsonConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotanalytics_datastore#parquet_configuration IotanalyticsDatastore#parquet_configuration}.
	ParquetConfiguration *IotanalyticsDatastoreFileFormatConfigurationParquetConfiguration `field:"optional" json:"parquetConfiguration" yaml:"parquetConfiguration"`
}

