package lambdaeventsourcemapping


type LambdaEventSourceMappingMetricsConfig struct {
	// The metrics you want your event source mapping to produce.
	//
	// Include ``EventCount`` to receive event source mapping metrics related to the number of events processed by your event source mapping. For more information about these metrics, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/lambda_event_source_mapping#metrics LambdaEventSourceMapping#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
}

