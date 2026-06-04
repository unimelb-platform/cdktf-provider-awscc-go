package iotthinggroup


type IotThingGroupThingGroupProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#attribute_payload IotThingGroup#attribute_payload}.
	AttributePayload *IotThingGroupThingGroupPropertiesAttributePayload `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_thing_group#thing_group_description IotThingGroup#thing_group_description}.
	ThingGroupDescription *string `field:"optional" json:"thingGroupDescription" yaml:"thingGroupDescription"`
}

