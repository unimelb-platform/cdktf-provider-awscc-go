package ioteventsdetectormodel


type IoteventsDetectorModelDetectorModelDefinitionStatesOnInputEventsActionsIotSiteWisePropertyValueTimestamp struct {
	// The nanosecond offset converted from `timeInSeconds`. The valid range is between `0-999999999`. You can also specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#time_in_seconds IoteventsDetectorModel#time_in_seconds}
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// The timestamp, in seconds, in the Unix epoch format.
	//
	// The valid range is between `1-31556889864403199`. You can also specify an expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#offset_in_nanos IoteventsDetectorModel#offset_in_nanos}
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}

