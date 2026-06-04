package pcscomputenodegroup


type PcsComputeNodeGroupSlurmConfiguration struct {
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_compute_node_group#slurm_custom_settings PcsComputeNodeGroup#slurm_custom_settings}
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

