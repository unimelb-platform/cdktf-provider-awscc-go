package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#file_format_descriptor LookoutmetricsAnomalyDetector#file_format_descriptor}.
	FileFormatDescriptor *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor `field:"required" json:"fileFormatDescriptor" yaml:"fileFormatDescriptor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#role_arn LookoutmetricsAnomalyDetector#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#historical_data_path_list LookoutmetricsAnomalyDetector#historical_data_path_list}.
	HistoricalDataPathList *[]*string `field:"optional" json:"historicalDataPathList" yaml:"historicalDataPathList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#templated_path_list LookoutmetricsAnomalyDetector#templated_path_list}.
	TemplatedPathList *[]*string `field:"optional" json:"templatedPathList" yaml:"templatedPathList"`
}

