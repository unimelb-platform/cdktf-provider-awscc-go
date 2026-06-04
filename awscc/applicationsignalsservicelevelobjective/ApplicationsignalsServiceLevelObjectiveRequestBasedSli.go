package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveRequestBasedSli struct {
	// The arithmetic operation used when comparing the specified metric to the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective#comparison_operator ApplicationsignalsServiceLevelObjective#comparison_operator}
	ComparisonOperator *string `field:"optional" json:"comparisonOperator" yaml:"comparisonOperator"`
	// The value that the SLI metric is compared to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective#metric_threshold ApplicationsignalsServiceLevelObjective#metric_threshold}
	MetricThreshold *float64 `field:"optional" json:"metricThreshold" yaml:"metricThreshold"`
	// This structure contains the information about the metric that is used for a request-based SLO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/applicationsignals_service_level_objective#request_based_sli_metric ApplicationsignalsServiceLevelObjective#request_based_sli_metric}
	RequestBasedSliMetric *ApplicationsignalsServiceLevelObjectiveRequestBasedSliRequestBasedSliMetric `field:"optional" json:"requestBasedSliMetric" yaml:"requestBasedSliMetric"`
}

