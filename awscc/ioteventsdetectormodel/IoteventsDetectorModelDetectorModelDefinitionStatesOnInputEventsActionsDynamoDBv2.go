package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2 struct {
	// The name of the DynamoDB table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#table_name IoteventsDetectorModel#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the detector model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#payload IoteventsDetectorModel#payload}
	Payload *IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsDynamoDBv2Payload `field:"optional" json:"payload" yaml:"payload"`
}

