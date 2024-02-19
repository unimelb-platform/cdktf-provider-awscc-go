package databrewdataset


type DatabrewDatasetPathOptionsParametersDatasetParameter struct {
	// Parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#name DatabrewDataset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parameter type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#type DatabrewDataset#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Add the value of this parameter as a column in a dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#create_column DatabrewDataset#create_column}
	CreateColumn interface{} `field:"optional" json:"createColumn" yaml:"createColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#datetime_options DatabrewDataset#datetime_options}.
	DatetimeOptions *DatabrewDatasetPathOptionsParametersDatasetParameterDatetimeOptions `field:"optional" json:"datetimeOptions" yaml:"datetimeOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#filter DatabrewDataset#filter}.
	Filter *DatabrewDatasetPathOptionsParametersDatasetParameterFilter `field:"optional" json:"filter" yaml:"filter"`
}

