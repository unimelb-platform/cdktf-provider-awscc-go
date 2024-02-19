package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListTimestampColumn struct {
	// A timestamp format for the timestamps in the dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#column_format LookoutmetricsAnomalyDetector#column_format}
	ColumnFormat *string `field:"optional" json:"columnFormat" yaml:"columnFormat"`
	// Name of a column in the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#column_name LookoutmetricsAnomalyDetector#column_name}
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
}

