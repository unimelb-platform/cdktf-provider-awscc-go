package ioteventsalarmmodel


type IoteventsAlarmModelAlarmRule struct {
	// A rule that compares an input property value to a threshold value with a comparison operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#simple_rule IoteventsAlarmModel#simple_rule}
	SimpleRule *IoteventsAlarmModelAlarmRuleSimpleRule `field:"optional" json:"simpleRule" yaml:"simpleRule"`
}

