package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActions struct {
	// Information needed to clear the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#clear_timer IoteventsDetectorModel#clear_timer}
	ClearTimer *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsClearTimer `field:"optional" json:"clearTimer" yaml:"clearTimer"`
	// Writes to the DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). One column of the DynamoDB table receives all attribute-value pairs in the payload that you specify. For more information, see [Actions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-event-actions.html) in *AWS IoT Events Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#dynamo_db IoteventsDetectorModel#dynamo_db}
	DynamoDb *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Defines an action to write to the Amazon DynamoDB table that you created.
	//
	// The default action payload contains all attribute-value pairs that have the information about the detector model instance and the event that triggered the action. You can also customize the [payload](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Payload.html). A separate column of the DynamoDB table receives one attribute-value pair in the payload that you specify.
	//
	// You can use expressions for parameters that are strings. For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#dynamo_d_bv_2 IoteventsDetectorModel#dynamo_d_bv_2}
	DynamoDBv2 *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2 `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#firehose IoteventsDetectorModel#firehose}
	Firehose *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Sends an AWS IoT Events input, passing in information about the detector model instance and the event that triggered the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#iot_events IoteventsDetectorModel#iot_events}
	IotEvents *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#iot_site_wise IoteventsDetectorModel#iot_site_wise}
	IotSiteWise *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWise `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Information required to publish the MQTT message through the AWS IoT message broker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#iot_topic_publish IoteventsDetectorModel#iot_topic_publish}
	IotTopicPublish *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotTopicPublish `field:"optional" json:"iotTopicPublish" yaml:"iotTopicPublish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#lambda IoteventsDetectorModel#lambda}.
	Lambda *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Information required to reset the timer.
	//
	// The timer is reset to the previously evaluated result of the duration. The duration expression isn't reevaluated when you reset the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#reset_timer IoteventsDetectorModel#reset_timer}
	ResetTimer *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsResetTimer `field:"optional" json:"resetTimer" yaml:"resetTimer"`
	// Information needed to set the timer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#set_timer IoteventsDetectorModel#set_timer}
	SetTimer *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetTimer `field:"optional" json:"setTimer" yaml:"setTimer"`
	// Information about the variable and its new value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#set_variable IoteventsDetectorModel#set_variable}
	SetVariable *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSetVariable `field:"optional" json:"setVariable" yaml:"setVariable"`
	// Information required to publish the Amazon SNS message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#sns IoteventsDetectorModel#sns}
	Sns *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSns `field:"optional" json:"sns" yaml:"sns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#sqs IoteventsDetectorModel#sqs}.
	Sqs *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsSqs `field:"optional" json:"sqs" yaml:"sqs"`
}

