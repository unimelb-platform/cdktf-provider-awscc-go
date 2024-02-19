package ioteventsdetectormodel


type IoteventsDetectorModelTags struct {
	// Key of the Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#key IoteventsDetectorModel#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#value IoteventsDetectorModel#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

