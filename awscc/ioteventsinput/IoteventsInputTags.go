package ioteventsinput


type IoteventsInputTags struct {
	// Key of the Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_input#key IoteventsInput#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value of the Tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_input#value IoteventsInput#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

