package lookoutmetricsanomalydetector


type LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#csv_format_descriptor LookoutmetricsAnomalyDetector#csv_format_descriptor}.
	CsvFormatDescriptor *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptorCsvFormatDescriptor `field:"optional" json:"csvFormatDescriptor" yaml:"csvFormatDescriptor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#json_format_descriptor LookoutmetricsAnomalyDetector#json_format_descriptor}.
	JsonFormatDescriptor *LookoutmetricsAnomalyDetectorMetricSetListMetricSourceS3SourceConfigFileFormatDescriptorJsonFormatDescriptor `field:"optional" json:"jsonFormatDescriptor" yaml:"jsonFormatDescriptor"`
}

