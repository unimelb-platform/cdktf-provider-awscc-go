package quicksightdashboard


type QuicksightDashboardDashboardPublishOptions struct {
	// <p>Ad hoc (one-time) filtering option.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#ad_hoc_filtering_option QuicksightDashboard#ad_hoc_filtering_option}
	AdHocFilteringOption *QuicksightDashboardDashboardPublishOptionsAdHocFilteringOption `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// <p>Export to .csv option.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#export_to_csv_option QuicksightDashboard#export_to_csv_option}
	ExportToCsvOption *QuicksightDashboardDashboardPublishOptionsExportToCsvOption `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// <p>Sheet controls option.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_dashboard#sheet_controls_option QuicksightDashboard#sheet_controls_option}
	SheetControlsOption *QuicksightDashboardDashboardPublishOptionsSheetControlsOption `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
}

