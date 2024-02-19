package stepfunctionsstatemachinealias

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StepfunctionsStateMachineAliasConfig struct {
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
	// The settings to enable gradual state machine deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_alias#deployment_preference StepfunctionsStateMachineAlias#deployment_preference}
	DeploymentPreference *StepfunctionsStateMachineAliasDeploymentPreference `field:"optional" json:"deploymentPreference" yaml:"deploymentPreference"`
	// An optional description of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_alias#description StepfunctionsStateMachineAlias#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_alias#name StepfunctionsStateMachineAlias#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The routing configuration of the alias.
	//
	// One or two versions can be mapped to an alias to split StartExecution requests of the same state machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/stepfunctions_state_machine_alias#routing_configuration StepfunctionsStateMachineAlias#routing_configuration}
	RoutingConfiguration interface{} `field:"optional" json:"routingConfiguration" yaml:"routingConfiguration"`
}

