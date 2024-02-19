package appstreamappblock


type AppstreamAppBlockSourceS3Location struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#s3_bucket AppstreamAppBlock#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_app_block#s3_key AppstreamAppBlock#s3_key}.
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
}

