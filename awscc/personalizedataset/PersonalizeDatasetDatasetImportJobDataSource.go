package personalizedataset


type PersonalizeDatasetDatasetImportJobDataSource struct {
	// The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/personalize_dataset#data_location PersonalizeDataset#data_location}
	DataLocation *string `field:"optional" json:"dataLocation" yaml:"dataLocation"`
}

