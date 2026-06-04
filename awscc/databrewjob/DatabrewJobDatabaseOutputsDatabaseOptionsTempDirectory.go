package databrewjob


type DatabrewJobDatabaseOutputsDatabaseOptionsTempDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#bucket DatabrewJob#bucket}.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#bucket_owner DatabrewJob#bucket_owner}.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/databrew_job#key DatabrewJob#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

