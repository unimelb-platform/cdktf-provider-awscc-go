package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValue struct {
	// A structure that contains an asset property value.
	//
	// For more information, see [Variant](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_Variant.html) in the *AWS IoT SiteWise API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#value IoteventsAlarmModel#value}
	Value *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValueValue `field:"required" json:"value" yaml:"value"`
	// The quality of the asset property value.
	//
	// The value must be `GOOD`, `BAD`, or `UNCERTAIN`. You can also specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#quality IoteventsAlarmModel#quality}
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// A structure that contains timestamp information. For more information, see [TimeInNanos](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_TimeInNanos.html) in the *AWS IoT SiteWise API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#timestamp IoteventsAlarmModel#timestamp}
	Timestamp *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValueTimestamp `field:"optional" json:"timestamp" yaml:"timestamp"`
}

