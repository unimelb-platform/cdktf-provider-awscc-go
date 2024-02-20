package batchjobdefinition


type BatchJobDefinitionContainerPropertiesVolumesEfsVolumeConfigurationAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#access_point_id BatchJobDefinition#access_point_id}.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#iam BatchJobDefinition#iam}.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}

