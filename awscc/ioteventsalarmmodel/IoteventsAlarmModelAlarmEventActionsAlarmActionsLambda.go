package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsLambda struct {
	// The ARN of the Lambda function that is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#function_arn IoteventsAlarmModel#function_arn}
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsLambdaPayload `field:"optional" json:"payload" yaml:"payload"`
}

