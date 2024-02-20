package imagebuilderimagepipeline


type ImagebuilderImagePipelineImageTestsConfiguration struct {
	// Defines if tests should be executed when building this image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_pipeline#image_tests_enabled ImagebuilderImagePipeline#image_tests_enabled}
	ImageTestsEnabled interface{} `field:"optional" json:"imageTestsEnabled" yaml:"imageTestsEnabled"`
	// The maximum time in minutes that tests are permitted to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_pipeline#timeout_minutes ImagebuilderImagePipeline#timeout_minutes}
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

