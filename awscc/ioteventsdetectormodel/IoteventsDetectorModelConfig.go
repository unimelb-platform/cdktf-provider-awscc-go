package ioteventsdetectormodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IoteventsDetectorModelConfig struct {
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
	// Information that defines how a detector operates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#detector_model_definition IoteventsDetectorModel#detector_model_definition}
	DetectorModelDefinition *IoteventsDetectorModelDetectorModelDefinition `field:"required" json:"detectorModelDefinition" yaml:"detectorModelDefinition"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#role_arn IoteventsDetectorModel#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A brief description of the detector model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#detector_model_description IoteventsDetectorModel#detector_model_description}
	DetectorModelDescription *string `field:"optional" json:"detectorModelDescription" yaml:"detectorModelDescription"`
	// The name of the detector model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#detector_model_name IoteventsDetectorModel#detector_model_name}
	DetectorModelName *string `field:"optional" json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#evaluation_method IoteventsDetectorModel#evaluation_method}
	EvaluationMethod *string `field:"optional" json:"evaluationMethod" yaml:"evaluationMethod"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new detector instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct detector instance, the device must send a message payload that contains the same attribute-value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#key IoteventsDetectorModel#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_detector_model#tags IoteventsDetectorModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

