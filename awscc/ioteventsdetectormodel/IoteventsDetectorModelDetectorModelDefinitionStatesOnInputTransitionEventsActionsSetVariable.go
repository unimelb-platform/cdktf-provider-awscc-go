package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputTransitionEventsActionsSetVariable struct {
	// The new value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#value IoteventsDetectorModel#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#variable_name IoteventsDetectorModel#variable_name}
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
}

