package databrewdataset


type DatabrewDatasetPathOptionsFilesLimit struct {
	// Maximum number of files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#max_files DatabrewDataset#max_files}
	MaxFiles *float64 `field:"required" json:"maxFiles" yaml:"maxFiles"`
	// Order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#order DatabrewDataset#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Ordered by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#ordered_by DatabrewDataset#ordered_by}
	OrderedBy *string `field:"optional" json:"orderedBy" yaml:"orderedBy"`
}

