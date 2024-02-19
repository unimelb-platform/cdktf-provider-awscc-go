package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionJobResources struct {
	// Configuration for the cluster used to run model monitoring jobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#cluster_config SagemakerDataQualityJobDefinition#cluster_config}
	ClusterConfig *SagemakerDataQualityJobDefinitionJobResourcesClusterConfig `field:"required" json:"clusterConfig" yaml:"clusterConfig"`
}

