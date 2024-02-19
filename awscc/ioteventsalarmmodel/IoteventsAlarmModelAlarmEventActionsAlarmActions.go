package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActions struct {
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the alarm model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#dynamo_db IoteventsAlarmModel#dynamo_db}
	DynamoDb *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the alarm model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
	//
	// You can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#dynamo_d_bv_2 IoteventsAlarmModel#dynamo_d_bv_2}
	DynamoDBv2 *IoteventsAlarmModelAlarmEventActionsAlarmActionsDynamoDBv2 `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the alarm model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#firehose IoteventsAlarmModel#firehose}
	Firehose *IoteventsAlarmModelAlarmEventActionsAlarmActionsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends an AWS IoT Events input, passing in information about the alarm model instance and the event that triggered the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#iot_events IoteventsAlarmModel#iot_events}
	IotEvents *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the alarm model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#iot_site_wise IoteventsAlarmModel#iot_site_wise}
	IotSiteWise *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWise `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Information required to publish the MQTT message through the AWS IoT message broker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#iot_topic_publish IoteventsAlarmModel#iot_topic_publish}
	IotTopicPublish *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotTopicPublish `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#lambda IoteventsAlarmModel#lambda}.
	Lambda *IoteventsAlarmModelAlarmEventActionsAlarmActionsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Information required to publish the Amazon SNS message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#sns IoteventsAlarmModel#sns}
	Sns *IoteventsAlarmModelAlarmEventActionsAlarmActionsSns `field:"optional" json:"sns" yaml:"sns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#sqs IoteventsAlarmModel#sqs}.
	Sqs *IoteventsAlarmModelAlarmEventActionsAlarmActionsSqs `field:"optional" json:"sqs" yaml:"sqs"`
}

