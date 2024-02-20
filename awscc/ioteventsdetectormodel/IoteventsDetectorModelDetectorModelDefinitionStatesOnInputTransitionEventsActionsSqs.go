package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqs struct {
	// The URL of the SQS queue where the data is written.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#queue_url IoteventsDetectorModel#queue_url}
	QueueUrl *string `field:"required" json:"queueUrl" yaml:"queueUrl"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSqsPayload `field:"optional" json:"payload" yaml:"payload"`
	// Set this to `TRUE` if you want the data to be base-64 encoded before it is written to the queue.
	//
	// Otherwise, set this to `FALSE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#use_base_64 IoteventsDetectorModel#use_base_64}
	UseBase64 interface{} `field:"optional" json:"useBase64" yaml:"useBase64"`
}

