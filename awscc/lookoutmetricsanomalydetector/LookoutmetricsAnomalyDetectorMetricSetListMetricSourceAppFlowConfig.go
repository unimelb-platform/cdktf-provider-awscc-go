package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceAppFlowConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#flow_name LookoutmetricsAnomalyDetector#flow_name}.
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#role_arn LookoutmetricsAnomalyDetector#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

