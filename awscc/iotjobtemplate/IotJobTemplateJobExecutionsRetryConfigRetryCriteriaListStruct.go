package iotjobtemplate


type IotJobTemplateJobExecutionsRetryConfigRetryCriteriaListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#failure_type IotJobTemplate#failure_type}.
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#number_of_retries IotJobTemplate#number_of_retries}.
	NumberOfRetries *float64 `field:"optional" json:"numberOfRetries" yaml:"numberOfRetries"`
}

