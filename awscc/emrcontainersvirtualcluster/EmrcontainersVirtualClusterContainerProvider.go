package emrcontainersvirtualcluster


type EmrcontainersVirtualClusterContainerProvider struct {
	// The ID of the container cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#id EmrcontainersVirtualCluster#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#info EmrcontainersVirtualCluster#info}.
	Info *EmrcontainersVirtualClusterContainerProviderInfo `field:"required" json:"info" yaml:"info"`
	// The type of the container provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/emrcontainers_virtual_cluster#type EmrcontainersVirtualCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

