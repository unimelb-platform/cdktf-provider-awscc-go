package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfigSegmentOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evidently_launch#evaluation_order EvidentlyLaunch#evaluation_order}.
	EvaluationOrder *float64 `field:"optional" json:"evaluationOrder" yaml:"evaluationOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evidently_launch#segment EvidentlyLaunch#segment}.
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/evidently_launch#weights EvidentlyLaunch#weights}.
	Weights interface{} `field:"optional" json:"weights" yaml:"weights"`
}

