package groundstationmissionprofile


type GroundstationMissionProfileStreamsKmsKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_mission_profile#kms_alias_arn GroundstationMissionProfile#kms_alias_arn}.
	KmsAliasArn *string `field:"optional" json:"kmsAliasArn" yaml:"kmsAliasArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/groundstation_mission_profile#kms_key_arn GroundstationMissionProfile#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

