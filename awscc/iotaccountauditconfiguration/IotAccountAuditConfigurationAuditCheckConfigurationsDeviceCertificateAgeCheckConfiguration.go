package iotaccountauditconfiguration


type IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateAgeCheckConfiguration struct {
	// The configValue for configuring audit checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_account_audit_configuration#cert_age_threshold_in_days IotAccountAuditConfiguration#cert_age_threshold_in_days}
	CertAgeThresholdInDays *string `field:"optional" json:"certAgeThresholdInDays" yaml:"certAgeThresholdInDays"`
}

