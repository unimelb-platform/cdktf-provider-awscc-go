package evidentlylaunch


type EvidentlyLaunchMetricMonitors struct {
	// The JSON path to reference the entity id in the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#entity_id_key EvidentlyLaunch#entity_id_key}
	EntityIdKey *string `field:"required" json:"entityIdKey" yaml:"entityIdKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#metric_name EvidentlyLaunch#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The JSON path to reference the numerical metric value in the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#value_key EvidentlyLaunch#value_key}
	ValueKey *string `field:"required" json:"valueKey" yaml:"valueKey"`
	// Event patterns have the same structure as the events they match.
	//
	// Rules use event patterns to select events. An event pattern either matches an event or it doesn't.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#event_pattern EvidentlyLaunch#event_pattern}
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/evidently_launch#unit_label EvidentlyLaunch#unit_label}.
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
}

