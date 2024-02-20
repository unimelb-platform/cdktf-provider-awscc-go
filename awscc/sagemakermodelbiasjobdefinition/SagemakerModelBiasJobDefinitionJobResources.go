package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionJobResources struct {
	// Configuration for the cluster used to run model monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#cluster_config SagemakerModelBiasJobDefinition#cluster_config}
	ClusterConfig *SagemakerModelBiasJobDefinitionJobResourcesClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

