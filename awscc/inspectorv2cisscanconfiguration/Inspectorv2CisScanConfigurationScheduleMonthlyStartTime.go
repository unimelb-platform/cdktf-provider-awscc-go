package inspectorv2cisscanconfiguration


type Inspectorv2CisScanConfigurationScheduleMonthlyStartTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#time_of_day Inspectorv2CisScanConfiguration#time_of_day}.
	TimeOfDay *string `field:"optional" json:"timeOfDay" yaml:"timeOfDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#time_zone Inspectorv2CisScanConfiguration#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

