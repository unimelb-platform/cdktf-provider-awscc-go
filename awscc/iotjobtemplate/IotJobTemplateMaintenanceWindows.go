package iotjobtemplate


type IotJobTemplateMaintenanceWindows struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#duration_in_minutes IotJobTemplate#duration_in_minutes}.
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#start_time IotJobTemplate#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

