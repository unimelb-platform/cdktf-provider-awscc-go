package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumesPersistentVolumeClaim struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#claim_name BatchJobDefinition#claim_name}.
	ClaimName *string `field:"optional" json:"claimName" yaml:"claimName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#read_only BatchJobDefinition#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

