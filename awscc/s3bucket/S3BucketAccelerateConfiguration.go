package s3bucket


type S3BucketAccelerateConfiguration struct {
	// Configures the transfer acceleration state for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#acceleration_status S3Bucket#acceleration_status}
	AccelerationStatus *string `field:"required" json:"accelerationStatus" yaml:"accelerationStatus"`
}

