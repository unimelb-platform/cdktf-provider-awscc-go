package datasynctask


type DatasyncTaskSchedule struct {
	// A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_task#schedule_expression DatasyncTask#schedule_expression}
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

