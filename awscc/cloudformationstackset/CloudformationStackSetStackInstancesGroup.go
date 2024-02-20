package cloudformationstackset


type CloudformationStackSetStackInstancesGroup struct {
	// The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#deployment_targets CloudformationStackSet#deployment_targets}
	DeploymentTargets *CloudformationStackSetStackInstancesGroupDeploymentTargets `field:"required" json:"deploymentTargets" yaml:"deploymentTargets"`
	// The names of one or more Regions where you want to create stack instances using the specified AWS account(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#regions CloudformationStackSet#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// A list of stack set parameters whose values you want to override in the selected stack instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#parameter_overrides CloudformationStackSet#parameter_overrides}
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
}

