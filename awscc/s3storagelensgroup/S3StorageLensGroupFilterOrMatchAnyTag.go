package s3storagelensgroup


type S3StorageLensGroupFilterOrMatchAnyTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_storage_lens_group#key S3StorageLensGroup#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3_storage_lens_group#value S3StorageLensGroup#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

