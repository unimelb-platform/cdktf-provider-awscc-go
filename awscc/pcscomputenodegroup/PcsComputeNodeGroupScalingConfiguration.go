package pcscomputenodegroup


type PcsComputeNodeGroupScalingConfiguration struct {
	// The upper bound of the number of instances allowed in the compute fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_compute_node_group#max_instance_count PcsComputeNodeGroup#max_instance_count}
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// The lower bound of the number of instances allowed in the compute fleet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_compute_node_group#min_instance_count PcsComputeNodeGroup#min_instance_count}
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

