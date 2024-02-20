package iotjobtemplate


type IotJobTemplateAbortConfigCriteriaListStruct struct {
	// The type of job action to take to initiate the job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#action IotJobTemplate#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The type of job execution failures that can initiate a job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#failure_type IotJobTemplate#failure_type}
	FailureType *string `field:"required" json:"failureType" yaml:"failureType"`
	// The minimum number of things which must receive job execution notifications before the job can be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#min_number_of_executed_things IotJobTemplate#min_number_of_executed_things}
	MinNumberOfExecutedThings *float64 `field:"required" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// The minimum percentage of job execution failures that must occur to initiate the job abort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#threshold_percentage IotJobTemplate#threshold_percentage}
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

