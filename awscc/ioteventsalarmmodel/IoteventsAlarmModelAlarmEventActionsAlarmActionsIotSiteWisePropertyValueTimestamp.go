package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValueTimestamp struct {
	// The nanosecond offset converted from ``timeInSeconds``. The valid range is between 0-999999999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_alarm_model#offset_in_nanos IoteventsAlarmModel#offset_in_nanos}
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
	// The timestamp, in seconds, in the Unix epoch format. The valid range is between 1-31556889864403199.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_alarm_model#time_in_seconds IoteventsAlarmModel#time_in_seconds}
	TimeInSeconds *string `field:"optional" json:"timeInSeconds" yaml:"timeInSeconds"`
}

