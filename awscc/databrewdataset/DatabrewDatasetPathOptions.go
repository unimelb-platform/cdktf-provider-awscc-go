package databrewdataset


type DatabrewDatasetPathOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#files_limit DatabrewDataset#files_limit}.
	FilesLimit *DatabrewDatasetPathOptionsFilesLimit `field:"optional" json:"filesLimit" yaml:"filesLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#last_modified_date_condition DatabrewDataset#last_modified_date_condition}.
	LastModifiedDateCondition *DatabrewDatasetPathOptionsLastModifiedDateCondition `field:"optional" json:"lastModifiedDateCondition" yaml:"lastModifiedDateCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#parameters DatabrewDataset#parameters}.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

