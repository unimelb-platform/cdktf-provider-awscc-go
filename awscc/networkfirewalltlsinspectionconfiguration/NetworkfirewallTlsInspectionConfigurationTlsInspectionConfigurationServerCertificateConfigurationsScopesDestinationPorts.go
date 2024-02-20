package networkfirewalltlsinspectionconfiguration


type NetworkfirewallTlsInspectionConfigurationTlsInspectionConfigurationServerCertificateConfigurationsScopesDestinationPorts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_tls_inspection_configuration#from_port NetworkfirewallTlsInspectionConfiguration#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_tls_inspection_configuration#to_port NetworkfirewallTlsInspectionConfiguration#to_port}.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

