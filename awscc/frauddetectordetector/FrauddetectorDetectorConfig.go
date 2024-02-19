package frauddetectordetector

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrauddetectorDetectorConfig struct {
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
	// The ID of the detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#detector_id FrauddetectorDetector#detector_id}
	DetectorId *string `field:"required" json:"detectorId" yaml:"detectorId"`
	// The event type to associate this detector with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#event_type FrauddetectorDetector#event_type}
	EventType *FrauddetectorDetectorEventType `field:"required" json:"eventType" yaml:"eventType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#rules FrauddetectorDetector#rules}.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// The models to associate with this detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#associated_models FrauddetectorDetector#associated_models}
	AssociatedModels interface{} `field:"optional" json:"associatedModels" yaml:"associatedModels"`
	// The description of the detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#description FrauddetectorDetector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The desired detector version status for the detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#detector_version_status FrauddetectorDetector#detector_version_status}
	DetectorVersionStatus *string `field:"optional" json:"detectorVersionStatus" yaml:"detectorVersionStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#rule_execution_mode FrauddetectorDetector#rule_execution_mode}.
	RuleExecutionMode *string `field:"optional" json:"ruleExecutionMode" yaml:"ruleExecutionMode"`
	// Tags associated with this detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/frauddetector_detector#tags FrauddetectorDetector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

