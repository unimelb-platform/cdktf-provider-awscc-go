package datasynctask


type DatasyncTaskTaskReportConfig struct {
	// Specifies where DataSync uploads your task report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#destination DatasyncTask#destination}
	Destination *DatasyncTaskTaskReportConfigDestination `field:"required" json:"destination" yaml:"destination"`
	// Specifies the type of task report that you want.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#output_type DatasyncTask#output_type}
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// Specifies whether your task report includes the new version of each object transferred into an S3 bucket, this only applies if you enable versioning on your bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#object_version_ids DatasyncTask#object_version_ids}
	ObjectVersionIds *string `field:"optional" json:"objectVersionIds" yaml:"objectVersionIds"`
	// Customizes the reporting level for aspects of your task report.
	//
	// For example, your report might generally only include errors, but you could specify that you want a list of successes and errors just for the files that Datasync attempted to delete in your destination location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#overrides DatasyncTask#overrides}
	Overrides *DatasyncTaskTaskReportConfigOverrides `field:"optional" json:"overrides" yaml:"overrides"`
	// Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#report_level DatasyncTask#report_level}
	ReportLevel *string `field:"optional" json:"reportLevel" yaml:"reportLevel"`
}

