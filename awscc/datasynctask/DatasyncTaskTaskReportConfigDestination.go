package datasynctask


type DatasyncTaskTaskReportConfigDestination struct {
	// Specifies the Amazon S3 bucket where DataSync uploads your task report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#s3 DatasyncTask#s3}
	S3 *DatasyncTaskTaskReportConfigDestinationS3 `field:"optional" json:"s3" yaml:"s3"`
}

