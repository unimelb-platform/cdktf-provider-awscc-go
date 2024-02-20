package s3multiregionaccesspoint


type S3MultiRegionAccessPointRegions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point#bucket S3MultiRegionAccessPoint#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_multi_region_access_point#bucket_account_id S3MultiRegionAccessPoint#bucket_account_id}.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
}

