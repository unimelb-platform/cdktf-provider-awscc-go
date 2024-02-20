package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionNetworkConfig struct {
	// Whether to encrypt all communications between distributed processing jobs.
	//
	// Choose True to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#enable_inter_container_traffic_encryption SagemakerModelExplainabilityJobDefinition#enable_inter_container_traffic_encryption}
	EnableInterContainerTrafficEncryption interface{} `field:"optional" json:"enableInterContainerTrafficEncryption" yaml:"enableInterContainerTrafficEncryption"`
	// Whether to allow inbound and outbound network calls to and from the containers used for the processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#enable_network_isolation SagemakerModelExplainabilityJobDefinition#enable_network_isolation}
	EnableNetworkIsolation interface{} `field:"optional" json:"enableNetworkIsolation" yaml:"enableNetworkIsolation"`
	// Specifies a VPC that your training jobs and hosted models have access to.
	//
	// Control access to and from your training and model containers by configuring the VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#vpc_config SagemakerModelExplainabilityJobDefinition#vpc_config}
	VpcConfig *SagemakerModelExplainabilityJobDefinitionNetworkConfigVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

