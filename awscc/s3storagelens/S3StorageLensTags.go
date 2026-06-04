package s3storagelens


type S3StorageLensTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_storage_lens#key S3StorageLens#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_storage_lens#value S3StorageLens#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

