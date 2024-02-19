package iotjobtemplate


type IotJobTemplateAbortConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_job_template#criteria_list IotJobTemplate#criteria_list}.
	CriteriaList interface{} `field:"required" json:"criteriaList" yaml:"criteriaList"`
}

