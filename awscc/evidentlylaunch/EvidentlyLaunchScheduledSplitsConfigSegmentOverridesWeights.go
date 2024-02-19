package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfigSegmentOverridesWeights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#group_name EvidentlyLaunch#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#split_weight EvidentlyLaunch#split_weight}.
	SplitWeight *float64 `field:"required" json:"splitWeight" yaml:"splitWeight"`
}

