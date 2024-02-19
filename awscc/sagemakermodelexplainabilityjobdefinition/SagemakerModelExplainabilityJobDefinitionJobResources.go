package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionJobResources struct {
	// Configuration for the cluster used to run model monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#cluster_config SagemakerModelExplainabilityJobDefinition#cluster_config}
	ClusterConfig *SagemakerModelExplainabilityJobDefinitionJobResourcesClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

