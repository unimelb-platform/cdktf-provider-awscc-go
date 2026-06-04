package guarddutydetector


type GuarddutyDetectorFeatures struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_detector#additional_configuration GuarddutyDetector#additional_configuration}.
	AdditionalConfiguration interface{} `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_detector#name GuarddutyDetector#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/guardduty_detector#status GuarddutyDetector#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

