package mskcluster


type MskClusterLoggingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/msk_cluster#broker_logs MskCluster#broker_logs}.
	BrokerLogs *MskClusterLoggingInfoBrokerLogs `field:"required" json:"brokerLogs" yaml:"brokerLogs"`
}

