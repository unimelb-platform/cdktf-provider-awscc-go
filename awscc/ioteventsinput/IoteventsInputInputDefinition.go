package ioteventsinput


type IoteventsInputInputDefinition struct {
	// The attributes from the JSON payload that are made available by the input.
	//
	// Inputs are derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such message contains a JSON payload, and those attributes (and their paired values) specified here are available for use in the `condition` expressions used by detectors that monitor this input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_input#attributes IoteventsInput#attributes}
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
}

