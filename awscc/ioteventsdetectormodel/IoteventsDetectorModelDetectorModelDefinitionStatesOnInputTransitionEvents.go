package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEvents struct {
	// A Boolean expression that when `TRUE` causes the `actions` to be performed and the `nextState` to be entered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#condition IoteventsDetectorModel#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// The name of the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#event_name IoteventsDetectorModel#event_name}
	EventName *string `field:"required" json:"eventName" yaml:"eventName"`
	// The next state to enter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#next_state IoteventsDetectorModel#next_state}
	NextState *string `field:"required" json:"nextState" yaml:"nextState"`
	// The actions to be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#actions IoteventsDetectorModel#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
}

