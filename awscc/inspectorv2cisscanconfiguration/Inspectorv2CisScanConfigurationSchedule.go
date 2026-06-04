package inspectorv2cisscanconfiguration


type Inspectorv2CisScanConfigurationSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#daily Inspectorv2CisScanConfiguration#daily}.
	Daily *Inspectorv2CisScanConfigurationScheduleDaily `field:"optional" json:"daily" yaml:"daily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#monthly Inspectorv2CisScanConfiguration#monthly}.
	Monthly *Inspectorv2CisScanConfigurationScheduleMonthly `field:"optional" json:"monthly" yaml:"monthly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#one_time Inspectorv2CisScanConfiguration#one_time}.
	OneTime *string `field:"optional" json:"oneTime" yaml:"oneTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#weekly Inspectorv2CisScanConfiguration#weekly}.
	Weekly *Inspectorv2CisScanConfigurationScheduleWeekly `field:"optional" json:"weekly" yaml:"weekly"`
}

