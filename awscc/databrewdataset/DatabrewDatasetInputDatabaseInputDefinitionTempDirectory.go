package databrewdataset


type DatabrewDatasetInputDatabaseInputDefinitionTempDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_dataset#bucket DatabrewDataset#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Bucket owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_dataset#bucket_owner DatabrewDataset#bucket_owner}
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_dataset#key DatabrewDataset#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

