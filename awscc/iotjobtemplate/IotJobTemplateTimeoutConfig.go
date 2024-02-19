package iotjobtemplate


type IotJobTemplateTimeoutConfig struct {
	// Specifies the amount of time, in minutes, this device has to finish execution of this job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#in_progress_timeout_in_minutes IotJobTemplate#in_progress_timeout_in_minutes}
	InProgressTimeoutInMinutes *float64 `field:"required" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

