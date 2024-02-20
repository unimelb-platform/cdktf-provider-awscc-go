package quicksightdataset


type QuicksightDataSetPhysicalTableMapS3Source struct {
	// <p>The amazon Resource Name (ARN) for the data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}
	DataSourceArn *string `field:"required" json:"dataSourceArn" yaml:"dataSourceArn"`
	// <p>A physical table type for as S3 data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#input_columns QuicksightDataSet#input_columns}
	InputColumns interface{} `field:"required" json:"inputColumns" yaml:"inputColumns"`
	// <p>Information about the format for a source file or files.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#upload_settings QuicksightDataSet#upload_settings}
	UploadSettings *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings `field:"optional" json:"uploadSettings" yaml:"uploadSettings"`
}

