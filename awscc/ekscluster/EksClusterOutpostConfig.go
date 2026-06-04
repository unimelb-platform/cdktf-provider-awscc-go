package ekscluster


type EksClusterOutpostConfig struct {
	// Specify the Instance type of the machines that should be used to create your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#control_plane_instance_type EksCluster#control_plane_instance_type}
	ControlPlaneInstanceType *string `field:"optional" json:"controlPlaneInstanceType" yaml:"controlPlaneInstanceType"`
	// Specify the placement group of the control plane machines for your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#control_plane_placement EksCluster#control_plane_placement}
	ControlPlanePlacement *EksClusterOutpostConfigControlPlanePlacement `field:"optional" json:"controlPlanePlacement" yaml:"controlPlanePlacement"`
	// Specify one or more Arn(s) of Outpost(s) on which you would like to create your cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/eks_cluster#outpost_arns EksCluster#outpost_arns}
	OutpostArns *[]*string `field:"optional" json:"outpostArns" yaml:"outpostArns"`
}

