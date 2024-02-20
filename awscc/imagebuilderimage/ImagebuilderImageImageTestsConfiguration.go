package imagebuilderimage


type ImagebuilderImageImageTestsConfiguration struct {
	// ImageTestsEnabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image#image_tests_enabled ImagebuilderImage#image_tests_enabled}
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// TimeoutMinutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image#timeout_minutes ImagebuilderImage#timeout_minutes}
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

