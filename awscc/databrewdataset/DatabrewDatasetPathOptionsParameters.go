package databrewdataset


type DatabrewDatasetPathOptionsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_dataset#dataset_parameter DatabrewDataset#dataset_parameter}.
	DatasetParameter *DatabrewDatasetPathOptionsParametersDatasetParameter `field:"optional" json:"datasetParameter" yaml:"datasetParameter"`
	// Parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_dataset#path_parameter_name DatabrewDataset#path_parameter_name}
	PathParameterName *string `field:"optional" json:"pathParameterName" yaml:"pathParameterName"`
}

