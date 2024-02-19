package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsCustomComponentConfigurationConfigurationDetailsHanaPrometheusExporter struct {
	// A flag which indicates agreeing to install SAP HANA DB client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#agree_to_install_hanadb_client ApplicationinsightsApplication#agree_to_install_hanadb_client}
	AgreeToInstallHanadbClient interface{} `field:"required" json:"agreeToInstallHanadbClient" yaml:"agreeToInstallHanadbClient"`
	// The HANA DB port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#hana_port ApplicationinsightsApplication#hana_port}
	HanaPort *string `field:"required" json:"hanaPort" yaml:"hanaPort"`
	// The secret name which manages the HANA DB credentials e.g. {   "username": "<>",   "password": "<>" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#hana_secret_name ApplicationinsightsApplication#hana_secret_name}
	HanaSecretName *string `field:"required" json:"hanaSecretName" yaml:"hanaSecretName"`
	// HANA DB SID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#hanasid ApplicationinsightsApplication#hanasid}
	Hanasid *string `field:"required" json:"hanasid" yaml:"hanasid"`
	// Prometheus exporter port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#prometheus_port ApplicationinsightsApplication#prometheus_port}
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

