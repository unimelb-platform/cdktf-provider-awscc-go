package mskcluster


type MskClusterOpenMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/msk_cluster#prometheus MskCluster#prometheus}.
	Prometheus *MskClusterOpenMonitoringPrometheus `field:"optional" json:"prometheus" yaml:"prometheus"`
}

