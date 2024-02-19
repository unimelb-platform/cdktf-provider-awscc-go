package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2Payload struct {
	// The content of the payload.
	//
	// You can use a string expression that includes quoted strings (`'<string>'`), variables (`$variable.<variable-name>`), input values (`$input.<input-name>.<path-to-datum>`), string concatenations, and quoted strings that contain `${}` as the content. The recommended maximum size of a content expression is 1 KB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#content_expression IoteventsAlarmModel#content_expression}
	ContentExpression *string `field:"required" json:"contentExpression" yaml:"contentExpression"`
	// The value of the payload type can be either `STRING` or `JSON`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#type IoteventsAlarmModel#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

