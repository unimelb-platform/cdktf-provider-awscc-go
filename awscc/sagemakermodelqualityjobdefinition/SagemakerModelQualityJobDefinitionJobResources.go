package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionJobResources struct {
	// Configuration for the cluster used to run model monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#cluster_config SagemakerModelQualityJobDefinition#cluster_config}
	ClusterConfig *SagemakerModelQualityJobDefinitionJobResourcesClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

