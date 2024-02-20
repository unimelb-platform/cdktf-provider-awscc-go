package databrewdataset


type DatabrewDatasetFormatOptions struct {
	// Csv options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#csv DatabrewDataset#csv}
	Csv *DatabrewDatasetFormatOptionsCsv `field:"optional" json:"csv" yaml:"csv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#excel DatabrewDataset#excel}.
	Excel *DatabrewDatasetFormatOptionsExcel `field:"optional" json:"excel" yaml:"excel"`
	// Json options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#json DatabrewDataset#json}
	Json *DatabrewDatasetFormatOptionsJson `field:"optional" json:"json" yaml:"json"`
}

