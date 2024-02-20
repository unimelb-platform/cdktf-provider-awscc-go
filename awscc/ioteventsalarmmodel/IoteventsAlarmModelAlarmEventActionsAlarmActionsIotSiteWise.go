package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWise struct {
	// The ID of the asset that has the specified property. You can specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#asset_id IoteventsAlarmModel#asset_id}
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// A unique identifier for this entry.
	//
	// You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier. You can also specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#entry_id IoteventsAlarmModel#entry_id}
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The alias of the asset property. You can also specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#property_alias IoteventsAlarmModel#property_alias}
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset property. You can specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#property_id IoteventsAlarmModel#property_id}
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// A structure that contains value information. For more information, see [AssetPropertyValue](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetPropertyValue.html) in the *AWS IoT SiteWise API Reference*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#property_value IoteventsAlarmModel#property_value}
	PropertyValue *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotSiteWisePropertyValue `field:"optional" json:"propertyValue" yaml:"propertyValue"`
}

