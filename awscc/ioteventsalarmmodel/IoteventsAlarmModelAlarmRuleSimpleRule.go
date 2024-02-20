package ioteventsalarmmodel


type IoteventsAlarmModelAlarmRuleSimpleRule struct {
	// The comparison operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#comparison_operator IoteventsAlarmModel#comparison_operator}
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value on the left side of the comparison operator.
	//
	// You can specify an AWS IoT Events input attribute as an input property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#input_property IoteventsAlarmModel#input_property}
	InputProperty *string `field:"required" json:"inputProperty" yaml:"inputProperty"`
	// The value on the right side of the comparison operator.
	//
	// You can enter a number or specify an AWS IoT Events input attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#threshold IoteventsAlarmModel#threshold}
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
}

