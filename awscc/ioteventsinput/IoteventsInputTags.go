package ioteventsinput


type IoteventsInputTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_input#key IoteventsInput#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_input#value IoteventsInput#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

