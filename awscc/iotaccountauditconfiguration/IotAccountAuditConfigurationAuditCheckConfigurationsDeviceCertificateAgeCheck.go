package iotaccountauditconfiguration


type IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheck struct {
	// A structure containing the configName and corresponding configValue for configuring audit checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_account_audit_configuration#configuration IotAccountAuditConfiguration#configuration}
	Configuration *IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheckConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// True if the check is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_account_audit_configuration#enabled IotAccountAuditConfiguration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

