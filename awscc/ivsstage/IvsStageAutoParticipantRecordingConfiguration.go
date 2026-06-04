package ivsstage


type IvsStageAutoParticipantRecordingConfiguration struct {
	// Types of media to be recorded. Default: AUDIO_VIDEO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_stage#media_types IvsStage#media_types}
	MediaTypes *[]*string `field:"optional" json:"mediaTypes" yaml:"mediaTypes"`
	// ARN of the StorageConfiguration resource to use for individual participant recording.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivs_stage#storage_configuration_arn IvsStage#storage_configuration_arn}
	StorageConfigurationArn *string `field:"optional" json:"storageConfigurationArn" yaml:"storageConfigurationArn"`
}

