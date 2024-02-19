package quicksightdatasource


type QuicksightDataSourceDataSourceParametersS3ParametersManifestFileLocation struct {
	// <p>Amazon S3 bucket.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#bucket QuicksightDataSource#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// <p>Amazon S3 key that identifies an object.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#key QuicksightDataSource#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

