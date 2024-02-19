package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionModelBiasBaselineConfig struct {
	// The name of a processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#baselining_job_name SagemakerModelBiasJobDefinition#baselining_job_name}
	BaseliningJobName *string `field:"optional" json:"baseliningJobName" yaml:"baseliningJobName"`
	// The baseline constraints resource for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#constraints_resource SagemakerModelBiasJobDefinition#constraints_resource}
	ConstraintsResource *SagemakerModelBiasJobDefinitionModelBiasBaselineConfigConstraintsResource `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
}

