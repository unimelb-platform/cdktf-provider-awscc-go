package iotsoftwarepackageversion


type IotSoftwarePackageVersionArtifact struct {
	// The Amazon S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_software_package_version#s3_location IotSoftwarePackageVersion#s3_location}
	S3Location *IotSoftwarePackageVersionArtifactS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

