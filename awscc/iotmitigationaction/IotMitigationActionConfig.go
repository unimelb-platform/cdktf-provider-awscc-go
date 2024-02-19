package iotmitigationaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotMitigationActionConfig struct {
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
	// The set of parameters for this mitigation action.
	//
	// You can specify only one type of parameter (in other words, you can apply only one action for each defined mitigation action).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#action_params IotMitigationAction#action_params}
	ActionParams *IotMitigationActionActionParams `field:"required" json:"actionParams" yaml:"actionParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#role_arn IotMitigationAction#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A unique identifier for the mitigation action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#action_name IotMitigationAction#action_name}
	ActionName *string `field:"optional" json:"actionName" yaml:"actionName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iot_mitigation_action#tags IotMitigationAction#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

