package ioteventsalarmmodel


type IoteventsAlarmModelAlarmCapabilitiesInitializationConfiguration struct {
	// The value must be TRUE or FALSE.
	//
	// If FALSE, all alarm instances created based on the alarm model are activated. The default value is TRUE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#disabled_on_initialization IoteventsAlarmModel#disabled_on_initialization}
	DisabledOnInitialization interface{} `field:"optional" json:"disabledOnInitialization" yaml:"disabledOnInitialization"`
}

