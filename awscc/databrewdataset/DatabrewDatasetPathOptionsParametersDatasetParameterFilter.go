package databrewdataset


type DatabrewDatasetPathOptionsParametersDatasetParameterFilter struct {
	// Filtering expression for a parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#expression DatabrewDataset#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#values_map DatabrewDataset#values_map}.
	ValuesMap interface{} `field:"required" json:"valuesMap" yaml:"valuesMap"`
}

