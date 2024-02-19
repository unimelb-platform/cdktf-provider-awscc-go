package databrewdataset


type DatabrewDatasetPathOptionsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#dataset_parameter DatabrewDataset#dataset_parameter}.
	DatasetParameter *DatabrewDatasetPathOptionsParametersDatasetParameter `field:"required" json:"datasetParameter" yaml:"datasetParameter"`
	// Parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#path_parameter_name DatabrewDataset#path_parameter_name}
	PathParameterName *string `field:"required" json:"pathParameterName" yaml:"pathParameterName"`
}

