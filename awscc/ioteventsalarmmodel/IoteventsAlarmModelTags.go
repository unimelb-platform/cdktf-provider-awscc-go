package ioteventsalarmmodel


type IoteventsAlarmModelTags struct {
	// Key of the Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#key IoteventsAlarmModel#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#value IoteventsAlarmModel#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

