package lookoutmetricsanomalydetector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LookoutmetricsAnomalyDetectorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Configuration options for the AnomalyDetector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#anomaly_detector_config LookoutmetricsAnomalyDetector#anomaly_detector_config}
	AnomalyDetectorConfig *LookoutmetricsAnomalyDetectorAnomalyDetectorConfig `field:"required" json:"anomalyDetectorConfig" yaml:"anomalyDetectorConfig"`
	// List of metric sets for anomaly detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#metric_set_list LookoutmetricsAnomalyDetector#metric_set_list}
	MetricSetList interface{} `field:"required" json:"metricSetList" yaml:"metricSetList"`
	// A description for the AnomalyDetector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#anomaly_detector_description LookoutmetricsAnomalyDetector#anomaly_detector_description}
	AnomalyDetectorDescription *string `field:"optional" json:"anomalyDetectorDescription" yaml:"anomalyDetectorDescription"`
	// Name for the Amazon Lookout for Metrics Anomaly Detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#anomaly_detector_name LookoutmetricsAnomalyDetector#anomaly_detector_name}
	AnomalyDetectorName *string `field:"optional" json:"anomalyDetectorName" yaml:"anomalyDetectorName"`
	// KMS key used to encrypt the AnomalyDetector data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_anomaly_detector#kms_key_arn LookoutmetricsAnomalyDetector#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

