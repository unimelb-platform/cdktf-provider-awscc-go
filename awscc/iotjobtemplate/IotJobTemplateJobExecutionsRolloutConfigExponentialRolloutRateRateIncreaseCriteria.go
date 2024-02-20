package iotjobtemplate


type IotJobTemplateJobExecutionsRolloutConfigExponentialRolloutRateRateIncreaseCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#number_of_notified_things IotJobTemplate#number_of_notified_things}.
	NumberOfNotifiedThings *float64 `field:"optional" json:"numberOfNotifiedThings" yaml:"numberOfNotifiedThings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#number_of_succeeded_things IotJobTemplate#number_of_succeeded_things}.
	NumberOfSucceededThings *float64 `field:"optional" json:"numberOfSucceededThings" yaml:"numberOfSucceededThings"`
}

