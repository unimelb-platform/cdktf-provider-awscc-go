package ivsrecordingconfiguration


type IvsRecordingConfigurationDestinationConfiguration struct {
	// Recording S3 Destination Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ivs_recording_configuration#s3 IvsRecordingConfiguration#s3}
	S3 *IvsRecordingConfigurationDestinationConfigurationS3 `field:"optional" json:"s3" yaml:"s3"`
}

