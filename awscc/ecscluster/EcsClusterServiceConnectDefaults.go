package ecscluster


type EcsClusterServiceConnectDefaults struct {
	// Service Connect Namespace Name or ARN default for all services or tasks within this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#namespace EcsCluster#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

