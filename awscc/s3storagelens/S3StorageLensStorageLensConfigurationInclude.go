package s3storagelens


type S3StorageLensStorageLensConfigurationInclude struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#buckets S3StorageLens#buckets}.
	Buckets *[]*string `field:"optional" json:"buckets" yaml:"buckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#regions S3StorageLens#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

