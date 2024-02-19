package globalacceleratoraccelerator


type GlobalacceleratorAcceleratorTags struct {
	// Key of the tag. Value can be 1 to 127 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#key GlobalacceleratorAccelerator#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for the tag. Value can be 1 to 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/globalaccelerator_accelerator#value GlobalacceleratorAccelerator#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

