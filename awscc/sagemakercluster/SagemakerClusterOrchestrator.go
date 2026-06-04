package sagemakercluster


type SagemakerClusterOrchestrator struct {
	// Specifies parameter(s) related to EKS as orchestrator, e.g. the EKS cluster nodes will attach to,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_cluster#eks SagemakerCluster#eks}
	Eks *SagemakerClusterOrchestratorEks `field:"optional" json:"eks" yaml:"eks"`
}

