package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActions struct {
	// Specifies one or more supported actions to receive notifications when the alarm state changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#alarm_actions IoteventsAlarmModel#alarm_actions}
	AlarmActions interface{} `field:"optional" json:"alarmActions" yaml:"alarmActions"`
}

