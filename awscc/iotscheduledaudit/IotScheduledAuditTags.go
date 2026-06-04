package iotscheduledaudit


type IotScheduledAuditTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_scheduled_audit#key IotScheduledAudit#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_scheduled_audit#value IotScheduledAudit#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

