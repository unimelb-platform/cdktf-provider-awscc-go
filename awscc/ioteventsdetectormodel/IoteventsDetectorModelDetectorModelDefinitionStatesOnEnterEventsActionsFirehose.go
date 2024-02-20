package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehose struct {
	// The name of the Kinesis Data Firehose delivery stream where the data is written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#delivery_stream_name IoteventsDetectorModel#delivery_stream_name}
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnEnterEventsActionsFirehosePayload `field:"optional" json:"payload" yaml:"payload"`
	// A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#separator IoteventsDetectorModel#separator}
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}

