package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinition struct {
	// The state that is entered at the creation of each detector (instance).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#initial_state_name IoteventsDetectorModel#initial_state_name}
	InitialStateName *string `field:"required" json:"initialStateName" yaml:"initialStateName"`
	// Information about the states of the detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#states IoteventsDetectorModel#states}
	States interface{} `field:"required" json:"states" yaml:"states"`
}

