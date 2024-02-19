package frauddetectordetector


type FrauddetectorDetectorEventTypeLabels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#arn FrauddetectorDetector#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The time when the label was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#created_time FrauddetectorDetector#created_time}
	CreatedTime *string `field:"optional" json:"createdTime" yaml:"createdTime"`
	// The description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#description FrauddetectorDetector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#inline FrauddetectorDetector#inline}.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// The time when the label was last updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#last_updated_time FrauddetectorDetector#last_updated_time}
	LastUpdatedTime *string `field:"optional" json:"lastUpdatedTime" yaml:"lastUpdatedTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#name FrauddetectorDetector#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with this label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#tags FrauddetectorDetector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

