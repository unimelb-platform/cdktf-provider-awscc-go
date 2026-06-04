package iotmitigationaction


type IotMitigationActionActionParamsAddThingsToThingGroupParams struct {
	// Specifies if this mitigation action can move the things that triggered the mitigation action out of one or more dynamic thing groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_mitigation_action#override_dynamic_groups IotMitigationAction#override_dynamic_groups}
	OverrideDynamicGroups interface{} `field:"optional" json:"overrideDynamicGroups" yaml:"overrideDynamicGroups"`
	// The list of groups to which you want to add the things that triggered the mitigation action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iot_mitigation_action#thing_group_names IotMitigationAction#thing_group_names}
	ThingGroupNames *[]*string `field:"optional" json:"thingGroupNames" yaml:"thingGroupNames"`
}

