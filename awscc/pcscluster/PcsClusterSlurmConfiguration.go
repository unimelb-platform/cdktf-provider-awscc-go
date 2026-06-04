package pcscluster


type PcsClusterSlurmConfiguration struct {
	// The accounting configuration includes configurable settings for Slurm accounting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_cluster#accounting PcsCluster#accounting}
	Accounting *PcsClusterSlurmConfigurationAccounting `field:"optional" json:"accounting" yaml:"accounting"`
	// The shared Slurm key for authentication, also known as the cluster secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_cluster#auth_key PcsCluster#auth_key}
	AuthKey *PcsClusterSlurmConfigurationAuthKey `field:"optional" json:"authKey" yaml:"authKey"`
	// The time before an idle node is scaled down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_cluster#scale_down_idle_time_in_seconds PcsCluster#scale_down_idle_time_in_seconds}
	ScaleDownIdleTimeInSeconds *float64 `field:"optional" json:"scaleDownIdleTimeInSeconds" yaml:"scaleDownIdleTimeInSeconds"`
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_cluster#slurm_custom_settings PcsCluster#slurm_custom_settings}
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

