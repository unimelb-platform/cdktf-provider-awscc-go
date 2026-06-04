package inspectorv2cisscanconfiguration


type Inspectorv2CisScanConfigurationScheduleWeekly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#days Inspectorv2CisScanConfiguration#days}.
	Days *[]*string `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/inspectorv2_cis_scan_configuration#start_time Inspectorv2CisScanConfiguration#start_time}.
	StartTime *Inspectorv2CisScanConfigurationScheduleWeeklyStartTime `field:"optional" json:"startTime" yaml:"startTime"`
}

