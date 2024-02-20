package timestreamtable


type TimestreamTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation struct {
	// S3 configuration for location to store rejections from magnetic store writes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/timestream_table#s3_configuration TimestreamTable#s3_configuration}
	S3Configuration *TimestreamTableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

