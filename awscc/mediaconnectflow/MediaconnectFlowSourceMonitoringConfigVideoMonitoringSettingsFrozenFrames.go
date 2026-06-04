package mediaconnectflow


type MediaconnectFlowSourceMonitoringConfigVideoMonitoringSettingsFrozenFrames struct {
	// Indicates whether the FrozenFrames metric is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediaconnect_flow#state MediaconnectFlow#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// Specifies the number of consecutive seconds of a static image that triggers an event or alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediaconnect_flow#threshold_seconds MediaconnectFlow#threshold_seconds}
	ThresholdSeconds *float64 `field:"optional" json:"thresholdSeconds" yaml:"thresholdSeconds"`
}

