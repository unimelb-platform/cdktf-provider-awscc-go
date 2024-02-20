package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublish struct {
	// The MQTT topic of the message.
	//
	// You can use a string expression that includes variables (`$variable.<variable-name>`) and input values (`$input.<input-name>.<path-to-datum>`) as the topic string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#mqtt_topic IoteventsDetectorModel#mqtt_topic}
	MqttTopic *string `field:"required" json:"mqttTopic" yaml:"mqttTopic"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsIotTopicPublishPayload `field:"optional" json:"payload" yaml:"payload"`
}

