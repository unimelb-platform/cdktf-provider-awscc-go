package appstreamapplication


type AppstreamApplicationIconS3Location struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application#s3_bucket AppstreamApplication#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_application#s3_key AppstreamApplication#s3_key}.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}

