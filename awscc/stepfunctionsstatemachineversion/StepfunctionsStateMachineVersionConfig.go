package stepfunctionsstatemachineversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StepfunctionsStateMachineVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_version#state_machine_arn StepfunctionsStateMachineVersion#state_machine_arn}.
	StateMachineArn *string `field:"required" json:"stateMachineArn" yaml:"stateMachineArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_version#description StepfunctionsStateMachineVersion#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_version#state_machine_revision_id StepfunctionsStateMachineVersion#state_machine_revision_id}.
	StateMachineRevisionId *string `field:"optional" json:"stateMachineRevisionId" yaml:"stateMachineRevisionId"`
}

