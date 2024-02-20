package ioteventsalarmmodel


type IoteventsAlarmModelAlarmCapabilities struct {
	// Specifies whether to get notified for alarm state changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#acknowledge_flow IoteventsAlarmModel#acknowledge_flow}
	AcknowledgeFlow *IoteventsAlarmModelAlarmCapabilitiesAcknowledgeFlow `field:"optional" json:"acknowledgeFlow" yaml:"acknowledgeFlow"`
	// Specifies the default alarm state. The configuration applies to all alarms that were created based on this alarm model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#initialization_configuration IoteventsAlarmModel#initialization_configuration}
	InitializationConfiguration *IoteventsAlarmModelAlarmCapabilitiesInitializationConfiguration `field:"optional" json:"initializationConfiguration" yaml:"initializationConfiguration"`
}

