package iotaccountauditconfiguration


type IotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckConfiguration struct {
	// The configValue for configuring audit checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_account_audit_configuration#cert_expiration_threshold_in_days IotAccountAuditConfiguration#cert_expiration_threshold_in_days}
	CertExpirationThresholdInDays *string `field:"optional" json:"certExpirationThresholdInDays" yaml:"certExpirationThresholdInDays"`
}

