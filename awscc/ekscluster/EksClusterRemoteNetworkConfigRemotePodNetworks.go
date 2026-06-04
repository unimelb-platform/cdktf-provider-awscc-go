package ekscluster


type EksClusterRemoteNetworkConfigRemotePodNetworks struct {
	// Specifies the list of remote pod CIDRs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#cidrs EksCluster#cidrs}
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

