package lookoutmetricsalert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LookoutmetricsAlertConfig struct {
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
	// The action to be taken by the alert when an anomaly is detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#action LookoutmetricsAlert#action}
	Action *LookoutmetricsAlertAction `field:"required" json:"action" yaml:"action"`
	// A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#alert_sensitivity_threshold LookoutmetricsAlert#alert_sensitivity_threshold}
	AlertSensitivityThreshold *float64 `field:"required" json:"alertSensitivityThreshold" yaml:"alertSensitivityThreshold"`
	// The Amazon resource name (ARN) of the Anomaly Detector to alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#anomaly_detector_arn LookoutmetricsAlert#anomaly_detector_arn}
	AnomalyDetectorArn *string `field:"required" json:"anomalyDetectorArn" yaml:"anomalyDetectorArn"`
	// A description for the alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#alert_description LookoutmetricsAlert#alert_description}
	AlertDescription *string `field:"optional" json:"alertDescription" yaml:"alertDescription"`
	// The name of the alert. If not provided, a name is generated automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lookoutmetrics_alert#alert_name LookoutmetricsAlert#alert_name}
	AlertName *string `field:"optional" json:"alertName" yaml:"alertName"`
}

