package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorAnomalyDetectorConfig struct {
	// Frequency of anomaly detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#anomaly_detector_frequency LookoutmetricsAnomalyDetector#anomaly_detector_frequency}
	AnomalyDetectorFrequency *string `field:"required" json:"anomalyDetectorFrequency" yaml:"anomalyDetectorFrequency"`
}

