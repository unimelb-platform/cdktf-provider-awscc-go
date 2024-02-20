package databrewdataset


type DatabrewDatasetInputDatabaseInputDefinitionTempDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#bucket DatabrewDataset#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_dataset#key DatabrewDataset#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

