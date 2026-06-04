package syntheticscanary


type SyntheticsCanaryVisualReference struct {
	// Canary run id to be used as base reference for visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/synthetics_canary#base_canary_run_id SyntheticsCanary#base_canary_run_id}
	BaseCanaryRunId *string `field:"optional" json:"baseCanaryRunId" yaml:"baseCanaryRunId"`
	// List of screenshots used as base reference for visual testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/synthetics_canary#base_screenshots SyntheticsCanary#base_screenshots}
	BaseScreenshots interface{} `field:"optional" json:"baseScreenshots" yaml:"baseScreenshots"`
}

