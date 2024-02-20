package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricListStruct struct {
	// Operator used to aggregate metric values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#aggregation_function LookoutmetricsAnomalyDetector#aggregation_function}
	AggregationFunction *string `field:"required" json:"aggregationFunction" yaml:"aggregationFunction"`
	// Name of a column in the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#metric_name LookoutmetricsAnomalyDetector#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#namespace LookoutmetricsAnomalyDetector#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

