package guarddutydetector


type GuarddutyDetectorFeatures struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_detector#name GuarddutyDetector#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_detector#status GuarddutyDetector#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_detector#additional_configuration GuarddutyDetector#additional_configuration}.
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
}

