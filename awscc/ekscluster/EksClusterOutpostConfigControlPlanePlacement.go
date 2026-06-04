package ekscluster


type EksClusterOutpostConfigControlPlanePlacement struct {
	// Specify the placement group name of the control place machines for your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#group_name EksCluster#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
}

