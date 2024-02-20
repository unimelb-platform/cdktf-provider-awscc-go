package mskcluster


type MskClusterOpenMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#prometheus MskCluster#prometheus}.
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"required" json:"prometheus" yaml:"prometheus"`
}

