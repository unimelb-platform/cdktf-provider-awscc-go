package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEvents struct {
	// The name of the AWS IoT Events input where the data is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#input_name IoteventsDetectorModel#input_name}
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotEventsPayload `field:"optional" json:"payload" yaml:"payload"`
}

