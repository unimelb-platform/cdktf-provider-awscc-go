package certificatemanageraccount


type CertificatemanagerAccountExpiryEventsConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/certificatemanager_account#days_before_expiry CertificatemanagerAccount#days_before_expiry}.
	DaysBeforeExpiry *float64 `field:"optional" json:"daysBeforeExpiry" yaml:"daysBeforeExpiry"`
}

