package datasynctask


type DatasyncTaskTaskReportConfigOverridesVerified struct {
	// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#report_level DatasyncTask#report_level}
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
}

