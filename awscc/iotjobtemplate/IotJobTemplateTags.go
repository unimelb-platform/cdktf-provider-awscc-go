package iotjobtemplate


type IotJobTemplateTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_job_template#key IotJobTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_job_template#value IotJobTemplate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

