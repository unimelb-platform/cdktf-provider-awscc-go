package emrcontainersvirtualcluster


type EmrcontainersVirtualClusterContainerProviderInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#eks_info EmrcontainersVirtualCluster#eks_info}.
	EksInfo *EmrcontainersVirtualClusterContainerProviderInfoEksInfo `field:"required" json:"eksInfo" yaml:"eksInfo"`
}

